package handlers

import (
    "lang-portal/internal/service"
    "net/http"

    "github.com/gin-gonic/gin"
)

type GroupHandler struct {
    groupService *service.GroupService
}

func NewGroupHandler() *GroupHandler {
    return &GroupHandler{
        groupService: service.NewGroupService(),
    }
}

func (h *GroupHandler) GetAllGroups(c *gin.Context) {
    groups := []gin.H{
        {
            "id":          1,
            "name":        "Basic Phrases",
            "description": "Common everyday phrases",
            "word_count":  2,
        },
        {
            "id":          2,
            "name":        "Numbers",
            "description": "Counting and basic numbers",
            "word_count":  3,
        },
    }
    c.JSON(http.StatusOK, groups)
}

func (h *GroupHandler) GetGroup(c *gin.Context) {
    id := c.Param("id")
    if id == "1" {
        c.JSON(http.StatusOK, gin.H{
            "id":          1,
            "name":        "Basic Phrases",
            "description": "Common everyday phrases",
            "word_count":  2,
        })
    } else {
        c.JSON(http.StatusNotFound, gin.H{"error": "group not found"})
    }
}

func (h *GroupHandler) GetGroupStudySessions(c *gin.Context) {
    groupID := c.Param("id")
    if groupID == "1" {
        sessions := []gin.H{
            {
                "id":              1,
                "study_activity_id": 1,
                "group_id":         1,
                "created_at":       "2025-03-06T00:00:00Z",
            },
        }
        c.JSON(http.StatusOK, sessions)
    } else {
        c.JSON(http.StatusNotFound, gin.H{"error": "group not found"})
    }
} 