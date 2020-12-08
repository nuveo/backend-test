package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func (app *App) CreateWorkflow(w http.ResponseWriter, r *http.Request) {
	var proposal CreateWorkflowRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&proposal); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := NewCreateWorkflowResponse(nil, err)
		fmt.Fprint(w, response)
		return
	}

	workflow, err := app.service.Create(&proposal)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := NewCreateWorkflowResponse(nil, err)
		fmt.Fprint(w, response)
		return
	}

	response := NewCreateWorkflowResponse(workflow, nil)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, response)
}

func (app *App) UpdateWorkflowStatus(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	var proposal UpdateWorkflowRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&proposal); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := NewUpdateWorkflowResponse(nil, err)
		fmt.Fprint(w, response)
		return
	}

	if !proposal.Status.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := fmt.Sprintf("the provided status %s isn't valid", proposal.Status)
		err := errors.New(errorMessage)
		response := NewUpdateWorkflowResponse(nil, err)
		fmt.Fprint(w, response)
		return
	}

	workflow, err := app.service.UpdateStatus(parameters["uuid"], &proposal)

	if err == ErrorWorkFlowNotFound {
		w.WriteHeader(http.StatusNotFound)
		errorMessage := fmt.Sprintf("no workflow with uuid %s was found", parameters["uuid"])
		err := errors.New(errorMessage)
		response := NewUpdateWorkflowResponse(nil, err)
		fmt.Fprint(w, response)
		return

	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := NewUpdateWorkflowResponse(nil, err)
		fmt.Fprint(w, response)
		return
	}

	response := NewUpdateWorkflowResponse(workflow, nil)
	fmt.Fprint(w, response)
}

func (app *App) ListAllWorkflows(w http.ResponseWriter, r *http.Request) {
	workflows, err := app.service.ListAll()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := NewListAllWorkflowsResponse(nil, err)
		fmt.Fprint(w, response)
		return
	}

	response := NewListAllWorkflowsResponse(workflows, err)
	fmt.Fprint(w, response)
}

func (app *App) ConsumeWorkflow(w http.ResponseWriter, r *http.Request) {
	bytes, err := app.service.ConsumeWorkflow()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}

	csv, err := app.service.ConvertToCSV(bytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "DEU MERDA!")
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Add("Content-Disposition", "attachment; filename=\"data.csv\"")
	http.ServeContent(w, r, "data.csv", time.Now(), csv)
}
