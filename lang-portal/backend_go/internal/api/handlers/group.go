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
            "name":        "Basic Greetings",
            "word_count":  25,
        },
    }
    c.JSON(http.StatusOK, groups)
}

func (h *GroupHandler) GetGroup(c *gin.Context) {
    id := c.Param("id")
    if id == "1" {
        c.JSON(http.StatusOK, gin.H{
            "id":          1,
            "name":        "Basic Greetings",
            "word_count":  25,
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
                "id":               123,
                "activity_name":     "Vocabulary Quiz",
                "group_name":        "Basic Greetings",
                "start_time":        "2025-02-08T17:20:23-05:00",
                "end_time":          "2025-02-08T17:30:23-05:00",
                "review_items_count": 20,
            },
        }
        c.JSON(http.StatusOK, sessions)
    } else {
        c.JSON(http.StatusNotFound, gin.H{"error": "group not found"})
    }
} 