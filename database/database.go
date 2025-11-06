package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Init initializes the global DB connection.
func Init() {
	host := "localhost" //"database" // Docker service name
	user := "otochopego"
	pass := "gotochope"
	name := "otochope"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", user, pass, host, name)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatalf("Database unreachable: %v", err)
	}

	db = conn
	log.Printf("Connected to %s database at %s", name, host)
}

// Execute runs any query that doesn’t return rows (INSERT, UPDATE, DELETE, etc.)
func Execute(query string, args ...any) (sql.Result, error) {
	if db == nil {
		return nil, fmt.Errorf("database not initialized — call database.Init() first")
	}
	return db.Exec(query, args...)
}

// Query runs a SELECT and returns the rows.
func Query(query string, args ...any) (*sql.Rows, error) {
	if db == nil {
		return nil, fmt.Errorf("database not initialized — call database.Init() first")
	}
	return db.Query(query, args...)
}

// Close safely closes the DB connection
func Close() {
	if db != nil {
		_ = db.Close()
	}
}
