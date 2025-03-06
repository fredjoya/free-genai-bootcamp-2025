package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"lang-portal/pkg/database"
	"strconv"
)

func main() {
	log.Println("Starting application...")
	
	if err := database.InitDB(); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	r := gin.Default()
	api := r.Group("/api")

	// Dashboard endpoints
	api.GET("/dashboard/last_study_session", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"id": 123,
			"group_id": 456,
			"created_at": "2025-02-08T17:20:23-05:00",
			"study_activity_id": 789,
			"group_name": "Basic Greetings",
		})
	})

	api.GET("/dashboard/study_progress", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"total_words_studied": 3,
			"total_available_words": 124,
		})
	})

	api.GET("/dashboard/quick-stats", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success_rate": 80.0,
			"total_study_sessions": 4,
			"total_active_groups": 3,
			"study_streak_days": 4,
		})
	})

	// Words endpoints
	api.GET("/words", func(c *gin.Context) {
		rows, err := database.DB.Query(`
			SELECT w.id, w.arabic, w.transliteration, w.english,
				   COUNT(CASE WHEN wri.correct = 1 THEN 1 END) as correct_count,
				   COUNT(CASE WHEN wri.correct = 0 THEN 1 END) as wrong_count
			FROM words w
			LEFT JOIN word_review_items wri ON w.id = wri.word_id
			GROUP BY w.id
		`)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var words []gin.H
		for rows.Next() {
			var id int
			var arabic, transliteration, english string
			var correct, wrong int
			if err := rows.Scan(&id, &arabic, &transliteration, &english, &correct, &wrong); err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			words = append(words, gin.H{
				"id": id,
				"arabic": arabic,
				"transliteration": transliteration,
				"english": english,
				"correct_count": correct,
				"wrong_count": wrong,
			})
		}

		c.JSON(200, gin.H{
			"items": words,
			"pagination": gin.H{
				"current_page": 1,
				"total_pages": 1,
				"total_items": len(words),
				"items_per_page": 100,
			},
		})
	})

	api.GET("/words/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for _, word := range database.MockWords {
			if word["id"].(int) == id {
				c.JSON(200, word)
				return
			}
		}
		c.JSON(404, gin.H{"error": "Word not found"})
	})

	// Groups endpoints
	api.GET("/groups", func(c *gin.Context) {
		rows, err := database.DB.Query(`
			SELECT g.id, g.name, COUNT(wg.word_id) as word_count
			FROM groups g
			LEFT JOIN word_groups wg ON g.id = wg.group_id
			GROUP BY g.id
		`)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var groups []gin.H
		for rows.Next() {
			var id int
			var name string
			var wordCount int
			if err := rows.Scan(&id, &name, &wordCount); err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			groups = append(groups, gin.H{
				"id": id,
				"name": name,
				"word_count": wordCount,
			})
		}

		c.JSON(200, gin.H{
			"items": groups,
			"pagination": gin.H{
				"current_page": 1,
				"total_pages": 1,
				"total_items": len(groups),
				"items_per_page": 100,
			},
		})
	})

	api.GET("/groups/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for _, group := range database.MockGroups {
			if group["id"].(int) == id {
				c.JSON(200, group)
				return
			}
		}
		c.JSON(404, gin.H{"error": "Group not found"})
	})

	// Study sessions endpoints
	api.POST("/study_sessions", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"id": 124,
			"group_id": 123,
		})
	})

	api.GET("/study_sessions/:id/words", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"items": database.MockWords,
			"pagination": gin.H{
				"current_page": 1,
				"total_pages": 1,
				"total_items": len(database.MockWords),
				"items_per_page": 100,
			},
		})
	})

	// Reset endpoints
	api.POST("/reset_history", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"message": "Study history has been reset successfully",
		})
	})

	api.POST("/full_reset", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"message": "Database has been reset and reseeded successfully",
		})
	})

	api.GET("/test", func(c *gin.Context) {
		var count int
		err := database.DB.QueryRow("SELECT COUNT(*) FROM words").Scan(&count)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"word_count": count})
	})

	log.Println("Starting server on :8080")
	r.Run(":8080")
} 