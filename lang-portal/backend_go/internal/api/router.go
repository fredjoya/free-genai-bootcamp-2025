package api

import (
	"lang-portal/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and returns the Gin router with all routes configured
func SetupRouter() *gin.Engine {
	// Initialize handlers
	wordHandler := handlers.NewWordHandler()
	groupHandler := handlers.NewGroupHandler()
	studyHandler := handlers.NewStudyHandler()
	dashboardHandler := handlers.NewDashboardHandler()
	activityHandler := handlers.NewLearningActivityHandler()

	r := gin.Default()

	// Health check endpoint
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Dashboard endpoints
	r.GET("/api/dashboard/last_study_session", dashboardHandler.GetLastStudySession)
	r.GET("/api/dashboard/study_progress", dashboardHandler.GetStudyProgress)
	r.GET("/api/dashboard/quick-stats", dashboardHandler.GetQuickStats)

	// Words endpoints
	r.GET("/api/words", wordHandler.GetAllWords)
	r.GET("/api/words/:id", wordHandler.GetWord)
	r.GET("/api/groups/:id/words", wordHandler.GetWordsByGroup)
	r.POST("/api/study_sessions/:word_id/review", wordHandler.AddWordReview)

	// Groups endpoints
	r.GET("/api/groups", groupHandler.GetAllGroups)
	r.GET("/api/groups/:id", groupHandler.GetGroup)
	r.GET("/api/groups/:id/study_sessions", groupHandler.GetGroupStudySessions)

	// Study Sessions endpoints
	r.GET("/api/study_sessions", studyHandler.GetAllStudySessions)
	r.GET("/api/study_sessions/:id", studyHandler.GetStudySession)
	r.GET("/api/study_sessions/:id/words", studyHandler.GetStudySessionWords)
	r.POST("/api/study_sessions", studyHandler.CreateStudySession)
	r.POST("/api/reset_history", studyHandler.ResetHistory)
	r.POST("/api/full_reset", studyHandler.FullReset)

	// Learning Activities routes
	r.GET("/api/study_activities/:id", activityHandler.GetLearningActivity)
	r.GET("/api/study_activities/:id/study_sessions", activityHandler.GetLearningActivitySessions)
	r.POST("/api/study_activities", activityHandler.CreateLearningActivity)

	return r
} 