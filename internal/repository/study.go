package repository

import (
	"database/sql"
	"fmt"

	"github.com/your-project/models"
)

type StudyRepository struct {
	db *sql.DB
}

func (r *StudyRepository) GetStudyProgress() (*models.StudyProgress, error) {
	query := `
		SELECT
			COUNT(*) as total_reviews,
			SUM(CASE WHEN review_type = 'correct' THEN 1 ELSE 0 END) as correct_reviews,
			SUM(CASE WHEN review_type = 'incorrect' THEN 1 ELSE 0 END) as incorrect_reviews
		FROM word_reviews
		WHERE created_at >= datetime('now', '-30 days')
	`
	
	progress := &models.StudyProgress{}
	err := r.db.QueryRow(query).Scan(
		&progress.TotalReviews,
		&progress.CorrectReviews,
		&progress.IncorrectReviews,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting study progress: %v", err)
	}

	if progress.TotalReviews > 0 {
		progress.CompletionRate = float64(progress.CorrectReviews) / float64(progress.TotalReviews) * 100
	}

	return progress, nil
}

func (r *StudyRepository) CreateStudySession(groupID int) (*models.StudySession, error) {
	query := `
		INSERT INTO study_sessions (group_id, started_at)
		VALUES (?, datetime('now'))
		RETURNING id, group_id, started_at, completed_at
	`
	
	session := &models.StudySession{}
	err := r.db.QueryRow(query, groupID).Scan(
		&session.ID,
		&session.GroupID,
		&session.StartedAt,
		&session.CompletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating study session: %v", err)
	}

	return session, nil
}

func (r *StudyRepository) GetStudySessionWords(sessionID, page, pageSize int) (*models.PaginatedResponse, error) {
	query := `
		SELECT 
			w.id,
			w.word,
			w.translation,
			w.example,
			wr.review_type,
			wr.created_at
		FROM words w
		JOIN word_reviews wr ON w.id = wr.word_id
		WHERE wr.study_session_id = ?
		ORDER BY wr.created_at DESC
		LIMIT ? OFFSET ?
	`
	
	// Get total count
	var total int
	countQuery := `
		SELECT COUNT(DISTINCT w.id)
		FROM words w
		JOIN word_reviews wr ON w.id = wr.word_id
		WHERE wr.study_session_id = ?
	`
	if err := r.db.QueryRow(countQuery, sessionID).Scan(&total); err != nil {
		return nil, fmt.Errorf("error getting total words: %v", err)
	}

	offset := (page - 1) * pageSize
	rows, err := r.db.Query(query, sessionID, pageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("error querying words: %v", err)
	}
	defer rows.Close()

	var words []models.Word
	for rows.Next() {
		var w models.Word
		if err := rows.Scan(
			&w.ID,
			&w.Word,
			&w.Translation,
			&w.Example,
			&w.ReviewType,
			&w.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning word: %v", err)
		}
		words = append(words, w)
	}

	return &models.PaginatedResponse{
		Items: words,
		Pagination: models.Pagination{
			CurrentPage:   page,
			TotalPages:   (total + pageSize - 1) / pageSize,
			TotalItems:   total,
			ItemsPerPage: pageSize,
		},
	}, nil
} 