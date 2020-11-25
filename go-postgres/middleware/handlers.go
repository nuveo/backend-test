package middleware

import (
	"database/sql"
	"encoding/csv"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"go-postgres/models" // models package where Workflow schema is defined
	"go-postgres/queue"
	"log"
	"net/http" // used to access the request and response object of the api
	"os"       // used to read the environment variable

	"github.com/gorilla/mux"   // used to get the params from the route
	"github.com/joho/godotenv" // package used to read the .env file
	"github.com/lib/pq"        // postgres golang driver
)

var fila queue.IdQueue

// create connection with postgres db
func createConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

// CreateWorkflow create a workflow in the postgres db
func CreateWorkflow(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty workflow of type models.Workflow
	var ds models.DataStatus

	// decode the json request to workflow
	err := json.NewDecoder(r.Body).Decode(&ds)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v\n", err)
	}

	fmt.Printf("Queue size: %v\n", fila.Size())

	// call insert workflow function and pass the workflow
	workflow := insertWorkflow(ds)

	// put UUID of the workflow in the end of the queue
	fila.Enqueue(workflow.UUID)

	fmt.Printf("Queue size: %v\n", fila.Size())

	// format a response object
	res := models.Workflow{
		UUID:   workflow.UUID,
		Status: workflow.Status,
		Data:   workflow.Data,
		Steps:  workflow.Steps,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// GetWorkflow will return a single workflow by its id
func GetWorkflow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if fila.IsEmpty() {
		log.Fatalf("Empty queue.\n")
	}

	// get the id from the head of the queue
	uuid := fila.First()

	// call the getWorkflow function with workflow id to retrieve a single workflow
	workflow, err := getWorkflow(*uuid)

	if err != nil {
		log.Fatalf("Unable to get workflow. %v\n", err)
	}

	fmt.Printf("Queue size: %v\n", fila.Size())

	if fila.IsEmpty() {
		log.Fatalf("Empty queue.\n")
	} else {
		// consume workflow id from the head of the queue
		fila.Dequeue()

		//update workflow status to CONSUMED
		updateWorkflow(*uuid, models.WorkflowStatus(models.CONSUMED))
	}

	fmt.Printf("Queue size: %v\n", fila.Size())

	// format a json data to be converted into csv
	data := models.DataJSON{}
	err = json.Unmarshal([]byte(workflow.Data), &data)

	if err != nil {
		log.Fatalf("Unable to convert json file. %v\n", err)
	}

	// create csv file with the id nas the name
	file, err := os.Create("./output/" + *uuid + ".csv")

	if err != nil {
		log.Fatalf("Unable to create csv file. %v\n", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	// write data of json into row
	var row []string
	row = append(row, data.Name)
	row = append(row, data.Description)

	// write lines at the csv
	writer.Write([]string{"name", "description"})
	writer.Write(row)

	defer writer.Flush()
}

// GetAllWorkflow will return all the workflows
func GetAllWorkflow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get all the workflows in the db
	workflows, err := getAllWorkflows()

	if err != nil {
		log.Fatalf("Unable to get all workflow. %v\n", err)
	}

	fmt.Printf("Queue size: %v\n", fila.Size())

	// send all the workflows as response
	json.NewEncoder(w).Encode(workflows)
}

// UpdateWorkflow update workflow's detail in the postgres db
func UpdateWorkflow(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the workflowid from the request params, key is "id"
	params := mux.Vars(r)

	// get the id from the workflow to update
	uuid := params["UUID"]

	fmt.Printf("Queue size: %v\n", fila.Size())

	// update the workflow as INSERTED again and return the number os rows affected with the operation
	updatedRows := updateWorkflow(uuid, models.WorkflowStatus(models.INSERTED))

	if updatedRows > 0 {
		// since the workflow is inserted again we add the uuid to the end of the queue
		fila.Enqueue(uuid)
	} else {
		log.Fatalf("UUID not found.\n")
	}

	fmt.Printf("Queue size: %v\n", fila.Size())

	// format the message string
	msg := fmt.Sprintf("Workflow updated successfully, readding to the queue. Total rows/record affected %v\n", updatedRows)

	// format the response message
	res := models.Response{
		ID:      uuid,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// DeleteWorkflow delete workflow's detail in the postgres db
func DeleteWorkflow(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the workflowid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id in string to int
	uuid := params["UUID"]

	// delete workflow and return the number os rows affected with the operation
	deletedRows := deleteWorkflow(uuid)

	if deletedRows == 0 {
		log.Fatalf("UUID not found.\n")
	}

	// format the message string
	msg := fmt.Sprintf("Workflow deleted successfully. Total rows/record affected %v\n", deletedRows)

	// format the reponse message
	res := models.Response{
		ID:      uuid,
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

//------------------------- handler functions ----------------
// insert one workflow in the DB
func insertWorkflow(ds models.DataStatus) models.Workflow {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	// returning workflowid will return the id of the inserted workflow
	sqlStatement := `INSERT INTO workflows (Data, Steps) VALUES ($1, $2) RETURNING UUID, status, data, steps`

	// the inserted id will store in this id
	var workflow models.Workflow

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, ds.Data, pq.Array(ds.Steps)).Scan(&workflow.UUID, &workflow.Status, &workflow.Data, pq.Array(&workflow.Steps))

	if err != nil {
		log.Fatalf("Unable to execute the query. %v\n", err)
	}

	fmt.Printf("Inserted a single record %v\n", workflow)

	// return the inserted id
	return workflow
}

// get one workflow from the DB by its workflowid
func getWorkflow(id string) (models.Workflow, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a workflow of models.Workflow type
	var workflow models.Workflow

	// create the select sql query
	sqlStatement := `SELECT * FROM workflows WHERE uuid=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to workflow
	err := row.Scan(&workflow.UUID, &workflow.Status, &workflow.Data, pq.Array(&workflow.Steps))

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return workflow, nil

	case nil:
		return workflow, nil

	default:
		log.Fatalf("Unable to scan the row. %v\n", err)
	}

	// return empty workflow on error
	return workflow, err
}

// get one workflow from the DB by its workflowid
func getAllWorkflows() ([]models.Workflow, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var workflows []models.Workflow

	// create the select sql query
	sqlStatement := `SELECT * FROM workflows`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v\n", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var workflow models.Workflow

		// unmarshal the row object to workflow
		err = rows.Scan(&workflow.UUID, &workflow.Status, &workflow.Data, pq.Array(&workflow.Steps))

		if err != nil {
			log.Fatalf("Unable to scan the row. %v\n", err)
		}

		// append the workflow in the workflows slice
		workflows = append(workflows, workflow)
	}

	// return empty workflow on error
	return workflows, err
}

// update workflow in the DB
func updateWorkflow(id string, status models.WorkflowStatus) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE workflows SET status=$2 WHERE uuid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, status)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v\n", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v\n", err)
	}

	fmt.Printf("Total rows/record affected %v\n", rowsAffected)

	return rowsAffected
}

// delete workflow in the DB
func deleteWorkflow(id string) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM workflows WHERE uuid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v\n", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v\n", err)
	}

	fmt.Printf("Total rows/record affected %v\n", rowsAffected)

	return rowsAffected
}
