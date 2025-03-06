package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// WordPart represents a part of an Arabic word
type WordPart struct {
	Arabic       string   `json:"arabic"`
	Transliteration []string `json:"transliteration"`
}

// WordParts is a custom type for handling JSON array of WordPart
type WordParts []WordPart

// Value implements the driver.Valuer interface for WordParts
func (wp WordParts) Value() (driver.Value, error) {
	return json.Marshal(wp)
}

// Scan implements the sql.Scanner interface for WordParts
func (wp *WordParts) Scan(value interface{}) error {
	if value == nil {
		*wp = WordParts{}
		return nil
	}
	return json.Unmarshal(value.([]byte), wp)
}

// Word represents a vocabulary word in the system
type Word struct {
	ID             int64     `json:"id"`
	Arabic         string    `json:"arabic"`
	Transliteration string    `json:"transliteration"`
	English        string    `json:"english"`
	Parts          WordParts `json:"parts"`
	CreatedAt      time.Time `json:"created_at"`
	CorrectCount   int       `json:"correct_count,omitempty"`
	WrongCount     int       `json:"wrong_count,omitempty"`
	Groups         []Group   `json:"groups,omitempty"`
}

// WordWithStats represents a word with its study statistics
type WordWithStats struct {
	Word
	CorrectCount int `json:"correct_count"`
	WrongCount   int `json:"wrong_count"`
}

// WordReview represents a review of a word
type WordReview struct {
	WordID        int64     `json:"word_id"`
	StudySessionID int64     `json:"study_session_id"`
	Correct       bool      `json:"correct"`
	CreatedAt     time.Time `json:"created_at"`
} 