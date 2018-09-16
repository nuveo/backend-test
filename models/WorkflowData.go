// Package models provides entities to workflow API
package models

// WorkflowData represents a Workflow Data attibute. It will be used to produce
// a csv file
type WorkflowData struct {
	Name        string
	Description string
}
