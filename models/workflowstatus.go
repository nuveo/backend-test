package models

//WorkflowStatus represents a Enum type of workflow's status
type WorkflowStatus int

//Those are the possibles valeus for an Enum type of workflow's status
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
