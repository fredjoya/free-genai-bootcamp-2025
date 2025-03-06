package repository

import (
	"database/sql"
	"fmt"
	"lang-portal/internal/models"
	"lang-portal/pkg/database"
	"lang-portal/pkg/pagination"
)

// StudyRepository handles database operations for study sessions and activities
type StudyRepository struct {
	db *sql.DB
}

// NewStudyRepository creates a new study repository
func NewStudyRepository() *StudyRepository {
	return &StudyRepository{
		db: database.GetDB(),
	}
}

// CreateStudySession creates a new study session
func (r *StudyRepository) CreateStudySession(session *models.StudySession) error {
	query := `
		INSERT INTO study_sessions (study_activity_id, group_id, created_at)
		VALUES (?, ?, CURRENT_TIMESTAMP)
	`

	result, err := r.db.Exec(query, session.StudyActivityID, session.GroupID)
	if err != nil {
		return fmt.Errorf("error creating study session: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting last insert id: %v", err)
	}

	// Get the created session with group name
	err = r.db.QueryRow(`
		SELECT ss.id, ss.study_activity_id, ss.group_id, ss.created_at,
			g.name as group_name
		FROM study_sessions ss
		JOIN groups g ON ss.group_id = g.id
		WHERE ss.id = ?
	`, id).Scan(
		&session.ID,
		&session.StudyActivityID,
		&session.GroupID,
		&session.CreatedAt,
		&session.GroupName,
	)
	if err != nil {
		return fmt.Errorf("error getting created session: %v", err)
	}

	return nil
}

// GetAllStudySessions retrieves all study sessions with pagination
func (r *StudyRepository) GetAllStudySessions(page, itemsPerPage int) ([]models.StudySession, int, error) {
	offset := pagination.GetOffset(page, itemsPerPage)

	// Get total count
	var total int
	err := r.db.QueryRow("SELECT COUNT(*) FROM study_sessions").Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("error getting total sessions: %v", err)
	}

	// Get sessions with review counts
	query := `
		SELECT ss.id, ss.study_activity_id, ss.group_id, ss.created_at,
			g.name as group_name,
			COUNT(wri.id) as review_items_count
		FROM study_sessions ss
		JOIN groups g ON ss.group_id = g.id
		LEFT JOIN word_review_items wri ON ss.id = wri.study_session_id
		GROUP BY ss.id
		ORDER BY ss.created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.Query(query, itemsPerPage, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("error querying sessions: %v", err)
	}
	defer rows.Close()

	var sessions []models.StudySession
	for rows.Next() {
		var session models.StudySession
		err := rows.Scan(
			&session.ID,
			&session.StudyActivityID,
			&session.GroupID,
			&session.CreatedAt,
			&session.GroupName,
			&session.ReviewItemsCount,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("error scanning session: %v", err)
		}
		sessions = append(sessions, session)
	}

	return sessions, total, nil
}

// GetStudySession retrieves a study session by ID
func (r *StudyRepository) GetStudySession(id int64) (*models.StudySession, error) {
	query := `
		SELECT ss.id, ss.study_activity_id, ss.group_id, ss.created_at,
			g.name as group_name,
			COUNT(wri.id) as review_items_count
		FROM study_sessions ss
		JOIN groups g ON ss.group_id = g.id
		LEFT JOIN word_review_items wri ON ss.id = wri.study_session_id
		WHERE ss.id = ?
		GROUP BY ss.id
	`

	var session models.StudySession
	err := r.db.QueryRow(query, id).Scan(
		&session.ID,
		&session.StudyActivityID,
		&session.GroupID,
		&session.CreatedAt,
		&session.GroupName,
		&session.ReviewItemsCount,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("error querying session: %v", err)
	}

	return &session, nil
}

// GetStudySessionWords retrieves all words reviewed in a study session
func (r *StudyRepository) GetStudySessionWords(sessionID int64, page, itemsPerPage int) ([]models.WordWithStats, int, error) {
	offset := pagination.GetOffset(page, itemsPerPage)

	// Get total count
	var total int
	err := r.db.QueryRow(`
		SELECT COUNT(DISTINCT w.id) FROM words w
		JOIN word_review_items wri ON w.id = wri.word_id
		WHERE wri.study_session_id = ?
	`, sessionID).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("error getting total words: %v", err)
	}

	// Get words with review status
	query := `
		SELECT w.id, w.arabic, w.transliteration, w.english, w.parts, w.created_at,
			COUNT(CASE WHEN wri.correct = 1 THEN 1 END) as correct_count,
			COUNT(CASE WHEN wri.correct = 0 THEN 1 END) as wrong_count
		FROM words w
		JOIN word_review_items wri ON w.id = wri.word_id
		WHERE wri.study_session_id = ?
		GROUP BY w.id
		ORDER BY w.created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.Query(query, sessionID, itemsPerPage, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("error querying words: %v", err)
	}
	defer rows.Close()

	var words []models.WordWithStats
	for rows.Next() {
		var word models.WordWithStats
		err := rows.Scan(
			&word.ID,
			&word.Arabic,
			&word.Transliteration,
			&word.English,
			&word.Parts,
			&word.CreatedAt,
			&word.CorrectCount,
			&word.WrongCount,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("error scanning word: %v", err)
		}
		words = append(words, word)
	}

	return words, total, nil
}

// GetDashboardStats retrieves statistics for the dashboard
func (r *StudyRepository) GetDashboardStats() (*models.DashboardStats, error) {
	query := `
		SELECT 
			ROUND(AVG(CASE WHEN correct = 1 THEN 100 ELSE 0 END), 1) as success_rate,
			COUNT(DISTINCT study_session_id) as total_study_sessions,
			COUNT(DISTINCT group_id) as total_active_groups,
			COUNT(DISTINCT date(created_at)) as study_streak_days
		FROM word_review_items wri
		JOIN study_sessions ss ON wri.study_session_id = ss.id
	`

	var stats models.DashboardStats
	err := r.db.QueryRow(query).Scan(
		&stats.SuccessRate,
		&stats.TotalStudySessions,
		&stats.TotalActiveGroups,
		&stats.StudyStreakDays,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting dashboard stats: %v", err)
	}

	return &stats, nil
}

// GetStudyProgress retrieves study progress statistics
func (r *StudyRepository) GetStudyProgress() (*models.StudyProgress, error) {
	query := `
		SELECT 
			COUNT(DISTINCT wri.word_id) as total_words_studied,
			(SELECT COUNT(*) FROM words) as total_available_words
		FROM word_review_items wri
	`

	var progress models.StudyProgress
	err := r.db.QueryRow(query).Scan(
		&progress.TotalWordsStudied,
		&progress.TotalAvailableWords,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting study progress: %v", err)
	}

	return &progress, nil
}

// ResetHistory resets all study history
func (r *StudyRepository) ResetHistory() error {
	_, err := r.db.Exec("DELETE FROM word_review_items")
	if err != nil {
		return fmt.Errorf("error resetting history: %v", err)
	}
	return nil
}

// FullReset resets the entire database and reseeds it
func (r *StudyRepository) FullReset() error {
	// Delete all data from tables in the correct order
	queries := []string{
		"DELETE FROM word_review_items",
		"DELETE FROM study_sessions",
		"DELETE FROM study_activities",
		"DELETE FROM word_groups",
		"DELETE FROM words",
		"DELETE FROM groups",
	}

	for _, query := range queries {
		if _, err := r.db.Exec(query); err != nil {
			return fmt.Errorf("error executing query %s: %v", query, err)
		}
	}

	// TODO: Reseed the database with initial data
	// This would typically involve reading from seed files and inserting the data

	return nil
} 