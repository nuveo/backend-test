package repositories

import (
	"backend-test/models"

	"github.com/satori/go.uuid"
)

//WorkflowRepository ...
type WorkflowRepository interface {
	FindAll() ([]models.Workflow, error)
	FindByUUID(uuid.UUID) (models.Workflow, error)
	Save(models.Workflow) (models.Workflow, error)
	Update(workflowNew models.Workflow) (models.Workflow, error)
	ConsumeFromQueue() ([]models.Workflow, error)
}
