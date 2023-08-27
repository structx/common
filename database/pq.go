package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	// Import the pq Postgres driver
	_ "github.com/lib/pq"
)

// NewPQ returns a new connection to a PostgreSQL database
func NewPQ() (*sql.DB, error) {

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		return nil, errors.New("DB_DSN must be set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	return db, nil
}
