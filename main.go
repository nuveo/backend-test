// Entrypoint for Workflow API
package main

import (
	"backend-test/consumers"
	"backend-test/routers"
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
	//Crate routes
	router := routers.NewRouter()

	//Consumes a workflow item from a queue
	consumer := consumers.WorkflowConsumer{}
	go consumer.Run()

	// These two lines are important in order to allow access from the front-end
	// side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PATH"})

	// Launch server with CORS validations
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins,
		allowedMethods)(router)))
}
