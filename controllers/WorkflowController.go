// Package controllers provides controllers types
package controllers

import (
	"backend-test/exceptions"
	"backend-test/models"
	"backend-test/repositories"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

//Controller receives a request and defines the function that handles it
type Controller struct {
	Repo repositories.WorkflowRepository
}

//ListWorkflows finds all workflows and returns as json list. If error, return
//a json with error's message
func (c *Controller) ListWorkflows(w http.ResponseWriter, r *http.Request) {

	//list of all workflows
	workflowList, err := c.Repo.FindAll()

	//Return a custom exception with erros's message
	if err != nil {
		data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}

	//Setting response header
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//Marshals to json type
	data, _ := json.Marshal(workflowList)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

//AddWorkflow POST inserts a workflow to repository data. Returns the Workflow
//that was created
func (c *Controller) AddWorkflow(w http.ResponseWriter, r *http.Request) {

	var workflow models.Workflow

	//read the body of the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}

	//Marshal body from a json string to a workflow item
	if err := r.Body.Close(); err != nil {
		data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}

	// Unmarshall body contents as a type Workflow
	if err := json.Unmarshal(body, &workflow); err != nil {

		w.WriteHeader(422) // unprocessable entity
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(data)
			return
		}
	}

	// adds the workflow to the repository
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

//UpdateWorkflow Update status from a specific workflow
func (c *Controller) UpdateWorkflow(w http.ResponseWriter, r *http.Request) {

	var workflow models.Workflow

	vars := mux.Vars(r)

	// read the body of the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
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
		w.WriteHeader(422)
		data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
		w.Write(data)
		return
	}

	if err := json.NewEncoder(w).Encode(err); err != nil {
		data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
		w.Write(data)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Gets UUID from path variable /{UUID}
	uuidValue := vars["UUID"]
	//Create a UUID type from string

	uuidWorkflow, err := uuid.FromString(uuidValue)
	if err != nil {
		data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
		w.Write(data)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// updates the workflow in the repository
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

//ConsumeWorkflows finds all workflows was consumed from a queue and returns as
//json list. If error, return a json with error's message
func (c *Controller) ConsumeWorkflows(w http.ResponseWriter, r *http.Request) {

	var workflowData models.WorkflowData

	//Get a list for consumed items from a queue
	workflowList, err := c.Repo.ConsumeFromQueue()
	if err != nil {
		data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(data)
		return
	}

	//write workflow list as csv file
	record := []string{}
	b := &bytes.Buffer{}   // creates IO Writer
	wr := csv.NewWriter(b) // creates a csv writer that uses the io buffer.

	//Adds headers
	record = append(record, "Name")
	record = append(record, "Description")
	wr.Write(record)
	record = nil

	//Adds workflow items to csv writer
	for _, workflow := range workflowList {

		err := json.Unmarshal(workflow.Data, &workflowData)
		if err != nil {
			data, _ := json.Marshal(exceptions.WorkflowException{Message: err.Error()})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(data)
			return
		}
		record = append(record, workflowData.Name)
		record = append(record, workflowData.Description)

		wr.Write(record)
		record = nil
	}
	//Flush data to writer
	wr.Flush()

	//Setting header
	w.Header().Set("Content-Type", "text/csv; charset=UTF-8")
	w.Header().Set("Content-Disposition", "attachment;filename=WorkflowData.csv")
	w.WriteHeader(http.StatusOK)

	w.Write(b.Bytes())
	return
}
