package middlewares

import (
	"backend-test/helpers"
	"backend-test/models"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type response struct {
	ID      string `json: "id,string,omitempty"`
	Message string `json:"message,omitempty"`
}

var (
	queue helpers.ElementQueue
)

// Create connection with postgres
func createConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error when loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("\nSuccessfully connected!")

	queue.Create()

	return db
}

// CreateWorkflow create a new workflow
func CreateWorkflow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var workflow models.Workflow

	err := json.NewDecoder(r.Body).Decode(&workflow)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err.Error())
	}

	insertID := insertWorkflow(workflow)

	res := response{
		ID:      insertID,
		Message: "Workflow created successfully.",
	}

	json.NewEncoder(w).Encode(res)
}

// GetAllWorkflow get all workflows from database
func GetAllWorkflow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	workflows, err := getAllWorkflows()

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}

	json.NewEncoder(w).Encode(workflows)
}

// ConsumeWorkflow get a workflow from your uuid
func ConsumeWorkflow(w http.ResponseWriter, r *http.Request) {

	if queue.IsEmpty() {
		log.Fatal("Unable dequeue")
	}

	uuid := queue.Dequeue()

	workflow, err := getWorkflow(uuid)

	if err != nil {
		log.Fatalf("Unable to get workflow. %v", err)
	}

	updateWorkflow(uuid, workflow)

	generateCsv(workflow)

	json.NewEncoder(w).Encode(workflow)

}

// UpdateWorkflow update a workflow from your uuid
func UpdateWorkflow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	uuid := params["uuid"]

	var workflow models.Workflow

	json.NewDecoder(r.Body).Decode(&workflow)

	updateRows := updateWorkflow(string(uuid), workflow)

	msg := fmt.Sprintf("Workflow updated sucessfully. %v", updateRows)

	res := response{
		ID:      string(uuid),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)

}

// DeletWorkflow delete a workflow from your uuid
func DeleteWorkflow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	uuid, err := strconv.Atoi(params["uuid"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	deletedWorkflow := deleteWorkflow(string(uuid))

	msg := fmt.Sprintf("Workflow deleted sucessfully. %v", deletedWorkflow)

	res := response{
		ID:      string(uuid),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

// Insert a workflow in database
func insertWorkflow(workflow models.Workflow) string {

	log.Println("Creating a new workflow")

	db := createConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO workflow (data, steps) VALUES ($1, $2) RETURNING UUID`

	var uuid string

	err := db.QueryRow(sqlStatement, workflow.Data, pq.Array(workflow.Steps)).Scan(&uuid)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", uuid)

	queue.Enqueue(uuid)

	return uuid
}

// Get one workflow from uuid
func getWorkflow(uuid string) (models.Workflow, error) {
	db := createConnection()

	defer db.Close()

	var workflow models.Workflow

	sqlStatement := `SELECT * FROM workflow WHERE uuid=$1`

	row := db.QueryRow(sqlStatement, uuid)

	// unmarshal the row object to user
	err := row.Scan(&workflow.UUID, &workflow.Status, &workflow.Data, pq.Array(&workflow.Steps))

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return workflow, nil
	case nil:
		return workflow, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return workflow, err
}

// Get all workflows from the DB.
func getAllWorkflows() ([]models.Workflow, error) {
	db := createConnection()

	defer db.Close()

	var workflows []models.Workflow

	sqlStatement := `SELECT * FROM workflow`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var workflow models.Workflow

		err = rows.Scan(&workflow.UUID, &workflow.Status, &workflow.Data, pq.Array(&workflow.Steps))

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		workflows = append(workflows, workflow)

	}

	return workflows, err
}

// Update user in the DB
func updateWorkflow(uuid string, workflow models.Workflow) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `UPDATE workflow SET status=$2 WHERE uuid=$1`

	res, err := db.Exec(sqlStatement, uuid, workflow.Status)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v\n", rowsAffected)

	return rowsAffected
}

// Delete workflow in database
func deleteWorkflow(uuid string) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM workflow WHERE uuid=$1`

	res, err := db.Exec(sqlStatement, uuid)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// Generate workflow file .csv
func generateCsv(workflow models.Workflow) {
	path, _ := filepath.Abs("./")

	folder := filepath.Join(".", "spreadsheets")

	os.MkdirAll(folder, os.ModePerm)
	file, err := os.Create(fmt.Sprintf("%s/spreadsheets/%s.csv", path, workflow.UUID))

	if err != nil {
		log.Fatalf("Error while create files. %v", err)

	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var list []string

	list = append(list, string(workflow.Data))

	log.Printf("generate a CSV file with workflow data")
}
