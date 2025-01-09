package storage

import (
	"database/sql"
	"fmt"

	"restapi/internal/config"
	"restapi/internal/lib/hashtool"
	"restapi/internal/models"
	migrationTool "restapi/internal/storage/migrationTool"

	_ "github.com/lib/pq"
)

// PostgresStorage implements the Storage interface for PostgreSQL
type PostgresStorage struct {
	db *sql.DB
}

// NewPostgresStorage creates a new PostgresStorage
func NewPostgresStorage(cfg *config.Config) (*PostgresStorage, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DatabaseName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Run migration tool
	if err := migrationTool.ExecuteMigrationTool(
		db, cfg.MigrationPath, migrationTool.UpMigration); err != nil {
		return nil, fmt.Errorf("failed to run migration tool: %w", err)
	}

	return &PostgresStorage{db: db}, nil
}

// SaveUser inserts a new user record into the PostgreSQL database
func (ps *PostgresStorage) SaveUser(username, password string) error {
	hashedPassword, err := hashtool.BcryptHashing(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	stmt, err := ps.db.Prepare("INSERT INTO users (username, password) VALUES ($1, $2)")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

// GetUserByID retrieves a record from the PostgreSQL database by key
func (ps *PostgresStorage) GetUserByID(id int64) (*models.User, error) {
	stmt, err := ps.db.Prepare("SELECT id, username, password, created_at FROM users WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(id).Scan(&user.ID, &user.UserName, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}

	return &user, nil
}

// UsernameExists checks if a record with the given username exists in the PostgreSQL database
func (ps *PostgresStorage) UsernameExists(name string) (bool, error) {
	stmt, err := ps.db.Prepare("SELECT 1 FROM users WHERE username = $1")
	if err != nil {
		return false, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	var exists bool
	err = stmt.QueryRow(name).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to execute statement: %w", err)
	}

	return exists, nil
}

// UpdateUser updates a record in the PostgreSQL database
func (ps *PostgresStorage) UpdateUserPassword(id int64, password string) error {
	stmt, err := ps.db.Prepare("UPDATE users SET password = $1 WHERE id = $2")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	var hashedPassword string 
	if hashedPassword, err = hashtool.BcryptHashing(password); err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	_, err = stmt.Exec(hashedPassword, id)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

// DeleteUser deletes a record from the PostgreSQL database
func (ps *PostgresStorage) DeleteUser(id int64) error {
	stmt, err := ps.db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

// //SaveTask inserts a new task record into the PostgreSQL database
// func (ps *PostgresStorage) SaveTask(userId int64, content string) error {
// 	stmt, err := ps.db.Prepare("INSERT INTO tasks (user_id, content) VALUES ($1, $2)")
// 	if err != nil {
// 		return fmt.Errorf("failed to prepare statement: %w", err)
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(userId, content)
// }

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
	err := migrationTool.ExecuteMigrationTool(ps.db, "migrations", migrationTool.ResetMigration)
	if err != nil {
		return fmt.Errorf("failed to reset database: %w", err)
	}

	return nil
}