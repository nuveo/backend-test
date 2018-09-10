package repositories

import (
	"backend-test/models"

	"github.com/satori/go.uuid"
)

//WorkflowRepository ...
type WorkflowRepository interface {
	FindAll() []models.Workflow
	FindByUUID(uuid.UUID) models.Workflow
	Save(models.Workflow) bool
	Update(workflowNew models.Workflow) bool
	ConsumeFromQueue() []models.Workflow
}
