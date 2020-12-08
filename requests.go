package main

type CreateWorkflowRequest struct {
	Data  []interface{} `json:"data"`
	Steps []interface{} `json:"steps"`
}

type UpdateWorkflowRequest struct {
	Status Status `json:"status"`
}
