package main

import (
	"net/http"
)

type Route struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

func RegisterRoutes(app *App) {
	routes := []Route{
		{
			Path:        "/workflow",
			Method:      http.MethodPost,
			HandlerFunc: app.CreateWorkflow,
		},
		{
			Path:        "/workflow/{uuid}",
			Method:      http.MethodPatch,
			HandlerFunc: app.UpdateWorkflowStatus,
		},
		{
			Path:        "/workflow",
			Method:      http.MethodGet,
			HandlerFunc: app.ListAllWorkflows,
		},
		{
			Path:        "/workflow/consume",
			Method:      http.MethodGet,
			HandlerFunc: app.ConsumeWorkflow,
		},
	}

	for _, route := range routes {
		app.Router.HandleFunc(route.Path, route.HandlerFunc).Methods(route.Method)
	}
}
