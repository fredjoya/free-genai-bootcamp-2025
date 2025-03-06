package service

import (
	"fmt"
	"lang-portal/internal/models"
	"lang-portal/internal/repository"
	"lang-portal/pkg/pagination"
)

// GroupService handles business logic for group operations
type GroupService struct {
	groupRepo *repository.GroupRepository
}

// NewGroupService creates a new group service
func NewGroupService() *GroupService {
	return &GroupService{
		groupRepo: repository.NewGroupRepository(),
	}
}

// GetAllGroups retrieves all groups with pagination
func (s *GroupService) GetAllGroups(page, itemsPerPage int) (*pagination.PaginatedResponse[models.Group], error) {
	groups, total, err := s.groupRepo.GetAll(page, itemsPerPage)
	if err != nil {
		return nil, fmt.Errorf("error getting groups: %v", err)
	}

	return &pagination.PaginatedResponse[models.Group]{
		Items: groups,
		Pagination: pagination.Pagination{
			CurrentPage:  page,
			TotalPages:   (total + itemsPerPage - 1) / itemsPerPage,
			TotalItems:   total,
			ItemsPerPage: itemsPerPage,
		},
	}, nil
}

// GetGroup retrieves a group by ID
func (s *GroupService) GetGroup(id int64) (*models.Group, error) {
	group, err := s.groupRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error getting group: %v", err)
	}
	return group, nil
}

// GetGroupStudySessions retrieves study sessions for a group with pagination
func (s *GroupService) GetGroupStudySessions(groupID int64, page, itemsPerPage int) (*pagination.PaginatedResponse[models.StudySession], error) {
	sessions, total, err := s.groupRepo.GetStudySessions(groupID, page, itemsPerPage)
	if err != nil {
		return nil, fmt.Errorf("error getting group study sessions: %v", err)
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

// GetLastStudySession retrieves the most recent study session for a group
func (s *GroupService) GetLastStudySession(groupID int64) (*models.StudySession, error) {
	session, err := s.groupRepo.GetLastStudySession(groupID)
	if err != nil {
		return nil, fmt.Errorf("error getting last study session: %v", err)
	}
	return session, nil
} 