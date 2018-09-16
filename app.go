package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	a.Router.HandleFunc("/workflows", a.createWorkflow).Methods("POST")
	a.Router.HandleFunc("/workflows/{id:[0-9]+}", a.Workflow).Methods("GET")
	a.Router.HandleFunc("/workflows/{id:[0-9]+}", a.updateWorkflow).Methods("PATCH")
}

// WorkflowRequest reflects the attributes from workflow's table.
type WorkflowRequest struct {
	UUID   int             `json:"uuid"`
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
	Steps  string          `json:"steps"`
}

// Database starts a connection with the database.
func (a *App) Database(user, password, dbname string) error {
	credential := fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open(db, credential)
	if err != nil {
		return err
	}

	return nil
}

// Routes starts a new router and its endpoints.
func (a *App) Routes() {

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

// Run runs the application in the address 'addr'.
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

// Workflow returns the selected ID.
func (a *App) Workflow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("Returning workflow " + vars["id"])

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorReply(w, http.StatusBadRequest, "Invalid workflow ID")
		return
	}

	wf := Workflow{UUID: id}
	if err := wf.getWorkflow(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			errorReply(w, http.StatusNotFound, "Workflow not found")
		default:
			errorReply(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	reply(w, http.StatusOK, wf)
}

// Workflows returns all workflows from database.
func (a *App) Workflows(w http.ResponseWriter, r *http.Request) {
	log.Println("Returning all workflows")

	count, _ := strconv.Atoi(r.FormValue("count"))
	if count > 10 || count < 1 {
		count = 10
	}

	start, _ := strconv.Atoi(r.FormValue("start"))
	if start < 0 {
		start = 0
	}

	workflows, err := getWorkflows(a.DB, start, count)
	if err != nil {
		errorReply(w, http.StatusInternalServerError, err.Error())
		return
	}

	reply(w, http.StatusOK, workflows)
}

// createWorkflow creates a new workflow received from payload data.
func (a *App) createWorkflow(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating new workflow")

	var wfr WorkflowRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&wfr); err != nil {
		errorReply(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	wf := Workflow{
		UUID:   wfr.UUID,
		Status: wfr.Status,
		Data:   string(wfr.Data),
		Steps:  wfr.Steps,
	}

	queue.Enqueue(wf)

	if err := wf.insertWorkflow(a.DB); err != nil {
		errorReply(w, http.StatusInternalServerError, err.Error())
		return
	}

	reply(w, http.StatusCreated, wf.UUID)
}

// updateWorkflow updates selected workflow with received ID.
func (a *App) updateWorkflow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("Updating workflow " + vars["id"])

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		errorReply(w, http.StatusBadRequest, "Invalid workflow UUID")
		return
	}

	var wf Workflow
	wf.UUID = id
	if err := wf.getWorkflow(a.DB); err != nil {
		errorReply(w, http.StatusInternalServerError, err.Error())
		return
	}

	if wf.Data == "" {
		errorReply(w, http.StatusInternalServerError, "Workflow not found")
		return
	}
	wf.Status = "consumed"

	if err := wf.updateWorkflow(a.DB); err != nil {
		errorReply(w, http.StatusInternalServerError, err.Error())
		return
	}

	reply(w, http.StatusOK, wf)
}

// reply returns request.
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
