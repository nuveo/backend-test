package repositories

import (
	"backend-test/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"

	//
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// SERVER the DB server
const SERVER = "host=localhost port=5432 user=nuveo dbname=nuveo password=nuveo sslmode=disable"

//PostgresRepository ...
type PostgresRepository struct {
	_data []models.Workflow
}

// FindAll returns the list of worflow
func (r PostgresRepository) FindAll() []models.Workflow {
	db, err := gorm.Open("postgres", SERVER)

	db.LogMode(true)

	if err != nil {

		log.Fatalln("Error in connect to database", err)
	}
	defer db.Close()

	var workflow []models.Workflow
	db.Find(&workflow)

	return workflow
}

// Save adds a Workflow in the DB
func (r *PostgresRepository) Save(workflow models.Workflow) bool {

	r._data = append(r._data, workflow)
	fmt.Println("Added New Product ID- ", workflow.UUID)

	return true
}

// FindByUUID finds a Workflow by UUID
func (r PostgresRepository) FindByUUID(uuidValue uuid.UUID) models.Workflow {

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
func (r PostgresRepository) ConsumeFromQueue() []models.Workflow {
	return r._data
}
