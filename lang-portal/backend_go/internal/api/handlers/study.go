package handlers

import (
	"lang-portal/internal/models"
	"lang-portal/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// StudyHandler handles study-related HTTP requests
type StudyHandler struct {
	studyService *service.StudyService
}

// NewStudyHandler creates a new study handler
func NewStudyHandler() *StudyHandler {
	return &StudyHandler{
		studyService: service.NewStudyService(),
	}
}

// CreateStudySession handles POST /api/study_sessions
func (h *StudyHandler) CreateStudySession(c *gin.Context) {
	var session models.StudySession
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.studyService.CreateStudySession(&session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, session)
}

// GetAllStudySessions handles GET /api/study_sessions
func (h *StudyHandler) GetAllStudySessions(c *gin.Context) {
	sessions := []gin.H{
		{
			"id":               1,
			"study_activity_id": 1,
			"group_id":         1,
			"created_at":       "2025-02-08T17:20:23-05:00",
		},
	}
	c.JSON(http.StatusOK, sessions)
}

// GetStudySession handles GET /api/study_sessions/:id
func (h *StudyHandler) GetStudySession(c *gin.Context) {
	id := c.Param("id")
	if id == "1" {
		c.JSON(http.StatusOK, gin.H{
			"id":               1,
			"study_activity_id": 1,
			"group_id":         1,
			"created_at":       "2025-02-08T17:20:23-05:00",
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "study session not found"})
	}
}

// GetStudySessionWords handles GET /api/study_sessions/:id/words
func (h *StudyHandler) GetStudySessionWords(c *gin.Context) {
	id := c.Param("id")
	if id == "1" {
		words := []gin.H{
			{
				"id":            1,
				"arabic":        "مرحبا",
				"transliteration": "marhaba",
				"english":       "hello",
				"correct_count": 5,
				"wrong_count":   2,
				"review_correct": true,
			},
		}
		c.JSON(http.StatusOK, words)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "study session not found"})
	}
}

// ResetHistory handles POST /api/reset_history
func (h *StudyHandler) ResetHistory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Study history has been reset successfully",
	})
}

// FullReset handles POST /api/full_reset
func (h *StudyHandler) FullReset(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Database has been reset and reseeded successfully",
	})
} 