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
)

const createTable = `CREATE TABLE IF NOT EXISTS workflows (
	uuid SERIAL,
	status status_t DEFAULT 'inserted',
	data JSONB NOT NULL,
	steps text[],
	CONSTRAINT workflows_pkey PRIMARY KEY (uuid)
);`

func main() {
	// Receber flags com dados - TODO
	log.Println("Database credentials succesfully received")

	var a App
	if err := a.Database(luser, lpassword, ldbname); err != nil {
		log.Fatalf("The system couldn't open a database connection: %v", err)
	}
	log.Println("Database succesfully connected")

	if _, err := a.DB.Exec(createTable); err != nil {
		log.Fatalf("The system couldn't create workflow table: %v", err)
	}
	log.Println("Workflow table ready to go")

	a.Routes()
	log.Println("Routes succesfully initiated")

	log.Println("Application running at port " + port)
	a.Run(port)
}
