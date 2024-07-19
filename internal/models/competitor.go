package models

import (
	"database/sql"
)

// Define a Competitor type to hold the data for an individual competitor.
type Competitor struct {
	ID       int
	Name     string
	Location string
}

// Define a CompetitorModel which wraps a sql.DB connection pool.
type CompetitorModel struct {
	DB *sql.DB
}

// This will insert a new competitor into the database.
func (m *CompetitorModel) Insert(Name string, Location string) (int, error) {
	// SQL statement to execute.
	// Use placeholder parameters to avoid unwanted SQL injection.
	stmt := `INSERT INTO competitors (name, location)
	VALUES (?, ?)`

	// Execute the statement using the Exec method on the embedded connection pool.
	// This method returns a sql.Result type, which contains basic information about
	// what happened when the statement was executed.
	result, err := m.DB.Exec(stmt, Name, Location)
	if err != nil {
		return 0, err
	}

	// Use the LastInsertId() method to get the ID of our newly inserted record.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Convert ID from int64 to int type before returning.
	return int(id), nil
}
