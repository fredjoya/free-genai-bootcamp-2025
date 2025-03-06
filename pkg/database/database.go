package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

// Mock data storage to bypass SQLite issues temporarily
var (
	MockWords = []map[string]interface{}{
		{"id": 1, "arabic": "مرحبا", "transliteration": "marhaba", "english": "hello", "correct_count": 5, "wrong_count": 2},
		{"id": 2, "arabic": "شكرا", "transliteration": "shukran", "english": "thank you", "correct_count": 3, "wrong_count": 1},
	}
	MockGroups = []map[string]interface{}{
		{"id": 1, "name": "Basic Greetings", "word_count": 25},
		{"id": 2, "name": "Common Phrases", "word_count": 30},
	}
)

func InitDB() error {
	log.Println("Removing any existing database...")
	os.Remove("words.db")
	os.Remove("words.db-shm")
	os.Remove("words.db-wal")
	
	log.Println("Opening new database connection...")
	db, err := sql.Open("sqlite", "words.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	DB = db

	// Drop all tables if they exist
	log.Println("Dropping any existing tables...")
	dropTables := []string{
		"DROP TABLE IF EXISTS word_groups",
		"DROP TABLE IF EXISTS word_review_items",
		"DROP TABLE IF EXISTS study_sessions",
		"DROP TABLE IF EXISTS study_activities",
		"DROP TABLE IF EXISTS words",
		"DROP TABLE IF EXISTS groups",
	}

	for _, query := range dropTables {
		log.Printf("Executing: %s", query)
		if _, err := DB.Exec(query); err != nil {
			return fmt.Errorf("failed to drop table: %v", err)
		}
	}

	// Create tables with minimal schema
	log.Println("Creating new tables...")
	createTables := []string{
		`CREATE TABLE words (
			id INTEGER PRIMARY KEY,
			arabic TEXT,
			transliteration TEXT,
			english TEXT
		)`,
		`CREATE TABLE groups (
			id INTEGER PRIMARY KEY,
			name TEXT
		)`,
		`CREATE TABLE word_groups (
			id INTEGER PRIMARY KEY,
			word_id INTEGER,
			group_id INTEGER
		)`,
	}

	for _, query := range createTables {
		log.Printf("Executing: %s", query)
		if _, err := DB.Exec(query); err != nil {
			return fmt.Errorf("failed to create table: %v", err)
		}
	}

	// Insert test data with explicit IDs
	log.Println("Inserting test data...")
	testData := []string{
		"INSERT INTO words (id, arabic, transliteration, english) VALUES (1, 'مرحبا', 'marhaba', 'hello')",
		"INSERT INTO groups (id, name) VALUES (1, 'Basic Greetings')",
		"INSERT INTO word_groups (id, word_id, group_id) VALUES (1, 1, 1)",
	}

	for _, query := range testData {
		log.Printf("Executing: %s", query)
		if _, err := DB.Exec(query); err != nil {
			return fmt.Errorf("failed to insert test data: %v", err)
		}
	}

	log.Println("Database initialization completed successfully")
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}