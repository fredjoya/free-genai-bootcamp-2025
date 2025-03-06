package handlers

import (
    "github.com/gin-gonic/gin"
    "lang-portal/pkg/database"
)

func GetWords(c *gin.Context) {
    rows, err := database.DB.Query("SELECT id, arabic, transliteration, english FROM words")
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var words []gin.H
    for rows.Next() {
        var id int
        var arabic, transliteration, english string
        if err := rows.Scan(&id, &arabic, &transliteration, &english); err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        words = append(words, gin.H{
            "id": id,
            "arabic": arabic,
            "transliteration": transliteration,
            "english": english,
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
} 