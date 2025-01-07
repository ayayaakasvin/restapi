package migrationtool

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"strings"

	"restapi/internal/config"

	_ "github.com/lib/pq"
)

// MigrationTool is a tool for migrating and setuping the database
func MigrationToolExecutable(cfg *config.Config) error {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DatabaseName))
	if err != nil {
		return err
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		return err
	}

	// we should parse sql script from the file and execute it
	sqlScripts, err := parseSQLScriptFromFile(cfg.MigrationPath)

	if err != nil {
		return fmt.Errorf("failed to parse sql script from file: %w", err)
	}

	// for _, script := range sqlScripts {
	// 	fmt.Println(script) // for debug
	// }

	for _, script := range sqlScripts {
		if err := executeSQLScript(db, script); err != nil {
			return fmt.Errorf("failed to execute sql script: %w", err)
		}
	}

	return nil
}

func parseSQLScriptFromFile(filePath string) ([]string, error) {
	if filePath == "" {
		return nil, fmt.Errorf("file path is empty")
	}

	// we should read the directory and return the content
	entries, err := os.ReadDir(filePath)
	if err != nil {
		return nil, err
	}

	if len(entries) == 0 {
		return nil, fmt.Errorf("directory is empty")
	}

	var (
		sqlScripts []string
	)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if isUpMigration(entry.Name()) {
			content, err := readSQLScriptFromFile(filePath + "/" + entry.Name())
			if err != nil {
				return nil, err
			}

			sqlScripts = append(sqlScripts, content)
		}
	}

	return sqlScripts, nil
}

// function to execute the sql script
func executeSQLScript(db *sql.DB, script string) error {
	stmt, err := db.Prepare(script)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}

	defer stmt.Close()
	
	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	return nil
}

// to read the content of the file
func readSQLScriptFromFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// function to check if file is .up.sql format
func isUpMigration(fileName string) bool {
	return strings.HasSuffix(fileName, ".up.sql")
}
