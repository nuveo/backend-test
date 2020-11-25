package models

import "encoding/json"

type WorkflowStatus string

const (
	INSERTED WorkflowStatus = "inserted"
	CONSUMED WorkflowStatus = "consumed"
)

type Workflow struct {
	UUID   string          `json:"UUID"`
	Status WorkflowStatus  `json:"status"`
	Data   json.RawMessage `json:"data"`
	Steps  []string        `json:"steps"`
}

type Response struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type DataStatus struct {
	Data  json.RawMessage `json:"data"`
	Steps []string        `json:"steps"`
}

type DataJSON struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
