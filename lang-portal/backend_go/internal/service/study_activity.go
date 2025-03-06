package service

import (
	"lang-portal/internal/models"
	"lang-portal/internal/repository"
	"lang-portal/pkg/pagination"
)

// LearningActivityService handles business logic for learning activities
type LearningActivityService struct {
	activityRepo *repository.LearningActivityRepository
}

// NewLearningActivityService creates a new learning activity service
func NewLearningActivityService() *LearningActivityService {
	return &LearningActivityService{
		activityRepo: repository.NewLearningActivityRepository(),
	}
}

// GetLearningActivity retrieves a learning activity by ID
func (s *LearningActivityService) GetLearningActivity(id int64) (*models.LearningActivity, error) {
	return s.activityRepo.GetLearningActivity(id)
}

// GetLearningActivitySessions retrieves all study sessions for a specific activity
func (s *LearningActivityService) GetLearningActivitySessions(activityID int64, page, itemsPerPage int) (pagination.PaginatedResponse[models.LearningActivitySession], error) {
	return s.activityRepo.GetLearningActivitySessions(activityID, page, itemsPerPage)
}

// CreateLearningActivity creates a new learning activity
func (s *LearningActivityService) CreateLearningActivity(activity *models.LearningActivity) error {
	return s.activityRepo.CreateLearningActivity(activity)
} 