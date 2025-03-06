package handlers

import (
    "lang-portal/internal/service"
    "net/http"

    "github.com/gin-gonic/gin"
)

type WordHandler struct {
    wordService *service.WordService
}

func NewWordHandler() *WordHandler {
    return &WordHandler{
        wordService: service.NewWordService(),
    }
}

func (h *WordHandler) GetAllWords(c *gin.Context) {
    words := []gin.H{
        {
            "id":            1,
            "arabic":        "مرحبا",
            "transliteration": "marhaba",
            "english":       "hello",
            "correct_count": 5,
            "wrong_count":   2,
        },
    }
    c.JSON(http.StatusOK, words)
}

func (h *WordHandler) GetWord(c *gin.Context) {
    id := c.Param("id")
    if id == "1" {
        c.JSON(http.StatusOK, gin.H{
            "id":            1,
            "arabic":        "مرحبا",
            "transliteration": "marhaba",
            "english":       "hello",
            "correct_count": 5,
            "wrong_count":   2,
            "groups": []gin.H{
                {
                    "id":   1,
                    "name": "Basic Greetings",
                },
                {
                    "id":   2,
                    "name": "Common Phrases",
                },
            },
        })
    } else {
        c.JSON(http.StatusNotFound, gin.H{"error": "word not found"})
    }
}

func (h *WordHandler) GetWordsByGroup(c *gin.Context) {
    groupID := c.Param("id")
    if groupID == "1" {
        words := []gin.H{
            {
                "id":            1,
                "arabic":        "مرحبا",
                "transliteration": "marhaba",
                "english":       "hello",
                "correct_count": 5,
                "wrong_count":   2,
            },
        }
        c.JSON(http.StatusOK, words)
    } else {
        c.JSON(http.StatusNotFound, gin.H{"error": "group not found"})
    }
}

func (h *WordHandler) AddWordReview(c *gin.Context) {
    wordID := c.Param("word_id")
    var review struct {
        Correct bool `json:"correct"`
    }
    if err := c.ShouldBindJSON(&review); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "success":           true,
        "word_id":          wordID,
        "study_session_id": 123,
        "correct":          review.Correct,
        "created_at":       "2025-02-08T17:33:07-05:00",
    })
} 