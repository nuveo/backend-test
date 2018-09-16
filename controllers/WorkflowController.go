/********************************************************
* Author: Vagner Clementino
* Date: 16/09/2018
********************************************************/
package controllers

import (
	"backend-test/exceptions"
	"backend-test/models"
	"backend-test/repositories"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

//Controller ...
type Controller struct {
	Repo repositories.WorkflowRepository
}

//ListWorkflows GET /workflow
func (c *Controller) ListWorkflows(w http.ResponseWriter, r *http.Request) {
	workflows, err := c.Repo.FindAll() // list of all products
	if err != nil {
		data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data, _ := json.Marshal(workflows)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// AddWorkflow POST /workflow
func (c *Controller) AddWorkflow(w http.ResponseWriter, r *http.Request) {
	var workflow models.Workflow

	// read the body of the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	// log.Println(body)

	if err != nil {
		data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}

	if err := r.Body.Close(); err != nil {
		data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}

	if err := json.Unmarshal(body, &workflow); err != nil {
		// unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(data)
			return
		}
	}

	log.Println(workflow)
	// adds the product to the DB
	_, err = c.Repo.Save(workflow)
	if err != nil {
		data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(workflow)
	w.Write(data)
	return
}

//UpdateWorkflow Update status from a specific workflow PATCH /
func (c *Controller) UpdateWorkflow(w http.ResponseWriter, r *http.Request) {
	var workflow models.Workflow

	vars := mux.Vars(r)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}

	if err := r.Body.Close(); err != nil {
		data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}

	// unmarshall body contents as a type Candidate
	if err := json.Unmarshal(body, &workflow); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
		w.Write(data)
		return
	}

	if err := json.NewEncoder(w).Encode(err); err != nil {
		// log.Fatalln("Error UpdateProduct unmarshalling data", err)
		data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
		w.Write(data)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//uuidValue := strings.Replace(vars["uuid"], "-", "", -1) // param id
	uuidValue := vars["UUID"]
	log.Println("Parsed UUID:" + uuidValue)

	uuidWorkflow, err := uuid.FromString(uuidValue)
	if err != nil {
		log.Fatalln("Something went wrong:", err)
		data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
		w.Write(data)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// updates the product in the DB
	workflow.UUID = uuidWorkflow
	_, err = c.Repo.Update(workflow)
	if err != nil {
		data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	return
}

// ConsumeWorkflows GET /consume
func (c *Controller) ConsumeWorkflows(w http.ResponseWriter, r *http.Request) {
	workflowList, err := c.Repo.ConsumeFromQueue()
	if err != nil {
		data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}

	// just some test data to use for the wr.Writer() method below.
	record := []string{}
	var workflowData models.WorkflowData

	b := &bytes.Buffer{}   // creates IO Writer
	wr := csv.NewWriter(b) // creates a csv writer that uses the io buffer.

	//Adds headers
	record = append(record, "Name")
	record = append(record, "Description")
	wr.Write(record)

	record = nil
	for _, workflow := range workflowList {

		err2 := json.Unmarshal(workflow.Data, &workflowData)
		if err2 != nil {
			fmt.Println("error:", err2)
			os.Exit(1)
		}
		record = append(record, workflowData.Name)
		record = append(record, workflowData.Description)

		wr.Write(record)
		record = nil
	}
	wr.Flush()
	// data, _ := json.Marshal(workflows)
	w.Header().Set("Content-Type", "text/csv; charset=UTF-8")
	w.Header().Set("Content-Disposition", "attachment;filename=WorkflowData.csv")
	w.WriteHeader(http.StatusOK)
	w.Write(b.Bytes())
	return
}
