package models

import "time"

// Group represents a thematic group of words
type Group struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	WordCount int       `json:"word_count,omitempty"`
}

// StudySession represents a study session
type StudySession struct {
	ID                int64     `json:"id"`
	StudyActivityID   int64     `json:"study_activity_id"`
	GroupID           int64     `json:"group_id"`
	GroupName         string    `json:"group_name,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
	ReviewItemsCount  int       `json:"review_items_count,omitempty"`
}

// StudyActivity represents a study activity
type StudyActivity struct {
	ID              int64     `json:"id"`
	StudySessionID  int64     `json:"study_session_id"`
	GroupID         int64     `json:"group_id"`
	CreatedAt       time.Time `json:"created_at"`
}

// DashboardStats represents the quick stats for the dashboard
type DashboardStats struct {
	SuccessRate        float64 `json:"success_rate"`
	TotalStudySessions int     `json:"total_study_sessions"`
	TotalActiveGroups  int     `json:"total_active_groups"`
	StudyStreakDays    int     `json:"study_streak_days"`
}

// StudyProgress represents the study progress statistics
type StudyProgress struct {
	TotalWordsStudied    int `json:"total_words_studied"`
	TotalAvailableWords  int `json:"total_available_words"`
} 