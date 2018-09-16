package main

import (
	"log"
)

var (
	luser     = "postgres"
	lpassword = "1234"
	ldbname   = "nuveo"
	port      = ":8889"
	db        = "postgres"
	queue     ItemQueue
)

const createEnum = `DO $$ BEGIN
CREATE TYPE status_t AS ENUM ('inserted', 'consumed');
EXCEPTION
WHEN duplicate_object THEN null;
END $$;`
const createExtension = `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
const createTable = `CREATE TABLE IF NOT EXISTS workflows (
	uuid UUID DEFAULT uuid_generate_v4 (),
	status status_t DEFAULT 'inserted',
	data JSONB NOT NULL,
	steps text[],
	PRIMARY KEY (uuid)
);`

func main() {
	// Receber flags com dados
	log.Println("Database credentials succesfully received")

	var a App

	if err := a.Database(luser, lpassword, ldbname); err != nil {
		log.Fatalf("The system couldn't open a database connection: %v", err)
	}
	log.Println("Database succesfully connected")

	if _, err := a.DB.Exec(createEnum); err != nil {
		log.Fatalf("The system couldn't create status type: %v", err)
	}
	if _, err := a.DB.Exec(createTable); err != nil {
		log.Fatalf("The system couldn't create workflow table: %v", err)
	}
	log.Println("Workflow table ready to go")

	a.Routes()
	log.Println("Routes succesfully initiated")

	queue.New()
	log.Println("Queue succesfully created")

	log.Println("Application running at port " + port)
	a.Run(port)
}
