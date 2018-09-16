package main

import (
	"database/sql"
	"errors"
)

// workflow reflects the attributes from workflow's table.
type workflow struct {
	UUID   int      `json:"UUID"`
	Status int      `json:"status"`
	Data   float64  `json:"data"`
	Steps  []string `json:"steps"`
}

func (w *workflow) getWorkflow(db *sql.DB) error {
	return db.QueryRow("SELECT status, data, steps FROM workflows WHERE uuid=$1",
		w.UUID).Scan(&w.Status, &w.Data, &w.Steps)
}

func (w *workflow) insertWorkflow(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO workflows(status, data, steps) VALUES($1, $2, $3) RETURNING uuid",
		w.Status, w.Data, w.Steps).Scan(&w.UUID)
	if err != nil {
		return err
	}

	return nil
}

func (w *workflow) updateWorkflow(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE workflows SET status=$1, data=$2, steps=$3, WHERE uuid=$4",
			w.Status, w.Data, w.Steps, w.UUID)
	return err
}

func getWorkflows(db *sql.DB, start, count int) ([]workflow, error) {
	rows, err := db.Query(
		"SELECT uuid, status, data, steps FROM workflows LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	workflows := []workflow{}

	for rows.Next() {
		var w workflow
		if err := rows.Scan(&w.UUID, &w.Status, &w.Data, &w.Steps); err != nil {
			return nil, err
		}
		workflows = append(workflows, w)
	}

	return workflows, nil
}

func (w *workflow) consumeWorkflow(db *sql.DB) error {
	return errors.New("Not implemented")
}
