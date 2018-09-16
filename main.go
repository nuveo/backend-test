package main

import (
	"log"
	"os"
)

var (
	port  = ":8889"
	queue ItemQueue
)

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
	credential := Credential{User: os.Args[1], Password: os.Args[2], Database: os.Args[3]}

	var app App
	if err := app.Connect(credential.User, credential.Password, credential.Database); err != nil {
		log.Fatalf("Failed to open a database connection: %v", err)
	}

	if err := app.Database(); err != nil {
		log.Fatalf("Failed to prepare database: %v", err)
	}

	app.Routes()

	queue.New()

	log.Printf("Application running on localhost%s", port)
	app.Run(port)
}
