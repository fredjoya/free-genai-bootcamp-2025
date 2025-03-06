package handlers

import (
	"lang-portal/internal/models"
	"lang-portal/internal/service"
	"net/http"
	"strconv"

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
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	itemsPerPage, _ := strconv.Atoi(c.DefaultQuery("items_per_page", "100"))

	response, err := h.studyService.GetAllStudySessions(page, itemsPerPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetStudySession handles GET /api/study_sessions/:id
func (h *StudyHandler) GetStudySession(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session ID"})
		return
	}

	session, err := h.studyService.GetStudySession(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "study session not found"})
		return
	}

	c.JSON(http.StatusOK, session)
}

// GetStudySessionWords handles GET /api/study_sessions/:id/words
func (h *StudyHandler) GetStudySessionWords(c *gin.Context) {
	sessionID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	itemsPerPage, _ := strconv.Atoi(c.DefaultQuery("items_per_page", "100"))

	response, err := h.studyService.GetStudySessionWords(sessionID, page, itemsPerPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// ResetHistory handles POST /api/reset_history
func (h *StudyHandler) ResetHistory(c *gin.Context) {
	if err := h.studyService.ResetHistory(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Study history has been reset successfully"})
}

// FullReset handles POST /api/full_reset
func (h *StudyHandler) FullReset(c *gin.Context) {
	if err := h.studyService.FullReset(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Database has been reset and reseeded successfully"})
} 