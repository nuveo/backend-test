// Entrypoint for API
package main

import (
	"backend-test/workflow"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	// port := os.Getenv("PORT")
	port := "8000"

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := workflow.NewRouter() // create routes

	// These two lines are important in order to allow access from the front-end side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// Launch server with CORS validations
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
