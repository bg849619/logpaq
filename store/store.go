package logpaq

import (
	"database/sql"
	_ "embed"
	"path/filepath"
)

//go:embed schema.sql
var schemaSQL string

type Store struct {
	db     *sql.DB
	nodeID string
}

func Open(logID string, dataDir string) (*Store, error) {
	path := filepath.Join(dataDir, logID)
	// Create if the file doesn't exist.
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	if err := migrate(db); err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func migrate(db *sql.DB) error {
	_, err := db.Exec(schemaSQL)
	return err
}
