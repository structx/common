// Package kv provides a key-value store interface.
package kv

import (
	"errors"
	"fmt"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

// NewLevelDB returns a new LevelDB instance.
func NewLevelDB() (*leveldb.DB, error) {

	dbPath := os.Getenv("LEVELDB_PATH")
	if dbPath == "" {
		return nil, errors.New("LEVELDB_PATH is not set")
	}

	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open leveldb: %w", err)
	}

	return db, nil
}
