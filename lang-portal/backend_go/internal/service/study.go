package service

import (
	"fmt"
	"lang-portal/internal/models"
	"lang-portal/internal/repository"
	"lang-portal/pkg/pagination"
)

// StudyService handles business logic for study operations
type StudyService struct {
	studyRepo *repository.StudyRepository
	wordRepo  *repository.WordRepository
}

// NewStudyService creates a new study service
func NewStudyService() *StudyService {
	return &StudyService{
		studyRepo: repository.NewStudyRepository(),
		wordRepo:  repository.NewWordRepository(),
	}
}

// CreateStudySession creates a new study session
func (s *StudyService) CreateStudySession(session *models.StudySession) error {
	if err := s.studyRepo.CreateStudySession(session); err != nil {
		return fmt.Errorf("error creating study session: %v", err)
	}
	return nil
}

// GetAllStudySessions retrieves all study sessions with pagination
func (s *StudyService) GetAllStudySessions(page, itemsPerPage int) (*pagination.PaginatedResponse[models.StudySession], error) {
	sessions, total, err := s.studyRepo.GetAllStudySessions(page, itemsPerPage)
	if err != nil {
		return nil, fmt.Errorf("error getting study sessions: %v", err)
	}

	return &pagination.PaginatedResponse[models.StudySession]{
		Items: sessions,
		Pagination: pagination.Pagination{
			CurrentPage:  page,
			TotalPages:   (total + itemsPerPage - 1) / itemsPerPage,
			TotalItems:   total,
			ItemsPerPage: itemsPerPage,
		},
	}, nil
}

// GetStudySession retrieves a study session by ID
func (s *StudyService) GetStudySession(id int64) (*models.StudySession, error) {
	session, err := s.studyRepo.GetStudySession(id)
	if err != nil {
		return nil, fmt.Errorf("error getting study session: %v", err)
	}
	return session, nil
}

// GetStudySessionWords retrieves all words reviewed in a study session with pagination
func (s *StudyService) GetStudySessionWords(sessionID int64, page, itemsPerPage int) (*pagination.PaginatedResponse[models.WordWithStats], error) {
	words, total, err := s.studyRepo.GetStudySessionWords(sessionID, page, itemsPerPage)
	if err != nil {
		return nil, fmt.Errorf("error getting study session words: %v", err)
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

// AddWordReview adds a review for a word in a study session
func (s *StudyService) AddWordReview(review *models.WordReview) error {
	// Verify the study session exists
	_, err := s.studyRepo.GetStudySession(review.StudySessionID)
	if err != nil {
		return fmt.Errorf("error verifying study session: %v", err)
	}

	// Add the review
	if err := s.wordRepo.AddReview(review); err != nil {
		return fmt.Errorf("error adding word review: %v", err)
	}

	return nil
}

// GetDashboardStats retrieves statistics for the dashboard
func (s *StudyService) GetDashboardStats() (*models.DashboardStats, error) {
	stats, err := s.studyRepo.GetDashboardStats()
	if err != nil {
		return nil, fmt.Errorf("error getting dashboard stats: %v", err)
	}
	return stats, nil
}

// GetStudyProgress retrieves study progress statistics
func (s *StudyService) GetStudyProgress() (*models.StudyProgress, error) {
	progress, err := s.studyRepo.GetStudyProgress()
	if err != nil {
		return nil, fmt.Errorf("error getting study progress: %v", err)
	}
	return progress, nil
}

// ResetHistory resets all study history
func (s *StudyService) ResetHistory() error {
	if err := s.studyRepo.ResetHistory(); err != nil {
		return fmt.Errorf("error resetting history: %v", err)
	}
	return nil
}

// FullReset resets the entire database and reseeds it
func (s *StudyService) FullReset() error {
	if err := s.studyRepo.FullReset(); err != nil {
		return fmt.Errorf("error resetting database: %v", err)
	}
	return nil
} 