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

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/workflows", a.getWorkflows).Methods("GET")
	a.Router.HandleFunc("/workflow", a.createWorkflow).Methods("POST")
	a.Router.HandleFunc("/workflow/{id:[0-9]+}", a.getWorkflow).Methods("GET")
	a.Router.HandleFunc("/workflow/{id:[0-9]+}", a.updateWorkflow).Methods("PATCH")
}

// InitDatabase starts a connection with the 'dbname' database.
func (a *App) InitDatabase(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

// Run runs the application in the address 'addr'.
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func (a *App) getWorkflow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["uuid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid workflow ID")
		return
	}

	wf := workflow{UUID: id}
	if err := wf.getWorkflow(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Workflow not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, wf)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) getWorkflows(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	workflows, err := getWorkflows(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, workflows)
}

func (a *App) createWorkflow(w http.ResponseWriter, r *http.Request) {
	var wf workflow
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&wf); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := wf.insertWorkflow(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, wf)
}

func (a *App) updateWorkflow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid workflow UUID")
		return
	}

	var wf workflow
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&wf); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	wf.UUID = id

	if err := wf.updateWorkflow(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, wf)
}
