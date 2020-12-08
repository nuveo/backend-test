package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"io/ioutil"
	"os"
	"time"
)

var (
	ErrorWorkFlowNotFound = errors.New("couldn't find workflow")
)

type WorkflowServer interface {
	Create(r *CreateWorkflowRequest) (*Workflow, error)
	UpdateStatus(uuid string, r *UpdateWorkflowRequest) (*Workflow, error)
	ListAll() ([]*Workflow, error)
	ConsumeWorkflow() ([]byte, error)
	ConvertToCSV(data []byte) (*os.File, error)
}

type WorkflowService struct {
	db  *sql.DB
	pch *amqp.Channel
	cch *amqp.Channel
}

func (ws *WorkflowService) ConsumeWorkflow() ([]byte, error) {
	consume, err := ws.cch.Consume(
		"workflows",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	for {
		timeout := time.After(2 * time.Second)
		var uuid string
		var msg amqp.Delivery
		select {
		case <-timeout:
			return nil, errors.New("could not consume workflow")
		case msg = <-consume:
			uuid = string(msg.Body)
			ws.cch.Ack(msg.DeliveryTag, false)
		}

		row := ws.db.QueryRow("UPDATE workflows SET status='consumed' WHERE uuid=$1 AND status='inserted' RETURNING data", uuid)
		if err := row.Err(); err != nil && err != sql.ErrNoRows {
			ws.cch.Nack(msg.DeliveryTag, false, true)
			return nil, err
		} else if err := row.Err(); err == nil {
			var data []byte
			err := row.Scan(&data)
			if err != nil {
				ws.cch.Nack(msg.DeliveryTag, false, true)
				return nil, err
			}

			return data, nil
		}
	}
}

func (ws WorkflowService) ConvertToCSV(data []byte) (*os.File, error) {
	var list []map[string]interface{}
	file, err := ioutil.TempFile(os.TempDir(), "csv-")
	if err != nil {
		return nil, err
	}
	writer := csv.NewWriter(file)

	json.Unmarshal(data, &list)

	header := make([]string, 0)
	for columnName := range list[0] {
		header = append(header, columnName)
	}

	for _, obj := range list {
		row := make([]string, 0)

		for _, column := range header {
			if i, ok := obj[column].(float64); ok {
				s := fmt.Sprintf("%f", i)
				row = append(row, s)
			} else {
				row = append(row, obj[column].(string))
			}
		}
		err := writer.Write(row)
		if err != nil {
			return nil, err
		}
		writer.Flush()
	}

	return file, nil
}

func (ws *WorkflowService) Create(r *CreateWorkflowRequest) (*Workflow, error) {
	statement, err := ws.db.Prepare("INSERT INTO workflows (\"data\", steps) VALUES ($1, $2) RETURNING *")
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(r.Data)
	if err != nil {
		return nil, err
	}
	data := string(b)

	b, err = json.Marshal(r.Steps)
	if err != nil {
		return nil, err
	}

	steps := string(b)

	result := statement.QueryRow(data, steps)
	if result.Err() != nil {
		return nil, err
	}

	var workflow Workflow
	var returnedData []byte
	var returnedSteps []byte

	if err := result.Scan(&workflow.UUID, &workflow.Status, &returnedData, &returnedSteps); err != nil {
		return nil, err
	}
	json.Unmarshal(returnedData, &workflow.Data)
	json.Unmarshal(returnedSteps, &workflow.Steps)
	err = ws.pch.Publish("", "workflows", false, false,
		amqp.Publishing{
			Body: []byte(workflow.UUID),
		},
	)

	if err != nil {
		return nil, err
	}

	return &workflow, nil
}

func (ws *WorkflowService) UpdateStatus(uuid string, r *UpdateWorkflowRequest) (*Workflow, error) {
	result := ws.db.QueryRow("UPDATE workflows SET status=$2 WHERE uuid=$1 RETURNING *", uuid, r.Status)
	if err := result.Err(); err != nil {
		return nil, err
	}

	var workflow Workflow
	if err := result.Scan(&workflow.UUID, &workflow.Status, &workflow.Data, &workflow.Steps); err != nil {
		return nil, err
	}

	if len(workflow.UUID) == 0 {
		return nil, ErrorWorkFlowNotFound
	}
	return &workflow, nil
}

func (ws *WorkflowService) ListAll() ([]*Workflow, error) {
	results, err := ws.db.Query("SELECT * FROM workflows")
	if err != nil {
		return nil, err
	}
	workflows := make([]*Workflow, 0)
	for results.Next() {
		var w Workflow
		var data []byte
		var steps []byte
		if err := results.Scan(&w.UUID, &w.Status, &data, &steps); err != nil {
			return nil, err
		}
		json.Unmarshal(data, &w.Data)
		json.Unmarshal(steps, &w.Steps)
		workflows = append(workflows, &w)
	}

	return workflows, nil
}
