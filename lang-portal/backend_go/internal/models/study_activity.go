package models

import "time"

// LearningActivity represents a learning activity in the system
type LearningActivity struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	ThumbnailURL string   `json:"thumbnail_url"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// LearningActivitySession represents a study session for a specific activity
type LearningActivitySession struct {
	ID              int64     `json:"id"`
	ActivityName    string    `json:"activity_name"`
	GroupName       string    `json:"group_name"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	ReviewItemsCount int      `json:"review_items_count"`
} 