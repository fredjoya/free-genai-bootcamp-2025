package handlers

import (
	"lang-portal/internal/service"
	"net/http"

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
	// Get the most recent group's last study session
	groups, err := h.groupService.GetAllGroups(1, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(groups.Items) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no groups found"})
		return
	}

	session, err := h.groupService.GetLastStudySession(groups.Items[0].ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no study sessions found"})
		return
	}

	c.JSON(http.StatusOK, session)
}

// GetStudyProgress handles GET /api/dashboard/study_progress
func (h *DashboardHandler) GetStudyProgress(c *gin.Context) {
	progress, err := h.studyService.GetStudyProgress()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, progress)
}

// GetQuickStats handles GET /api/dashboard/quick-stats
func (h *DashboardHandler) GetQuickStats(c *gin.Context) {
	stats, err := h.studyService.GetDashboardStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
} 