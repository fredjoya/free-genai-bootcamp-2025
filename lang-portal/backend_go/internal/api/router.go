package api

import (
	"lang-portal/internal/api/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and returns the Gin router with all routes configured
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Health check
	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Dashboard endpoints
	dashboardHandler := handlers.NewDashboardHandler()
	router.GET("/api/dashboard/last_study_session", dashboardHandler.GetLastStudySession)
	router.GET("/api/dashboard/study_progress", dashboardHandler.GetStudyProgress)
	router.GET("/api/dashboard/quick-stats", dashboardHandler.GetQuickStats)

	// Words endpoints
	wordHandler := handlers.NewWordHandler()
	router.GET("/api/words", wordHandler.GetAllWords)
	router.GET("/api/words/:id", wordHandler.GetWord)
	router.GET("/api/groups/:id/words", wordHandler.GetWordsByGroup)
	router.POST("/api/study_sessions/:word_id/review", wordHandler.AddWordReview)

	// Groups endpoints
	groupHandler := handlers.NewGroupHandler()
	router.GET("/api/groups", groupHandler.GetAllGroups)
	router.GET("/api/groups/:id", groupHandler.GetGroup)
	router.GET("/api/groups/:id/study_sessions", groupHandler.GetGroupStudySessions)

	// Study sessions endpoints
	studyHandler := handlers.NewStudyHandler()
	router.GET("/api/study_sessions", studyHandler.GetAllStudySessions)
	router.GET("/api/study_sessions/:id", studyHandler.GetStudySession)
	router.GET("/api/study_sessions/:id/words", studyHandler.GetStudySessionWords)
	router.POST("/api/reset_history", studyHandler.ResetHistory)
	router.POST("/api/full_reset", studyHandler.FullReset)

	// Study activities endpoints
	activityHandler := handlers.NewActivityHandler()
	router.GET("/api/api/study_activities/:id", activityHandler.GetActivity)
	router.GET("/api/api/study_activities/:id/study_sessions", activityHandler.GetActivitySessions)
	router.POST("/api/study_activities", activityHandler.CreateActivity)

	return router
} 