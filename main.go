package main

import (
	"log"
	"os"
)

var (
	port  = ":8889"
	queue ItemQueue
)

const createEnum = `DO $$ BEGIN
CREATE TYPE status_t AS ENUM ('inserted', 'consumed');
EXCEPTION
WHEN duplicate_object THEN null;
END $$;`
const createExtension = `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
const createTable = `CREATE TABLE IF NOT EXISTS workflows (
	uuid UUID DEFAULT uuid_generate_v4(),
	status status_t DEFAULT 'inserted' NOT NULL,
	data JSONB NOT NULL,
	steps text[] NOT NULL,
	PRIMARY KEY (uuid)
);`

// Credentials contains user credentials.
type Credentials struct {
	User     string
	Password string
	Database string
}

func main() {
	if len(os.Args) <= 3 {
		log.Fatalf("The system couldn't start application: arguments are missing")
	}
	c := Credentials{User: os.Args[1], Password: os.Args[2], Database: os.Args[3]}
	log.Println("Database credentials successfully received")

	var a App
	if err := a.Connect(c.User, c.Password, c.Database); err != nil {
		log.Fatalf("The system couldn't open a database connection: %v", err)
	}
	log.Println("Database successfully connected")

	if err := a.Prepare(); err != nil {
		log.Fatalf("The system couldn't prepare database: %v", err)
	}
	log.Println("Database successfully prepared")

	a.Routes()
	log.Println("Routes successfully initiated")

	queue.New()
	log.Println("Queue successfully created")

	log.Println("Application running on localhost" + port)
	a.Run(port)
}
