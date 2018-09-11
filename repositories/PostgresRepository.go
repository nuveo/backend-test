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
var host = "localhost"
var port = "5432"
var user = "nuveo"
var dbname = "nuveo"
var password = "nuveo"
var sslmode = "disable"

var dns = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)

//PostgresRepository ...
type PostgresRepository struct {
	Db *gorm.DB
}

//NewConnection ...
func NewConnection() *gorm.DB {

	dbConn, err := gorm.Open("postgres", dns)

	dbConn.LogMode(true)

	if err != nil {

		log.Fatalln("Error in connect to database", err)
	}

	return dbConn

}

// FindAll returns the list of worflow
func (r PostgresRepository) FindAll() []models.Workflow {

	var workflow []models.Workflow
	r.Db.Find(&workflow)

	return workflow
}

// Save adds a Workflow in the DB
func (r *PostgresRepository) Save(workflow models.Workflow) bool {

	r.Db.Save(&workflow)
	fmt.Println("Added New Product ID- ", workflow.UUID)
	return true
}

//Update adds a Workflow in the DB
func (r PostgresRepository) Update(workflowNew models.Workflow) bool {

	fmt.Println("New Status Updated: ", workflowNew.Status.Value())
	errUpdate := r.Db.Model(&workflowNew).Where("uuid = ?", workflowNew.UUID).Update("status", workflowNew.Status).Error

	if errUpdate != nil {

		log.Fatalln("Error in insert data", errUpdate)

		return false
	}
	return true
}

// FindByUUID finds a Workflow by UUID
func (r PostgresRepository) FindByUUID(uuidValue uuid.UUID) models.Workflow {

	var workflow models.Workflow
	r.Db.Where("uuid = ?", uuidValue).First(&workflow)
	return workflow
}

//ConsumeFromQueue by Queue and returns the list of workflows
func (r PostgresRepository) ConsumeFromQueue() []models.Workflow {

	var workflow []models.Workflow
	r.Db.Find(&workflow)
	return workflow
}
