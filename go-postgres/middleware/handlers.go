package middleware

import (
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"go-postgres/models" // models package where Workflow schema is defined
	"log"
	"net/http" // used to access the request and response object of the api
	"os"       // used to read the environment variable

	// package used to covert string into int type
	"github.com/gorilla/mux" // used to get the params from the route

	"github.com/joho/godotenv" // package used to read the .env file
	_ "github.com/lib/pq"      // postgres golang driver
)

// response format
type response struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

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
	var workflow models.Workflow

	// decode the json request to workflow
	err := json.NewDecoder(r.Body).Decode(&workflow)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call insert workflow function and pass the workflow
	insertID := insertWorkflow(workflow)

	// format a response object
	res := response{
		ID:      insertID,
		Message: "Workflow created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// GetWorkflow will return a single workflow by its id
func GetWorkflow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the workflowid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := params["id"]

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the getWorkflow function with workflow id to retrieve a single workflow
	workflow, err := getWorkflow(string(id))

	if err != nil {
		log.Fatalf("Unable to get workflow. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(workflow)
}

// GetAllWorkflow will return all the workflows
func GetAllWorkflow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the workflows in the db
	workflows, err := getAllWorkflows()

	if err != nil {
		log.Fatalf("Unable to get all workflow. %v", err)
	}

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

	// convert the id type from string to int
	id, err := params["id"]

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// create an empty workflow of type models.Workflow
	var workflow models.WorkflowStatus

	// decode the json request to workflow
	err = json.NewDecoder(r.Body).Decode(&workflow)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call update workflow to update the workflow
	updatedRows := updateWorkflow(string(id), workflow)

	// format the message string
	msg := fmt.Sprintf("Workflow updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := response{
		ID:      string(id),
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
	id, err := params["id"]

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the deleteWorkflow, convert the int to int64
	deletedRows := deleteWorkflow(string(id))

	// format the message string
	msg := fmt.Sprintf("Workflow updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := response{
		ID:      string(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

//------------------------- handler functions ----------------
// insert one workflow in the DB
func insertWorkflow(workflow models.Workflow) string {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	// returning workflowid will return the id of the inserted workflow
	sqlStatement := `INSERT INTO workflows (Data, Steps) VALUES ($1, $2) RETURNING UUID`

	// the inserted id will store in this id
	var id string

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, workflow.Data, workflow.Steps).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
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
	err := row.Scan(&workflow.UUID, &workflow.Status, &workflow.Data, &workflow.Steps)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return workflow, nil
	case nil:
		return workflow, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
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
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var workflow models.Workflow

		// unmarshal the row object to workflow
		err = rows.Scan(&workflow.UUID, &workflow.Status, &workflow.Data, &workflow.Steps)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
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
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

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
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
