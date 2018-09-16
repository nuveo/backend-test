package main_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"." // main
)

var (
	a        main.App
	user     = "postgres"
	password = "1234"
	dbname   = "nuveo"
)

const createTable = `CREATE TABLE IF NOT EXISTS workflows (
	uuid UUID,
	status status_t DEFAULT 'inserted',
	data JSONB NOT NULL,
	steps text[],
	CONSTRAINT workflows_pkey PRIMARY KEY (uuid)
);`

func TestMain(m *testing.M) {
	a = main.App{}
	a.Connect(user, password, dbname)
	a.Routes()

	if _, err := a.DB.Exec(createTable); err != nil {
		log.Fatal(err)
	}

	code := m.Run()

	clearDatabase()

	os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
	clearDatabase()

	req, err := http.NewRequest("GET", "/workflows", nil)
	if err != nil {
		t.Errorf("Failed on creating a new GET request: %v", err)
	}

	response := executeRequest(t, req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentWorkflow(t *testing.T) {
	clearDatabase()

	uuid := "a6406644-414b-4e11-bdf5-c1438928dc14"
	req, err := http.NewRequest("GET", "/workflows/"+uuid, nil)
	if err != nil {
		t.Errorf("Failed on creating a new GET request: %v", err)
	}

	response := executeRequest(t, req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Workflow not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Workflow not found'. Got '%s'", m["error"])
	}
}

func TestCreateProduct(t *testing.T) {
	clearDatabase()

	payload := []byte(`{"status": "inserted","data":"{'teste': 'Teste'}", "steps":"["go", "eh","vida"]"}`)

	req, err := http.NewRequest("POST", "/workflows", bytes.NewBuffer(payload))
	if err != nil {
		t.Errorf("Failed on creating a new POST request: %v", err)
	}

	response := executeRequest(t, req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	// verificar dados retornados
}

// LAUREN corrigir recebendo UUID
func TestUpdateProduct(t *testing.T) {
	clearDatabase()

	count := 1

	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO workflows(status, data, steps) VALUES($1, $2, $3)", "inserted", "{\"teste1\": \"teste1\"}", "['go', 'eh','vida']")
	}

	req, err := http.NewRequest("GET", "/workflow/1", nil)
	if err != nil {
		t.Errorf("Failed on creating a new GET request: %v", err)
	}

	response := executeRequest(t, req)

	var originalProduct map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalProduct)

	payload := []byte(`{"status": "consumed"}`)

	req, err = http.NewRequest("PATCH", "/workflows/1", bytes.NewBuffer(payload))
	if err != nil {
		t.Errorf("Failed on creating a new PATCH request: %v", err)
	}

	response = executeRequest(t, req)

	checkResponseCode(t, http.StatusOK, response.Code)

	// verificar campos alterados
}

func executeRequest(t *testing.T, r *http.Request) *httptest.ResponseRecorder {
	t.Log("Executing request")

	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, r)

	return recorder
}

func checkResponseCode(t *testing.T, expected, actual int) {
	t.Log("Checking response code")
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func clearDatabase() {
	a.DB.Exec("DELETE FROM workflows")
	a.DB.Exec("ALTER SEQUENCE workflows_uuid_seq RESTART WITH 1")
}
