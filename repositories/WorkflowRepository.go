// Package respositories provides types to handles data access
package repositories

import (
	"backend-test/models"

	"github.com/satori/go.uuid"
)

//WorkflowRepository is a interface to data's access
type WorkflowRepository interface {
	// Return all workflow from repository as a slice of workflow
	FindAll() ([]models.Workflow, error)

	//Find a specific workflow item by UUID
	FindByUUID(uuid.UUID) (models.Workflow, error)

	//Find one or more workflow item by status. Return type is a workflow's
	//slice
	FindByStatus(status models.WorkflowStatus) ([]models.Workflow, error)

	//Persist a workflow to repository
	Save(models.Workflow) (models.Workflow, error)

	//Updates a workflow to repository
	Update(workflowNew models.Workflow) (models.Workflow, error)

	//Consume one or more workflow item that was consumed by a queue. Return
	//type is a workflow's slice
	ConsumeFromQueue() ([]models.Workflow, error)
}
