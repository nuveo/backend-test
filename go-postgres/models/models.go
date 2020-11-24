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

type WorkflowData struct {
	Name        string
	Description string
}
