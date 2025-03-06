package service

import (
	"fmt"
	"lang-portal/internal/models"
	"lang-portal/internal/repository"
	"lang-portal/pkg/pagination"
)

// WordService handles business logic for word operations
type WordService struct {
	wordRepo *repository.WordRepository
}

// NewWordService creates a new word service
func NewWordService() *WordService {
	return &WordService{
		wordRepo: repository.NewWordRepository(),
	}
}

// GetAllWords retrieves all words with pagination
func (s *WordService) GetAllWords(page, itemsPerPage int) (*pagination.PaginatedResponse[models.WordWithStats], error) {
	words, total, err := s.wordRepo.GetAll(page, itemsPerPage)
	if err != nil {
		return nil, fmt.Errorf("error getting words: %v", err)
	}

	return &pagination.PaginatedResponse[models.WordWithStats]{
		Items: words,
		Pagination: pagination.Pagination{
			CurrentPage:  page,
			TotalPages:   (total + itemsPerPage - 1) / itemsPerPage,
			TotalItems:   total,
			ItemsPerPage: itemsPerPage,
		},
	}, nil
}

// GetWord retrieves a word by ID
func (s *WordService) GetWord(id int64) (*models.Word, error) {
	word, err := s.wordRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error getting word: %v", err)
	}
	return word, nil
}

// GetWordsByGroup retrieves all words in a group with pagination
func (s *WordService) GetWordsByGroup(groupID int64, page, itemsPerPage int) (*pagination.PaginatedResponse[models.WordWithStats], error) {
	words, total, err := s.wordRepo.GetByGroup(groupID, page, itemsPerPage)
	if err != nil {
		return nil, fmt.Errorf("error getting words by group: %v", err)
	}

	return &pagination.PaginatedResponse[models.WordWithStats]{
		Items: words,
		Pagination: pagination.Pagination{
			CurrentPage:  page,
			TotalPages:   (total + itemsPerPage - 1) / itemsPerPage,
			TotalItems:   total,
			ItemsPerPage: itemsPerPage,
		},
	}, nil
}

// AddWordReview adds a review for a word
func (s *WordService) AddWordReview(review *models.WordReview) error {
	if err := s.wordRepo.AddReview(review); err != nil {
		return fmt.Errorf("error adding word review: %v", err)
	}
	return nil
} 