package main

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	Router  *mux.Router
	service WorkflowServer
}

func NewApp(service WorkflowServer) (*App, error) {
	if service == nil {
		return nil, errors.New("service can't be nil")
	}

	app := App{
		Router:  mux.NewRouter(),
		service: service,
	}

	app.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(writer, request)
		})
	})

	RegisterRoutes(&app)
	return &app, nil
}
