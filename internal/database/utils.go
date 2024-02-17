package database

// run sql file in the database

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func RunSQLFile(db *sql.DB, filename string) error {
	b, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	queries := strings.Split(string(b), ";")

	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}

		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("error executing query %s: %v", query, err)
		}
	}

	return nil
}
