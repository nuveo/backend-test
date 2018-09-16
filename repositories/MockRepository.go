// Package respositories provides types to handles data access
package repositories

import (
	"backend-test/models"
	"fmt"

	"github.com/satori/go.uuid"
)

//MockRepository ...
type MockRepository struct {
	_data []models.Workflow
}

// FindAll returns the list of worflow
func (r MockRepository) FindAll() []models.Workflow {
	return r._data
}

// Save adds a Workflow in the DB
func (r *MockRepository) Save(workflow models.Workflow) bool {

	r._data = append(r._data, workflow)
	fmt.Println("Added New Product ID- ", workflow.UUID)

	return true
}

// FindByUUID finds a Workflow by UUID
func (r MockRepository) FindByUUID(uuidValue uuid.UUID) models.Workflow {

	var workflow models.Workflow
	for _, element := range r._data {

		if element.UUID == uuidValue {

			workflow = element
			break
		}

	}
	return workflow
}

//ConsumeFromQueue by Queue and returns the list of workflows
func (r MockRepository) ConsumeFromQueue() []models.Workflow {
	return r._data
}
