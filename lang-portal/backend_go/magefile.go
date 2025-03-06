//go:build mage
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// DBInit initializes the SQLite database
func DBInit() error {
	db, err := sql.Open("sqlite3", "words.db")
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	// Read and execute migration files
	migrationsDir := "db/migrations"
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("error reading migrations directory: %v", err)
	}

	// Sort files to ensure correct order
	var migrationFiles []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			migrationFiles = append(migrationFiles, file.Name())
		}
	}
	sort.Strings(migrationFiles)

	// Execute each migration file
	for _, file := range migrationFiles {
		content, err := ioutil.ReadFile(filepath.Join(migrationsDir, file))
		if err != nil {
			return fmt.Errorf("error reading migration file %s: %v", file, err)
		}

		_, err = db.Exec(string(content))
		if err != nil {
			return fmt.Errorf("error executing migration %s: %v", file, err)
		}

		log.Printf("Executed migration: %s\n", file)
	}

	return nil
}

// DBSeed seeds the database with initial data
func DBSeed() error {
	db, err := sql.Open("sqlite3", "words.db")
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	// Read seed files
	seedsDir := "db/seeds"
	files, err := ioutil.ReadDir(seedsDir)
	if err != nil {
		return fmt.Errorf("error reading seeds directory: %v", err)
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		content, err := ioutil.ReadFile(filepath.Join(seedsDir, file.Name()))
		if err != nil {
			return fmt.Errorf("error reading seed file %s: %v", file.Name(), err)
		}

		var words []struct {
			Arabic       string `json:"arabic"`
			Transliteration string `json:"transliteration"`
			English     string `json:"english"`
			Parts       json.RawMessage `json:"parts"`
		}

		if err := json.Unmarshal(content, &words); err != nil {
			return fmt.Errorf("error parsing seed file %s: %v", file.Name(), err)
		}

		// Extract group name from filename (without .json extension)
		groupName := strings.TrimSuffix(file.Name(), ".json")

		// Insert group
		result, err := db.Exec("INSERT INTO groups (name) VALUES (?)", groupName)
		if err != nil {
			return fmt.Errorf("error inserting group %s: %v", groupName, err)
		}

		groupID, err := result.LastInsertId()
		if err != nil {
			return fmt.Errorf("error getting group ID: %v", err)
		}

		// Insert words and word-group relationships
		for _, word := range words {
			result, err := db.Exec(
				"INSERT INTO words (arabic, transliteration, english, parts) VALUES (?, ?, ?, ?)",
				word.Arabic, word.Transliteration, word.English, word.Parts,
			)
			if err != nil {
				return fmt.Errorf("error inserting word %s: %v", word.Arabic, err)
			}

			wordID, err := result.LastInsertId()
			if err != nil {
				return fmt.Errorf("error getting word ID: %v", err)
			}

			_, err = db.Exec(
				"INSERT INTO word_groups (word_id, group_id) VALUES (?, ?)",
				wordID, groupID,
			)
			if err != nil {
				return fmt.Errorf("error inserting word-group relationship: %v", err)
			}
		}

		log.Printf("Seeded data from %s\n", file.Name())
	}

	return nil
}

// DBReset resets the database by removing it and reinitializing
func DBReset() error {
	// Remove existing database
	if err := os.Remove("words.db"); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error removing database: %v", err)
	}

	// Initialize database
	if err := DBInit(); err != nil {
		return err
	}

	// Seed database
	if err := DBSeed(); err != nil {
		return err
	}

	log.Println("Database reset complete")
	return nil
} 