package dbmanager

import (
	"database/sql"
	"os"
)

// The InitDB function initializes a connection to a SQLite database and creates a table if it doesn't
// exist.
func InitDB() (*sql.DB, error) {
	connDB, err := sql.Open("sqlite3", "./databases/post.db")
	if err != nil {
		return nil, err
	}

	err = createTable(connDB)
	if err != nil {
		return nil, err
	}

	return connDB, nil
}

// The function `createTable` reads a SQL file and executes its content to create a table in the
// database.
func createTable(db *sql.DB) error {
	content, err := os.ReadFile("./databases/sqlRequests/createTable.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(content))
	if err != nil {
		return err
	}

	return nil
}
