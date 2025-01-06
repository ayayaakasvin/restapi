package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"restapi/internal/config"

	_ "github.com/lib/pq"
)

// PostgresStorage implements the Storage interface for PostgreSQL
type PostgresStorage struct {
	db *sql.DB
}

// NewPostgresStorage creates a new PostgresStorage
func NewPostgresStorage(connStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &PostgresStorage{db: db}, nil
}

// NewPostgresStorageWithConfig creates a new PostgresStorage with configuration from Config
func NewPostgresStorageWithConfig(cfg config.StorageConfig) (*PostgresStorage, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DatabaseName)
	return NewPostgresStorage(connStr)
}

// Create inserts a new record into the PostgreSQL database
func (ps *PostgresStorage) Create(key string, value interface{}) error {
	_, err := ps.db.Exec("INSERT INTO storage (key, value) VALUES ($1, $2)", key, value)
	return err
}

// Read retrieves a record from the PostgreSQL database
func (ps *PostgresStorage) Read(key string) (interface{}, error) {
	var value interface{}
	err := ps.db.QueryRow("SELECT value FROM storage WHERE key = $1", key).Scan(&value)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("record not found")
		}
		return nil, err
	}
	return value, nil
}

// Update modifies an existing record in the PostgreSQL database
func (ps *PostgresStorage) Update(key string, value interface{}) error {
	_, err := ps.db.Exec("UPDATE storage SET value = $1 WHERE key = $2", value, key)
	return err
}

// Delete removes a record from the PostgreSQL database
func (ps *PostgresStorage) Delete(key string) error {
	_, err := ps.db.Exec("DELETE FROM storage WHERE key = $1", key)
	return err
}

// Ping checks the connection to the PostgreSQL database
func (ps *PostgresStorage) Ping() error {
	return ps.db.Ping()
}

// Close closes the connection to the PostgreSQL database
func (ps *PostgresStorage) Close() error {
	return ps.db.Close()
}

// Reset clears all records from the PostgreSQL database
func (ps *PostgresStorage) Reset() error {
	_, err := ps.db.Exec("DELETE FROM storage")
	return err
}
