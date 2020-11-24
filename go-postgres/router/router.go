package router

import (
	"go-postgres/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/workflow/consume", middleware.GetWorkflow).Methods("GET", "OPTIONS")
	router.HandleFunc("/workflow", middleware.GetAllWorkflow).Methods("GET", "OPTIONS")
	router.HandleFunc("/workflow", middleware.CreateWorkflow).Methods("POST", "OPTIONS")
	router.HandleFunc("/workflow/{id}", middleware.UpdateWorkflow).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/workflow/{id}", middleware.DeleteWorkflow).Methods("DELETE", "OPTIONS")

	return router
}
