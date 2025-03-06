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

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// API routes
	api := r.Group("/api")
	{
		// Dashboard routes
		dashboard := api.Group("/dashboard")
		{
			dashboard.GET("/last_study_session", dashboardHandler.GetLastStudySession)
			dashboard.GET("/study_progress", dashboardHandler.GetStudyProgress)
			dashboard.GET("/quick-stats", dashboardHandler.GetQuickStats)
		}

		// Words routes
		api.GET("/words", wordHandler.GetAllWords)
		api.GET("/words/:id", wordHandler.GetWord)
		api.GET("/groups/:id/words", wordHandler.GetWordsByGroup)
		api.POST("/study_sessions/:word_id/review", wordHandler.AddWordReview)

		// Groups routes
		api.GET("/groups", groupHandler.GetAllGroups)
		api.GET("/groups/:id", groupHandler.GetGroup)
		api.GET("/groups/:id/study_sessions", groupHandler.GetGroupStudySessions)

		// Study routes
		api.GET("/study_sessions", studyHandler.GetAllStudySessions)
		api.GET("/study_sessions/:id", studyHandler.GetStudySession)
		api.GET("/study_sessions/:id/words", studyHandler.GetStudySessionWords)
		api.POST("/study_sessions", studyHandler.CreateStudySession)
		api.POST("/reset_history", studyHandler.ResetHistory)
		api.POST("/full_reset", studyHandler.FullReset)

		// Learning Activities routes
		api.GET("/study_activities/:id", activityHandler.GetLearningActivity)
		api.GET("/study_activities/:id/study_sessions", activityHandler.GetLearningActivitySessions)
		api.POST("/study_activities", activityHandler.CreateLearningActivity)
	}

	return r
} 