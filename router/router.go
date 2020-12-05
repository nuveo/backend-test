package router

import (
	"backend-test/middlewares"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/workflow", middlewares.GetAllWorkflow).Methods("GET", "OPTIONS")
	router.HandleFunc("/workflow", middlewares.CreateWorkflow).Methods("POST", "OPTIONS")
	router.HandleFunc("/workflow/{uuid}", middlewares.UpdateWorkflow).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/workflow/consume", middlewares.ConsumeWorkflow).Methods("GET", "OPTIONS")

	return router
}
