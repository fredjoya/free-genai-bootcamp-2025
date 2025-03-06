package repository

import (
	"database/sql"
	"lang-portal/internal/models"
	"lang-portal/pkg/database"
	"lang-portal/pkg/pagination"
)

// LearningActivityRepository handles database operations for learning activities
type LearningActivityRepository struct {
	db *sql.DB
}

// NewLearningActivityRepository creates a new learning activity repository
func NewLearningActivityRepository() *LearningActivityRepository {
	return &LearningActivityRepository{
		db: database.GetDB(),
	}
}

// GetLearningActivity retrieves a learning activity by ID
func (r *LearningActivityRepository) GetLearningActivity(id int64) (*models.LearningActivity, error) {
	var activity models.LearningActivity
	err := r.db.QueryRow(`
		SELECT id, name, thumbnail_url, description, created_at
		FROM study_activities
		WHERE id = ?
	`, id).Scan(&activity.ID, &activity.Name, &activity.ThumbnailURL, &activity.Description, &activity.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

// GetLearningActivitySessions retrieves all study sessions for a specific activity
func (r *LearningActivityRepository) GetLearningActivitySessions(activityID int64, page, itemsPerPage int) (pagination.PaginatedResponse[models.LearningActivitySession], error) {
	offset := (page - 1) * itemsPerPage

	// Get total count
	var total int
	err := r.db.QueryRow(`
		SELECT COUNT(*)
		FROM study_sessions ss
		JOIN study_activities sa ON ss.study_activity_id = sa.id
		WHERE sa.id = ?
	`, activityID).Scan(&total)
	if err != nil {
		return pagination.PaginatedResponse[models.LearningActivitySession]{}, err
	}

	// Get sessions
	rows, err := r.db.Query(`
		SELECT ss.id, sa.name, g.name, ss.created_at, ss.end_time,
			(SELECT COUNT(*) FROM word_review_items WHERE study_session_id = ss.id) as review_items_count
		FROM study_sessions ss
		JOIN study_activities sa ON ss.study_activity_id = sa.id
		JOIN groups g ON ss.group_id = g.id
		WHERE sa.id = ?
		ORDER BY ss.created_at DESC
		LIMIT ? OFFSET ?
	`, activityID, itemsPerPage, offset)
	if err != nil {
		return pagination.PaginatedResponse[models.LearningActivitySession]{}, err
	}
	defer rows.Close()

	var sessions []models.LearningActivitySession
	for rows.Next() {
		var session models.LearningActivitySession
		err := rows.Scan(&session.ID, &session.ActivityName, &session.GroupName, &session.StartTime, &session.EndTime, &session.ReviewItemsCount)
		if err != nil {
			return pagination.PaginatedResponse[models.LearningActivitySession]{}, err
		}
		sessions = append(sessions, session)
	}

	return pagination.NewPaginatedResponse(sessions, page, total, itemsPerPage), nil
}

// CreateLearningActivity creates a new learning activity
func (r *LearningActivityRepository) CreateLearningActivity(activity *models.LearningActivity) error {
	result, err := r.db.Exec(`
		INSERT INTO study_activities (name, thumbnail_url, description, created_at)
		VALUES (?, ?, ?, CURRENT_TIMESTAMP)
	`, activity.Name, activity.ThumbnailURL, activity.Description)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	activity.ID = id
	return nil
} 