package models

import (
	"encoding/json"

	"github.com/satori/go.uuid"
)

// Workflow represents an workflow item
type Workflow struct {
	UUID   uuid.UUID       `json:"uuid"`
	Status WorkflowStatus  `json:"status"`
	Data   json.RawMessage `json:"data"`
	Steps  []string        `json:"steps"`
}
