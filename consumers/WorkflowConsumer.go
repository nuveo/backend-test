// Package consumers provides consumers that mostly waits to receive messages
// from a rabbitqm's queue
package consumers

import (
	"backend-test/models"
	"backend-test/repositories"
	"encoding/json"
	"log"

	"github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

//WorkflowConsumer is a rabbitqm's consumer client
type WorkflowConsumer struct {
	Repo repositories.WorkflowRepository
}

//Run waits to receive messages from rabbitqm's qm and insert it on the
//database
func (wc *WorkflowConsumer) Run() error {

	log.Println("Starting Queue Consumer")

	//Connects do DB
	var db, _ = repositories.NewConnection()
	var repo = &repositories.PostgresRepository{Db: db}
	wc.Repo = repo

	// Connects to rabbitqm
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	//Opens a rabbitqm channel
	ch, _ := conn.Channel()
	defer ch.Close()

	//Binds to a rabbitqm queue
	q, _ := ch.QueueDeclare(
		"nuveo", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // argumentsj
	)

	//Reading messagem from the queue
	deliveries, _ := ch.Consume(
		q.Name,         // queue
		"workflow-api", // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		true,           // no-wait
		nil,            // args
	)

	//Starting a reading loop waiting for messages
	forever := make(chan bool)
	go func() {
		var workflow models.Workflow
		for d := range deliveries {

			//Transforms the receiving message to a json type
			if err := json.Unmarshal(d.Body, &workflow); err != nil {
				log.Println(err)
			}

			//Create a new UUID to a workflow item
			workflow.UUID, _ = uuid.NewV4()
			//Changes the workflow status to "Consumed"
			workflow.Status = models.Consumed

			//Saving into database
			wc.Repo.Save(workflow)
		}
	}()
	<-forever
	return nil
}
