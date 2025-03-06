package repository

import (
	"database/sql"
	"fmt"

	"github.com/your-project/models"
)

type DashboardRepository struct {
	db *sql.DB
}

func (r *DashboardRepository) GetDashboardStats() (*models.DashboardStats, error) {
	query := `
		SELECT
			(SELECT COUNT(*) FROM words) as total_words,
			(SELECT COUNT(*) FROM groups) as total_groups,
			(SELECT COUNT(*) FROM study_sessions) as total_sessions,
			(SELECT COUNT(*) FROM word_reviews WHERE created_at >= datetime('now', '-7 days')) as recent_activity
	`
	
	stats := &models.DashboardStats{}
	err := r.db.QueryRow(query).Scan(
		&stats.TotalWords,
		&stats.TotalGroups,
		&stats.TotalStudySessions,
		&stats.RecentActivity,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting dashboard stats: %v", err)
	}

	return stats, nil
} 