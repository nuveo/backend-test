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

// Credential contains user credentials.
type Credential struct {
	User     string
	Password string
	Database string
}

func main() {
	if len(os.Args) <= 3 {
		log.Fatalf("Failed to start application: arguments are missing")
	}
	c := Credential{User: os.Args[1], Password: os.Args[2], Database: os.Args[3]}

	var a App
	if err := a.Connect(c.User, c.Password, c.Database); err != nil {
		log.Fatalf("Failed to open a database connection: %v", err)
	}

	if err := a.Prepare(); err != nil {
		log.Fatalf("Failed to prepare database: %v", err)
	}

	a.Routes()

	queue.New()

	log.Printf("Application running on localhost%s", port)
	a.Run(port)
}
