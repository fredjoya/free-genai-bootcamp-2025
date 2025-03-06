package handlers

import (
	"lang-portal/internal/models"
	"lang-portal/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// LearningActivityHandler handles learning activity-related HTTP requests
type LearningActivityHandler struct {
	activityService *service.LearningActivityService
}

// NewLearningActivityHandler creates a new learning activity handler
func NewLearningActivityHandler() *LearningActivityHandler {
	return &LearningActivityHandler{
		activityService: service.NewLearningActivityService(),
	}
}

// GetLearningActivity handles GET /api/study_activities/:id
func (h *LearningActivityHandler) GetLearningActivity(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid activity ID"})
		return
	}

	activity, err := h.activityService.GetLearningActivity(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "learning activity not found"})
		return
	}

	c.JSON(http.StatusOK, activity)
}

// GetLearningActivitySessions handles GET /api/study_activities/:id/study_sessions
func (h *LearningActivityHandler) GetLearningActivitySessions(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid activity ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	itemsPerPage, _ := strconv.Atoi(c.DefaultQuery("items_per_page", "100"))

	response, err := h.activityService.GetLearningActivitySessions(id, page, itemsPerPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// CreateLearningActivity handles POST /api/study_activities
func (h *LearningActivityHandler) CreateLearningActivity(c *gin.Context) {
	var activity models.LearningActivity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.activityService.CreateLearningActivity(&activity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, activity)
} 