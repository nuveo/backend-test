package consumers

import (
	"backend-test/models"
	"backend-test/repositories"
	"encoding/json"
	"log"

	"github.com/satori/go.uuid"
	"github.com/streadway/amqp"
)

//WorkflowConsumer is ...
type WorkflowConsumer struct {
	Repo repositories.WorkflowRepository
}

//Run ...
func (wc *WorkflowConsumer) Run() error {

	log.Println("Starting Queue Consumer")
	var db, _ = repositories.NewConnection()
	var repo = &repositories.PostgresRepository{Db: db}
	wc.Repo = repo
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	q, _ := ch.QueueDeclare(
		"nuveo", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // argumentsj
	)
	deliveries, _ := ch.Consume(
		q.Name,         // queue
		"workflow-api", // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		true,           // no-wait
		nil,            // args
	)

	forever := make(chan bool)
	go func() {
		var workflow models.Workflow
		log.Println("Start reading message")
		for d := range deliveries {

			if err := json.Unmarshal(d.Body, &workflow); err != nil {
				log.Println(err)
			}
			workflow.Status = models.Consumed
			workflow.UUID, _ = uuid.NewV4()
			wc.Repo.Save(workflow)
			log.Println(workflow)
			log.Println("Message was read")
		}
	}()
	<-forever
	return nil
}
