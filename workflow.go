package main

import (
	"database/sql"
)

// Workflow reflects the attributes from Workflow's table.
type Workflow struct {
	UUID   string
	Status string
	Data   string
	Steps  string
}

// Get selects workflow from database by ID.
func (w *Workflow) Get(db *sql.DB) error {
	return db.QueryRow("SELECT status, data, steps FROM workflows WHERE uuid=$1",
		w.UUID).Scan(&w.Status, &w.Data, &w.Steps)
}

// Insert creates a new workflow in the database.
func (w *Workflow) Insert(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO workflows(status, data, steps) VALUES($1, $2, $3) RETURNING uuid",
		w.Status, w.Data, w.Steps).Scan(&w.UUID)
	if err != nil {
		return err
	}

	return nil
}

// Update changes workflow status.
func (w *Workflow) Update(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE workflows SET status=$1 WHERE uuid=$2",
			w.Status, w.UUID)
	return err
}

// Workflows returns all workflows from database.
func Workflows(db *sql.DB, start, count int) ([]Workflow, error) {
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
