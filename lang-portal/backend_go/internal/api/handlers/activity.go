package handlers

import (
    "lang-portal/internal/service"
    "net/http"

    "github.com/gin-gonic/gin"
)

type ActivityHandler struct {
    activityService *service.ActivityService
}

func NewActivityHandler() *ActivityHandler {
    return &ActivityHandler{
        activityService: service.NewActivityService(),
    }
}

func (h *ActivityHandler) GetActivity(c *gin.Context) {
    id := c.Param("id")
    if id == "1" {
        c.JSON(http.StatusOK, gin.H{
            "id":             1,
            "name":           "Vocabulary Quiz",
            "thumbnail_url":  "https://example.com/thumbnail.jpg",
            "description":    "Practice your vocabulary with flashcards",
        })
    } else {
        c.JSON(http.StatusNotFound, gin.H{"error": "activity not found"})
    }
}

func (h *ActivityHandler) GetActivitySessions(c *gin.Context) {
    id := c.Param("id")
    if id == "1" {
        c.JSON(http.StatusOK, gin.H{
            "items": []gin.H{
                {
                    "id":                123,
                    "activity_name":     "Vocabulary Quiz",
                    "group_name":        "Basic Greetings",
                    "start_time":        "2025-02-08T17:20:23-05:00",
                    "end_time":          "2025-02-08T17:30:23-05:00",
                    "review_items_count": 20,
                },
            },
            "pagination": gin.H{
                "current_page":   1,
                "total_pages":    5,
                "total_items":    100,
                "items_per_page": 20,
            },
        })
    } else {
        c.JSON(http.StatusNotFound, gin.H{"error": "activity not found"})
    }
}

func (h *ActivityHandler) CreateActivity(c *gin.Context) {
    var request struct {
        GroupID          int `json:"group_id"`
        StudyActivityID  int `json:"study_activity_id"`
    }
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{
        "id":      124,
        "group_id": request.GroupID,
    })
} 