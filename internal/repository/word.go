package repository

import (
	"database/sql"
	"fmt"
	"log"

	"lang-portal/internal/models"
)

type WordRepository struct {
	db *sql.DB
}

func NewWordRepository(db *sql.DB) *WordRepository {
	return &WordRepository{db: db}
}

func (r *WordRepository) GetAllWords(page, pageSize int) (*models.PaginatedResponse, error) {
	log.Println("Starting GetAllWords function")

	// Debug: Print all tables in the database
	rows, err := r.db.Query("SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		return nil, fmt.Errorf("error querying tables: %v", err)
	}
	defer rows.Close()

	log.Println("Existing tables:")
	for rows.Next() {
		var name string
		rows.Scan(&name)
		log.Printf("- %s", name)
	}

	// Count total words
	var total int
	countQuery := "SELECT COUNT(*) FROM words"
	log.Printf("Executing count query: %s", countQuery)
	if err := r.db.QueryRow(countQuery).Scan(&total); err != nil {
		return nil, fmt.Errorf("error counting words: %v", err)
	}
	log.Printf("Total words found: %d", total)

	// Get words with pagination
	query := `
		SELECT 
			id,
			word,
			translation,
			example,
			created_at
		FROM words
		ORDER BY id
		LIMIT ? OFFSET ?
	`
	log.Printf("Executing main query: %s", query)
	offset := (page - 1) * pageSize
	log.Printf("Page: %d, PageSize: %d, Offset: %d", page, pageSize, offset)

	rows, err = r.db.Query(query, pageSize, offset)
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
			&w.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning word: %v", err)
		}
		words = append(words, w)
	}

	log.Printf("Found %d words", len(words))

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

func (r *WordRepository) GetWord(id int) (*models.Word, error) {
	query := `
		SELECT 
			w.id,
			w.word,
			w.translation,
			w.example,
			COALESCE(wr.review_type, '') as review_type,
			w.created_at
		FROM words w
		LEFT JOIN (
			SELECT word_id, review_type, MAX(created_at) as latest_review
			FROM word_reviews
			GROUP BY word_id
		) wr ON w.id = wr.word_id
		WHERE w.id = ?
	`
	
	var word models.Word
	err := r.db.QueryRow(query, id).Scan(
		&word.ID,
		&word.Word,
		&word.Translation,
		&word.Example,
		&word.ReviewType,
		&word.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying word: %v", err)
	}

	return &word, nil
}

func (r *WordRepository) GetWordsByGroup(groupID, page, pageSize int) (*models.PaginatedResponse, error) {
	// Get total count
	var total int
	countQuery := `
		SELECT COUNT(*) 
		FROM words w
		JOIN word_groups wg ON w.id = wg.word_id
		WHERE wg.group_id = ?
	`
	if err := r.db.QueryRow(countQuery, groupID).Scan(&total); err != nil {
		return nil, fmt.Errorf("error counting words: %v", err)
	}

	// Get paginated words
	query := `
		SELECT 
			w.id,
			w.word,
			w.translation,
			w.example,
			COALESCE(wr.review_type, '') as review_type,
			w.created_at
		FROM words w
		JOIN word_groups wg ON w.id = wg.word_id
		LEFT JOIN (
			SELECT word_id, review_type, MAX(created_at) as latest_review
			FROM word_reviews
			GROUP BY word_id
		) wr ON w.id = wr.word_id
		WHERE wg.group_id = ?
		ORDER BY w.id
		LIMIT ? OFFSET ?
	`
	offset := (page - 1) * pageSize
	rows, err := r.db.Query(query, groupID, pageSize, offset)
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