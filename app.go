package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App contains the application's connections and dependencies.
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// initializeRoutes provides the application's endpoints and handles requests.
func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/workflows", a.Workflows).Methods("GET")
	a.Router.HandleFunc("/workflows", a.CreateWorkflow).Methods("POST")
	a.Router.HandleFunc("/workflows/{UUID}", a.UpdateWorkflow).Methods("PATCH")
	a.Router.HandleFunc("/workflows/consume", a.ConsumeWorkflow).Methods("GET")
}

// Connect starts a connection with the database.
func (a *App) Connect(user, password, dbname string) error {
	credentials := fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", credentials)
	if err != nil {
		return err
	}

	return nil
}

// Prepare installs missing table and/or extensions and prepares database.
func (a *App) Prepare() error {
	if _, err := a.DB.Exec(createEnum); err != nil {
		return err
	}
	if _, err := a.DB.Exec(createTable); err != nil {
		return err
	}
	return nil
}

// Routes starts a new router and its endpoints.
func (a *App) Routes() {

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

// Run runs the application in addr.
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

// Workflows returns all workflows from database.
func (a *App) Workflows(w http.ResponseWriter, r *http.Request) {
	log.Println("Returning all workflows")

	workflows, err := Workflows(a.DB)
	if err != nil {
		errorReply(w, http.StatusInternalServerError, err.Error())
		return
	}

	reply(w, http.StatusOK, workflows)
}

// CreateWorkflow creates a new workflow received from payload data.
func (a *App) CreateWorkflow(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating new workflow")

	type decoded struct {
		Status string          `json:"status"`
		Data   json.RawMessage `json:"data"`
		Steps  string          `json:"steps"`
	}

	var d decoded
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&d); err != nil {
		errorReply(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	workflow := Workflow{
		Status: d.Status,
		Data:   string(d.Data),
		Steps:  d.Steps,
	}

	if err := workflow.Insert(a.DB); err != nil {
		errorReply(w, http.StatusInternalServerError, err.Error())
		return
	}

	queue.Enqueue(workflow)

	reply(w, http.StatusCreated, workflow)
}

// UpdateWorkflow updates selected workflow with received ID.
func (a *App) UpdateWorkflow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("Updating workflow " + vars["UUID"])

	var workflow Workflow
	workflow.UUID = vars["UUID"]
	if err := workflow.Get(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			errorReply(w, http.StatusNotFound, "Workflow not found")
		default:
			errorReply(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	queue.Remove(workflow.UUID)

	if workflow.Status == "consumed" {
		errorReply(w, http.StatusInternalServerError, "Workflow lala was already consumed")
		return
	}

	workflow.Status = "consumed"

	if err := workflow.Update(a.DB); err != nil {
		errorReply(w, http.StatusInternalServerError, err.Error())
		return
	}

	reply(w, http.StatusOK, workflow)
}

// ConsumeWorkflow consumes workflows from queue.
func (a *App) ConsumeWorkflow(w http.ResponseWriter, r *http.Request) {
	path, _ := filepath.Abs("./")

	folder := filepath.Join(".", "data")
	os.MkdirAll(folder, os.ModePerm)

	if queue.IsEmpty() {
		errorReply(w, http.StatusInternalServerError, "Empty queue")
		return
	}

	workflow := queue.Dequeue()
	uuid := workflow.UUID
	fileName := uuid + ".csv"

	file, err := os.Create(path + "/data/" + fileName)
	if err != nil {
		errorReply(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var data []string
	data = append(data, workflow.Data)

	if err := writer.Write(data); err != nil {
		errorReply(w, http.StatusInternalServerError, err.Error())
		return
	}

	workflow.Status = "consumed"
	if err := workflow.Update(a.DB); err != nil {
		errorReply(w, http.StatusInternalServerError, err.Error())
		return
	}

	reply(w, http.StatusOK, "CSV file "+fileName+" generated successfully")
}

// reply returns request with header.
func reply(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// errorReply returns request with an error message.
func errorReply(w http.ResponseWriter, code int, message string) {
	reply(w, code, map[string]string{"error": message})
}
