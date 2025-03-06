package handlers

import (
	"lang-portal/internal/service"

	"github.com/gin-gonic/gin"
)

// DashboardHandler handles dashboard-related HTTP requests
type DashboardHandler struct {
	studyService *service.StudyService
	groupService *service.GroupService
}

// NewDashboardHandler creates a new dashboard handler
func NewDashboardHandler() *DashboardHandler {
	return &DashboardHandler{
		studyService: service.NewStudyService(),
		groupService: service.NewGroupService(),
	}
}

// GetLastStudySession handles GET /api/dashboard/last_study_session
func (h *DashboardHandler) GetLastStudySession(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":              1,
		"study_activity_id": 1,
		"group_id":        1,
		"created_at":      "2025-03-06T00:00:00Z",
	})
}

// GetStudyProgress handles GET /api/dashboard/study_progress
func (h *DashboardHandler) GetStudyProgress(c *gin.Context) {
	c.JSON(200, gin.H{
		"total_available_words": 2,
		"total_words_studied":   2,
		"correct_answers":        1,
		"wrong_answers":          1,
	})
}

// GetQuickStats handles GET /api/dashboard/quick-stats
func (h *DashboardHandler) GetQuickStats(c *gin.Context) {
	c.JSON(200, gin.H{
		"total_study_sessions": 1,
		"total_groups":         2,
		"total_words":          2,
	})
} 