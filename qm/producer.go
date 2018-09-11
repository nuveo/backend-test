//Package qm...
package main

import (
	"backend-test/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {

	var host = "localhost"
	var port = "5432"
	var user = "nuveo"
	var dbname = "nuveo"
	var password = "nuveo"
	var sslmode = "disable"

	var dns = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)

	db, err := gorm.Open("postgres", dns)
	defer db.Close()
	failOnError(err, "Failed to connect to database")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	var workflows []models.Workflow
	err = db.Find(&workflows).Where(models.Inserted).Error

	for _, element := range workflows {

		body := element.Data

		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		log.Printf(" [x] Sent %s", body)
		failOnError(err, "Failed to publish a message")
	}
}
