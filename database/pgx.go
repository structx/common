// Package database contains the database models and logic for the application
package database

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPGXPool returns a new pgxpool.Pool
func NewPGXPool(ctx context.Context) (*pgxpool.Pool, error) {

	// get the DSN from the environment
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		return nil, errors.New("DB_DSN is not set")
	}

	// parse the DSN
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.ParseConfig: %w", err)
	}

	// create the pool
	p, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.NewWithConfig: %w", err)
	}

	return p, nil
}
