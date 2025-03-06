package handlers

import (
	"lang-portal/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GroupHandler handles group-related HTTP requests
type GroupHandler struct {
	groupService *service.GroupService
}

// NewGroupHandler creates a new group handler
func NewGroupHandler() *GroupHandler {
	return &GroupHandler{
		groupService: service.NewGroupService(),
	}
}

// GetAllGroups handles GET /api/groups
func (h *GroupHandler) GetAllGroups(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	itemsPerPage, _ := strconv.Atoi(c.DefaultQuery("items_per_page", "100"))

	response, err := h.groupService.GetAllGroups(page, itemsPerPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetGroup handles GET /api/groups/:id
func (h *GroupHandler) GetGroup(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID"})
		return
	}

	group, err := h.groupService.GetGroup(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "group not found"})
		return
	}

	c.JSON(http.StatusOK, group)
}

// GetGroupStudySessions handles GET /api/groups/:id/study_sessions
func (h *GroupHandler) GetGroupStudySessions(c *gin.Context) {
	groupID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	itemsPerPage, _ := strconv.Atoi(c.DefaultQuery("items_per_page", "100"))

	response, err := h.groupService.GetGroupStudySessions(groupID, page, itemsPerPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
} 