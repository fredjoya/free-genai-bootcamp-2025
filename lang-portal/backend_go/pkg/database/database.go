package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// InitDB initializes the database connection and sets up the schema
func InitDB() error {
	// Get the current working directory
	dir := "data"

	// Create the database directory if it doesn't exist
	dbDir := filepath.Join(dir)
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

	// Initialize schema
	if err := initSchema(); err != nil {
		return fmt.Errorf("error initializing schema: %v", err)
	}

	// Seed test data
	if err := seedTestData(); err != nil {
		return fmt.Errorf("error seeding test data: %v", err)
	}

	log.Println("Database initialized successfully")
	return nil
}

// initSchema initializes the database schema
func initSchema() error {
	schema, err := ioutil.ReadFile("schema.sql")
	if err != nil {
		return fmt.Errorf("error reading schema file: %v", err)
	}

	// Execute the schema
	if _, err := DB.Exec(string(schema)); err != nil {
		return fmt.Errorf("error executing schema: %v", err)
	}

	return nil
}

// seedTestData seeds the database with test data
func seedTestData() error {
	seed, err := ioutil.ReadFile("seed.sql")
	if err != nil {
		return fmt.Errorf("error reading seed file: %v", err)
	}

	// Execute the seed data
	if _, err := DB.Exec(string(seed)); err != nil {
		return fmt.Errorf("error executing seed data: %v", err)
	}

	return nil
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	return DB
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

// WithTransaction executes a function within a transaction
func WithTransaction(fn func(*sql.Tx) error) error {
	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
} 