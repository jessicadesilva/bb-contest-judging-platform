package models

import (
	"database/sql"
	"errors"
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

func (m *CompetitorModel) Get(CompetitorID int) (Competitor, error) {
	// SQL statement to execute.
	stmt := `SELECT * FROM judgingplatform.competitors
		WHERE id = ?`

	// Execute the SQL statement using the QueryRow() method
	// on the connection pool. This returns a pointer to a sql.Row
	// object which holds the result from the database.
	row := m.DB.QueryRow(stmt, CompetitorID)

	// Initialize a new zeroed Competitor struct.
	var c Competitor

	// Copy the values from each field in sql.Row to
	// the corresponding filed in the Competitor struct
	// using row.Scan().
	err := row.Scan(&c.ID, &c.Name, &c.Location)

	if err != nil {
		// If the query returns no rows, we get a sql.ErrNoRows
		// error wich we can check for using errors.Is() to return
		// a ErrNoRecord error instead.
		if errors.Is(err, sql.ErrNoRows) {
			return Competitor{}, ErrNoRecord
		} else {
			return Competitor{}, err
		}
	}

	// If everything went OK, return the filled Competitor struct.
	return c, nil
}
