package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB() error {
	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %v", err)
	}

	// Create the database directory if it doesn't exist
	dbDir := filepath.Join(dir, "data")
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return fmt.Errorf("failed to create database directory: %v", err)
	}

	// Open the database file
	dbPath := filepath.Join(dbDir, "lang_portal.db")
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	DB = db
	log.Println("Database initialized successfully")
	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
} 