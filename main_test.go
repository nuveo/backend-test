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
	a         main.App
	luser     = "postgres"
	lpassword = "1234"
	ldbname   = "nuveo"
)

const createTable = `CREATE TABLE IF NOT EXISTS workflows (
	uuid SERIAL,
	status status_t DEFAULT 'inserted',
	data JSONB NOT NULL,
	steps text[],
	CONSTRAINT workflows_pkey PRIMARY KEY (uuid)
);`

func TestMain(m *testing.M) {
	a = main.App{}
	a.InitDatabase(
		luser,
		lpassword,
		ldbname)

	if _, err := a.DB.Exec(createTable); err != nil {
		log.Fatal(err)
	}

	code := m.Run()

	a.DB.Exec("DELETE FROM workflows")
	a.DB.Exec("ALTER SEQUENCE workflows_uuid_seq RESTART WITH 1")

	os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
	a.DB.Exec("DELETE FROM workflows")
	a.DB.Exec("ALTER SEQUENCE workflows_id_seq RESTART WITH 1")

	req, _ := http.NewRequest("GET", "/workflows", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestGetNonExistentProduct(t *testing.T) {
	a.DB.Exec("DELETE FROM workflows")
	a.DB.Exec("ALTER SEQUENCE workflows_uuid_seq RESTART WITH 1")

	req, _ := http.NewRequest("GET", "/workflow/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Workflow not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Workflow not found'. Got '%s'", m["error"])
	}
}

func TestCreateProduct(t *testing.T) {
	a.DB.Exec("DELETE FROM workflows")
	a.DB.Exec("ALTER SEQUENCE workflows_uuid_seq RESTART WITH 1")

	payload := []byte(`{"uuid":1, "status":"inserted",'{"teste1": "teste1"}', '{"hello"}'}`)

	req, _ := http.NewRequest("POST", "/workflow", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

}

func TestGetProduct(t *testing.T) {
	a.DB.Exec("DELETE FROM workflows")
	a.DB.Exec("ALTER SEQUENCE workflows_uuid_seq RESTART WITH 1")

	count := 1

	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO workflows(status, data, steps) VALUES($1, $2, $3)", "inserted", "{\"teste1\": \"teste1\"}", "{\"hello\"}")
	}

	req, _ := http.NewRequest("GET", "/workflow/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateProduct(t *testing.T) {
	a.DB.Exec("DELETE FROM workflows")
	a.DB.Exec("ALTER SEQUENCE workflows_uuid_seq RESTART WITH 1")

	count := 1

	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO workflows(status, data, steps) VALUES($1, $2, $3)", "inserted", "{\"teste1\": \"teste1\"}", "{\"hello\"}")
	}

	req, _ := http.NewRequest("GET", "/workflow/1", nil)
	response := executeRequest(req)
	var originalProduct map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalProduct)

	payload := []byte(`{"uuid":1, "status":"inserted",'{"teste2": "teste2"}', '{"bye"}'}`)

	req, _ = http.NewRequest("PUT", "/workflows/1", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["uuid"] != originalProduct["uuid"] {
		t.Errorf("Expected the uuid to remain the same (%v). Got %v", originalProduct["uuid"], m["uuid"])
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
