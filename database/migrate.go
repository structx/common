package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	migrator "github.com/golang-migrate/migrate/v4/database/pgx/v5"

	_ "github.com/golang-migrate/migrate/v4/source/file" // file driver
)

// Migrate runs database migrations
func Migrate() error {

	dsn := os.Getenv("DB_MIGRATIONS_DSN")
	if dsn == "" {
		return errors.New("$DB_MIGRATIONS_DSN must be set")
	}

	db := &migrator.Postgres{}
	driver, err := db.Open(dsn)
	if err != nil {
		return fmt.Errorf("failed to open migrator driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", os.Getenv("DB_MIGRATIONS_DIR")), "pgx", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrator: %w", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	return nil
}
