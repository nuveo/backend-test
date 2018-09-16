// Package consumers provides  custom exceptions to workflow REST API
package exceptions

//WorkflowException is a customized exception to Workflow API
type WorkflowException struct {
	Message string `json:"message"`
}
