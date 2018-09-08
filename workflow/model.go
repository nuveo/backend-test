package workflow

import (
	"encoding/json"

	"github.com/satori/go.uuid"
)

// type User struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// type JwtToken struct {
// 	Token string `json:"token"`
// }

type Exception struct {
	Message string `json:"message"`
}

type WorkflowStatus int

const (
	Inserted WorkflowStatus = 0
	Consumed WorkflowStatus = 1
)

func (status WorkflowStatus) String() string {

	statusName := []string{
		"Inserted",
		"Consumed"}

	if status != Inserted && status != Consumed {
		return "Unknown"
	}
	return statusName[status]
}

// Workflow represents an workflow item
type Workflow struct {
	UUID   uuid.UUID       `json:"uuid"`
	Status WorkflowStatus  `json:"status"`
	Data   json.RawMessage `json:"data"`
	Steps  []string        `json:"steps"`
}

// Workflows is an array of Workflow object
type Workflows []Workflow
