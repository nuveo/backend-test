package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load(".env")
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	connection, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatal(err)
	}

	pchannel, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}

	pchannel.QueueDeclare(
		"workflows",
		true,
		false,
		false,
		false,
		nil,
	)

	cchannel, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}

	service := &WorkflowService{db: db, cch: cchannel, pch: pchannel}

	app, err := NewApp(service)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		fmt.Fprintln(writer, "{\"data\":\"Hello, World!\"}")
	})

	http.Handle("/", app.Router)
	http.Handle("*", http.NotFoundHandler())
	fmt.Println("Starting server at http://localhost:8080/")
	fmt.Println("Say hello at http://localhost:8080/hello")
	fmt.Println("List workflows at http://localhost:8080/workflow")
	fmt.Println("Consume workflows at http://localhost:8080/workflow/consume")
	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
