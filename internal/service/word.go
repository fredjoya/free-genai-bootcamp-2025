package service

import (
	"log"
	"lang-portal/internal/models"
	"lang-portal/internal/repository"
)

type WordService struct {
	wordRepo *repository.WordRepository
}

func NewWordService(wordRepo *repository.WordRepository) *WordService {
	return &WordService{wordRepo: wordRepo}
}

func (s *WordService) GetAllWords(page, pageSize int) (*models.PaginatedResponse, error) {
	log.Printf("WordService.GetAllWords called with page=%d, pageSize=%d", page, pageSize)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	return s.wordRepo.GetAllWords(page, pageSize)
}

func (s *WordService) GetWord(id int) (*models.Word, error) {
	return s.wordRepo.GetWord(id)
}

func (s *WordService) GetWordsByGroup(groupID, page, pageSize int) (*models.PaginatedResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	return s.wordRepo.GetWordsByGroup(groupID, page, pageSize)
} 