// Package respositories provides types to handles data access
package repositories

import (
	"backend-test/models"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"

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

//PostgresRepository implements Repository interface to allow access to
//Postgres database
type PostgresRepository struct {
	Db *gorm.DB
}

//NewConnection  cretes a connection to a Postgres database
func NewConnection() (*gorm.DB, error) {

	dbConn, err := gorm.Open("postgres", dns)
	dbConn.LogMode(true)
	return dbConn, err
}

// FindAll returns the list of worflow
func (r PostgresRepository) FindAll() ([]models.Workflow, error) {

	var workflow []models.Workflow
	err := r.Db.Find(&workflow).Error
	return workflow, err
}

// Save adds a Workflow in the DB
func (r *PostgresRepository) Save(workflow models.Workflow) (models.Workflow, error) {

	err := r.Db.Save(&workflow).Error
	return workflow, err
}

//Updates a Workflow in the DB
func (r *PostgresRepository) Update(workflowNew models.Workflow) (models.Workflow, error) {

	err := r.Db.Model(&workflowNew).Where("uuid = ?", workflowNew.UUID).Update("status", workflowNew.Status).Error
	return workflowNew, err
}

// FindByUUID finds a Workflow by UUID
func (r PostgresRepository) FindByUUID(uuidValue uuid.UUID) (models.Workflow, error) {

	var workflow models.Workflow
	err := r.Db.Where("uuid = ?", uuidValue).First(&workflow).Error
	return workflow, err
}

// FindByUUID finds a Workflow by UUID
func (r PostgresRepository) FindByStatus(status models.WorkflowStatus) ([]models.Workflow, error) {

	var workflowList []models.Workflow
	err := r.Db.Where("status = ?", status).Find(&workflowList).Error
	return workflowList, err
}

//ConsumeFromQueue by Queue and returns the list of workflows
func (r PostgresRepository) ConsumeFromQueue() ([]models.Workflow, error) {

	workflowList, err := r.FindByStatus(models.Consumed)
	return workflowList, err
}
