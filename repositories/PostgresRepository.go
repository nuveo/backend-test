package repositories

import (
	"backend-test/models"
	"encoding/json"
	"fmt"
	"log"

	// "log"

	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"github.com/streadway/amqp"

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
func NewConnection() (*gorm.DB, error) {

	dbConn, err := gorm.Open("postgres", dns)
	dbConn.LogMode(true)
	// if err != nil {
	// 	log.Fatalln("Error in connect to database", err)
	// }
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

//ConsumeFromQueue by Queue and returns the list of workflows
func (r PostgresRepository) ConsumeFromQueue() ([]models.Workflow, error) {

	// workflowList := []models.Workflow{}
	var workflow models.Workflow
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	ch, err := conn.Channel()
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"nuveo", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // argumentsj
	)
	deliveries, err := ch.Consume(
		q.Name,         // queue
		"workflow-api", // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		true,           // no-wait
		nil,            // args
	)

	workflowList := []models.Workflow{}
	msgs := make(chan []byte)
	done := make(chan error)

	go func(deliveries <-chan amqp.Delivery, done chan error, message chan []byte) {

		log.Println("Start reading message")
		for d := range deliveries {
			message <- d.Body
			log.Println("Message was read")
		}
		done <- nil
	}(deliveries, done, msgs)
	for {
		if err := json.Unmarshal(<-msgs, &workflow); err != nil {
			log.Println(err)
		}
		workflow.Status = models.Consumed
		log.Println(workflow)
		log.Printf("%s\n", <-msgs)
		workflowList = append(workflowList, workflow)
		if len(workflowList) > 10 {
			break
		}
	}

	return workflowList, err
}
