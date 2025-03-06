package database

import (
	"database/sql"
	"fmt"
	"os"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

// InitDB initializes the database connection and sets up the schema
func InitDB() error {
	// Clean start
	os.Remove("words.db")
	
	db, err := sql.Open("sqlite", "words.db")
	if err != nil {
		return err
	}
	DB = db

	// First drop all tables
	dropTables := []string{
		"DROP TABLE IF EXISTS word_review_items;",
		"DROP TABLE IF EXISTS word_groups;",
		"DROP TABLE IF EXISTS study_sessions;",
		"DROP TABLE IF EXISTS study_activities;",
		"DROP TABLE IF EXISTS words;",
		"DROP TABLE IF EXISTS groups;",
	}

	for _, query := range dropTables {
		if _, err := DB.Exec(query); err != nil {
			return err
		}
	}

	// Create tables one by one
	createTables := []string{
		`CREATE TABLE words (
			id INTEGER PRIMARY KEY,
			arabic TEXT NOT NULL,
			transliteration TEXT NOT NULL,
			english TEXT NOT NULL,
			parts JSON
		);`,

		`CREATE TABLE groups (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL
		);`,

		`CREATE TABLE word_groups (
			word_id INTEGER,
			group_id INTEGER,
			PRIMARY KEY (word_id, group_id)
		);`,

		`CREATE TABLE study_activities (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,

		`CREATE TABLE study_sessions (
			id INTEGER PRIMARY KEY,
			study_activity_id INTEGER,
			group_id INTEGER,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,

		`CREATE TABLE word_review_items (
			word_id INTEGER,
			study_session_id INTEGER,
			correct BOOLEAN,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
	}

	for _, query := range createTables {
		if _, err := DB.Exec(query); err != nil {
			return err
		}
	}

	// Insert test data with separate statements
	testData := []string{
		"INSERT INTO words (id, arabic, transliteration, english) VALUES (1, 'مرحبا', 'marhaba', 'hello');",
		"INSERT INTO words (id, arabic, transliteration, english) VALUES (2, 'شكرا', 'shukran', 'thank you');",
		
		"INSERT INTO groups (id, name) VALUES (1, 'Basic Greetings');",
		"INSERT INTO groups (id, name) VALUES (2, 'Common Phrases');",
		
		"INSERT INTO study_activities (id, name, description) VALUES (1, 'Vocabulary Quiz', 'Practice vocabulary');",
		
		"INSERT INTO study_sessions (id, study_activity_id, group_id) VALUES (1, 1, 1);",
	}

	for _, query := range testData {
		if _, err := DB.Exec(query); err != nil {
			return err
		}
	}

	// Insert word_groups separately with explicit transaction
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	wordGroups := [][2]int{
		{1, 1},
		{2, 2},
	}

	for _, wg := range wordGroups {
		if _, err := tx.Exec("INSERT INTO word_groups (word_id, group_id) VALUES (?, ?)", wg[0], wg[1]); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	// Insert word reviews
	reviews := [][3]interface{}{
		{1, 1, true},
		{2, 1, false},
	}

	for _, r := range reviews {
		if _, err := DB.Exec("INSERT INTO word_review_items (word_id, study_session_id, correct) VALUES (?, ?, ?)", r[0], r[1], r[2]); err != nil {
			return err
		}
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