package repository

import (
	"database/sql"
	"fmt"
	"lang-portal/internal/models"
	"lang-portal/pkg/database"
	"lang-portal/pkg/pagination"
)

// WordRepository handles database operations for words
type WordRepository struct {
	db *sql.DB
}

// NewWordRepository creates a new word repository
func NewWordRepository() *WordRepository {
	return &WordRepository{
		db: database.GetDB(),
	}
}

// GetAll retrieves all words with pagination
func (r *WordRepository) GetAll(page, itemsPerPage int) ([]models.WordWithStats, int, error) {
	offset := pagination.GetOffset(page, itemsPerPage)

	// Get total count
	var total int
	err := r.db.QueryRow("SELECT COUNT(*) FROM words").Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("error getting total words: %v", err)
	}

	// Get words with stats
	query := `
		SELECT w.id, w.arabic, w.transliteration, w.english, w.parts, w.created_at,
			COUNT(CASE WHEN wri.correct = 1 THEN 1 END) as correct_count,
			COUNT(CASE WHEN wri.correct = 0 THEN 1 END) as wrong_count
		FROM words w
		LEFT JOIN word_review_items wri ON w.id = wri.word_id
		GROUP BY w.id
		ORDER BY w.created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.Query(query, itemsPerPage, offset)
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

// GetByID retrieves a word by ID with its groups
func (r *WordRepository) GetByID(id int64) (*models.Word, error) {
	query := `
		SELECT w.id, w.arabic, w.transliteration, w.english, w.parts, w.created_at,
			g.id, g.name, g.created_at
		FROM words w
		LEFT JOIN word_groups wg ON w.id = wg.word_id
		LEFT JOIN groups g ON wg.group_id = g.id
		WHERE w.id = ?
	`

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("error querying word: %v", err)
	}
	defer rows.Close()

	var word models.Word
	var groups []models.Group
	wordFound := false

	for rows.Next() {
		wordFound = true
		var group models.Group
		err := rows.Scan(
			&word.ID,
			&word.Arabic,
			&word.Transliteration,
			&word.English,
			&word.Parts,
			&word.CreatedAt,
			&group.ID,
			&group.Name,
			&group.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning word: %v", err)
		}
		if group.ID != 0 {
			groups = append(groups, group)
		}
	}

	if !wordFound {
		return nil, sql.ErrNoRows
	}

	word.Groups = groups
	return &word, nil
}

// GetByGroup retrieves all words in a group with pagination
func (r *WordRepository) GetByGroup(groupID int64, page, itemsPerPage int) ([]models.WordWithStats, int, error) {
	offset := pagination.GetOffset(page, itemsPerPage)

	// Get total count
	var total int
	err := r.db.QueryRow(`
		SELECT COUNT(*) FROM words w
		JOIN word_groups wg ON w.id = wg.word_id
		WHERE wg.group_id = ?
	`, groupID).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("error getting total words: %v", err)
	}

	// Get words with stats
	query := `
		SELECT w.id, w.arabic, w.transliteration, w.english, w.parts, w.created_at,
			COUNT(CASE WHEN wri.correct = 1 THEN 1 END) as correct_count,
			COUNT(CASE WHEN wri.correct = 0 THEN 1 END) as wrong_count
		FROM words w
		JOIN word_groups wg ON w.id = wg.word_id
		LEFT JOIN word_review_items wri ON w.id = wri.word_id
		WHERE wg.group_id = ?
		GROUP BY w.id
		ORDER BY w.created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.Query(query, groupID, itemsPerPage, offset)
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

// AddReview adds a word review
func (r *WordRepository) AddReview(review *models.WordReview) error {
	query := `
		INSERT INTO word_review_items (word_id, study_session_id, correct, created_at)
		VALUES (?, ?, ?, CURRENT_TIMESTAMP)
	`

	result, err := r.db.Exec(query, review.WordID, review.StudySessionID, review.Correct)
	if err != nil {
		return fmt.Errorf("error inserting review: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting last insert id: %v", err)
	}

	// Get the created review
	err = r.db.QueryRow(`
		SELECT created_at FROM word_review_items WHERE id = ?
	`, id).Scan(&review.CreatedAt)
	if err != nil {
		return fmt.Errorf("error getting review timestamp: %v", err)
	}

	return nil
} 