package exceptions

//WorkflowException is a customized exception to workflow-api
type WorkflowException struct {
	Message string `json:"message"`
}
