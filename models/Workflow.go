// Package models provides entities to workflow API
package models

import (
	"encoding/json"

	"github.com/lib/pq"
	"github.com/satori/go.uuid"
)

// Workflow represents an workflow item
type Workflow struct {
	UUID   uuid.UUID       `json:"uuid"`
	Status WorkflowStatus  `json:"status"`
	Data   json.RawMessage `json:"data"`
	Steps  pq.StringArray  `json:"steps" gorm:"type:text[]"`
}

// TableName set Workflow's table name to be `workflow`
func (Workflow) TableName() string {
	return "workflow"
}
