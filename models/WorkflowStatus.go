// Package models provides entities to workflow API
package models

//WorkflowStatus represents a Enum type of workflow's status
type WorkflowStatus int

//Those are the possibles valeus for an Enum type of workflow's status
const (
	Inserted WorkflowStatus = 0
	Consumed WorkflowStatus = 1
)

//Value returns the description from a WorkflowStatus Enum type
func (status WorkflowStatus) String() string {

	statusName := []string{
		"Inserted",
		"Consumed"}

	if status != Inserted && status != Consumed {
		return "Unknown"
	}
	return statusName[status]
}

//Value returns the integer valeu from a WorkflowStatus Enum type
func (status WorkflowStatus) Value() int {

	return int(status)
}
