package main

import (
	"database/sql"
	"errors"
)

// Workflow reflects the attributes from Workflow's table.
type Workflow struct {
	UUID   int    `json:"uuid"`
	Status string `json:"status"`
	Data   string `json:"data"`
	Steps  string `json:"steps"`
}

func (w *Workflow) getWorkflow(db *sql.DB) error {
	return db.QueryRow("SELECT status, data, steps FROM workflows WHERE uuid=$1",
		w.UUID).Scan(&w.Status, &w.Data, &w.Steps)
}

func (w *Workflow) insertWorkflow(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO workflows(status, data, steps) VALUES($1, $2, $3) RETURNING uuid",
		w.Status, w.Data, w.Steps).Scan(&w.UUID)
	if err != nil {
		return err
	}

	return nil
}

func (w *Workflow) updateWorkflow(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE workflows SET status=$1 WHERE uuid=$2",
			w.Status, w.UUID)
	return err
}

func getWorkflows(db *sql.DB, start, count int) ([]Workflow, error) {
	rows, err := db.Query(
		"SELECT uuid, status, data, steps FROM workflows LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	workflows := []Workflow{}

	for rows.Next() {
		var w Workflow
		if err := rows.Scan(&w.UUID, &w.Status, &w.Data, &w.Steps); err != nil {
			return nil, err
		}
		workflows = append(workflows, w)
	}

	return workflows, nil
}

func (w *Workflow) consumeWorkflow(db *sql.DB) error {
	return errors.New("Not implemented")
}
