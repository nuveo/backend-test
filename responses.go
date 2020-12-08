package main

import (
	"encoding/json"
	"log"
)

type CreateWorkflowResponse struct {
	Err  error `json:"error,omitempty"`
	Data struct {
		UUID string `json:"uuid"`
	} `json:"data,omitempty"`
}

func NewCreateWorkflowResponse(w *Workflow, err error) string {
	if (w == nil && err == nil) || (w != nil && err != nil) {
		log.Fatal("NewCreateWorkFlowResponse should have only a workflow or an error but not both or none!")
	}

	bytes, err := json.Marshal(CreateWorkflowResponse{
		Err: err,
		Data: struct {
			UUID string `json:"uuid"`
		}{
			UUID: w.UUID,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}

type UpdateWorkflowResponse struct {
	Err  error `json:"error,omitempty"`
	Data struct {
		UUID   string `json:"uuid"`
		Status Status `json:"status"`
	} `json:"data,omitempty"`
}

func NewUpdateWorkflowResponse(w *Workflow, err error) string {
	if (w == nil && err == nil) || (w != nil && err != nil) {
		log.Fatal("NewUpdateWorkFlowResponse should have only a workflow or an error but not both or none!")
	}

	bytes, err := json.Marshal(UpdateWorkflowResponse{
		Err: err,
		Data: struct {
			UUID   string `json:"uuid"`
			Status Status `json:"status"`
		}{
			UUID:   w.UUID,
			Status: w.Status,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}

type ListAllWorkflowsResponse struct {
	Err  error       `json:"error,omitempty"`
	Data []*Workflow `json:"data"`
}

func NewListAllWorkflowsResponse(ws []*Workflow, err error) string {
	if (ws == nil && err == nil) || (ws != nil && err != nil) {
		log.Fatal("ListAllWorkflowsResponse should have only a workflow or an error but not both or none!")
	}

	bytes, err := json.Marshal(ListAllWorkflowsResponse{
		Err:  err,
		Data: ws,
	})

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}
