package handlers

import (
	"lang-portal/internal/models"
	"lang-portal/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// WordHandler handles word-related HTTP requests
type WordHandler struct {
	wordService *service.WordService
}

// NewWordHandler creates a new word handler
func NewWordHandler() *WordHandler {
	return &WordHandler{
		wordService: service.NewWordService(),
	}
}

// GetAllWords handles GET /api/words
func (h *WordHandler) GetAllWords(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	itemsPerPage, _ := strconv.Atoi(c.DefaultQuery("items_per_page", "100"))

	response, err := h.wordService.GetAllWords(page, itemsPerPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetWord handles GET /api/words/:id
func (h *WordHandler) GetWord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid word ID"})
		return
	}

	word, err := h.wordService.GetWord(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "word not found"})
		return
	}

	c.JSON(http.StatusOK, word)
}

// GetWordsByGroup handles GET /api/groups/:id/words
func (h *WordHandler) GetWordsByGroup(c *gin.Context) {
	groupID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	itemsPerPage, _ := strconv.Atoi(c.DefaultQuery("items_per_page", "100"))

	response, err := h.wordService.GetWordsByGroup(groupID, page, itemsPerPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// AddWordReview handles POST /api/study_sessions/:word_id/review
func (h *WordHandler) AddWordReview(c *gin.Context) {
	wordID, err := strconv.ParseInt(c.Param("word_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid word ID"})
		return
	}

	var review struct {
		StudySessionID int64 `json:"study_session_id" binding:"required"`
		Correct       bool  `json:"correct" binding:"required"`
	}

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	wordReview := &models.WordReview{
		WordID:        wordID,
		StudySessionID: review.StudySessionID,
		Correct:       review.Correct,
	}

	if err := h.wordService.AddWordReview(wordReview); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wordReview)
} 