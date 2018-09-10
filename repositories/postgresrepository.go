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
const SERVER = "host=192.168.99.100 port=5432 user=nuveo dbname=nuveo password=nuveo sslmode=disable"

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

	db, err := gorm.Open("postgres", SERVER)

	db.LogMode(true)

	if err != nil {

		log.Fatalln("Error in connect to database", err)
	}
	defer db.Close()

	db.Save(&workflow)
	fmt.Println("Added New Product ID- ", workflow.UUID)

	return true
}

//Update adds a Workflow in the DB
func (r PostgresRepository) Update(workflowNew models.Workflow) bool {

	db, err := gorm.Open("postgres", SERVER)
	fmt.Println("New Status Updated: ", workflowNew.Status.Value())
	db.LogMode(true)

	if err != nil {

		log.Fatalln("Error in connect to database", err)
		return false
	}
	defer db.Close()
	errUpdate := db.Model(&workflowNew).Where("uuid = ?", workflowNew.UUID).Update("status", workflowNew.Status).Error

	if errUpdate != nil {

		log.Fatalln("Error in insert data", errUpdate)

		return false
	}
	return true
}

// FindByUUID finds a Workflow by UUID
func (r PostgresRepository) FindByUUID(uuidValue uuid.UUID) models.Workflow {

	var workflow models.Workflow
	db, err := gorm.Open("postgres", SERVER)

	db.LogMode(true)

	if err != nil {

		log.Fatalln("Error in connect to database", err)
	}
	defer db.Close()

	db.Where("uuid = ?", uuidValue).First(&workflow)
	return workflow
}

//ConsumeFromQueue by Queue and returns the list of workflows
func (r PostgresRepository) ConsumeFromQueue() []models.Workflow {
	return r._data
}
