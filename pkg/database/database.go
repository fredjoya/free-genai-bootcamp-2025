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
			group_id INTEGER
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

	// Insert word_groups separately
	stmt, err := DB.Prepare("INSERT INTO word_groups (word_id, group_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	wordGroups := [][2]int{
		{1, 1},
		{2, 2},
	}

	for _, wg := range wordGroups {
		if _, err := stmt.Exec(wg[0], wg[1]); err != nil {
			log.Printf("Warning: could not insert word_group %v: %v", wg, err)
			continue // Skip if there's an error
		}
	}

	// Insert word reviews
	stmt, err = DB.Prepare("INSERT INTO word_review_items (word_id, study_session_id, correct) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	reviews := [][3]interface{}{
		{1, 1, true},
		{2, 1, false},
	}

	for _, r := range reviews {
		if _, err := stmt.Exec(r[0], r[1], r[2]); err != nil {
			log.Printf("Warning: could not insert review %v: %v", r, err)
			continue
		}
	}

	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}