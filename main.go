package main

import (
	"backend-test/helpers"
	"backend-test/router"
	"fmt"
	"log"
	"net/http"
)

var (
	queue helpers.ElementQueue
)

func main() {

	r := router.Router()

	fmt.Println("Listening on :8080")

	queue.Create()

	log.Fatal(http.ListenAndServe(":8080", r))
}
