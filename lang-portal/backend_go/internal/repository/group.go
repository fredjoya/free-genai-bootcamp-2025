package repository

import (
	"database/sql"
	"fmt"
	"lang-portal/internal/models"
	"lang-portal/pkg/database"
	"lang-portal/pkg/pagination"
)

// GroupRepository handles database operations for groups
type GroupRepository struct {
	db *sql.DB
}

// NewGroupRepository creates a new group repository
func NewGroupRepository() *GroupRepository {
	return &GroupRepository{
		db: database.GetDB(),
	}
}

// GetAll retrieves all groups with pagination
func (r *GroupRepository) GetAll(page, itemsPerPage int) ([]models.Group, int, error) {
	offset := pagination.GetOffset(page, itemsPerPage)

	// Get total count
	var total int
	err := r.db.QueryRow("SELECT COUNT(*) FROM groups").Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("error getting total groups: %v", err)
	}

	// Get groups with word counts
	query := `
		SELECT g.id, g.name, g.created_at,
			COUNT(DISTINCT wg.word_id) as word_count
		FROM groups g
		LEFT JOIN word_groups wg ON g.id = wg.group_id
		GROUP BY g.id
		ORDER BY g.created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.Query(query, itemsPerPage, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("error querying groups: %v", err)
	}
	defer rows.Close()

	var groups []models.Group
	for rows.Next() {
		var group models.Group
		err := rows.Scan(
			&group.ID,
			&group.Name,
			&group.CreatedAt,
			&group.WordCount,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("error scanning group: %v", err)
		}
		groups = append(groups, group)
	}

	return groups, total, nil
}

// GetByID retrieves a group by ID
func (r *GroupRepository) GetByID(id int64) (*models.Group, error) {
	query := `
		SELECT g.id, g.name, g.created_at,
			COUNT(DISTINCT wg.word_id) as word_count
		FROM groups g
		LEFT JOIN word_groups wg ON g.id = wg.group_id
		WHERE g.id = ?
		GROUP BY g.id
	`

	var group models.Group
	err := r.db.QueryRow(query, id).Scan(
		&group.ID,
		&group.Name,
		&group.CreatedAt,
		&group.WordCount,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("error querying group: %v", err)
	}

	return &group, nil
}

// GetStudySessions retrieves study sessions for a group
func (r *GroupRepository) GetStudySessions(groupID int64, page, itemsPerPage int) ([]models.StudySession, int, error) {
	offset := pagination.GetOffset(page, itemsPerPage)

	// Get total count
	var total int
	err := r.db.QueryRow(`
		SELECT COUNT(*) FROM study_sessions
		WHERE group_id = ?
	`, groupID).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("error getting total sessions: %v", err)
	}

	// Get sessions with review counts
	query := `
		SELECT ss.id, ss.study_activity_id, ss.group_id, ss.created_at,
			g.name as group_name,
			COUNT(wri.id) as review_items_count
		FROM study_sessions ss
		JOIN groups g ON ss.group_id = g.id
		LEFT JOIN word_review_items wri ON ss.id = wri.study_session_id
		WHERE ss.group_id = ?
		GROUP BY ss.id
		ORDER BY ss.created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.Query(query, groupID, itemsPerPage, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("error querying sessions: %v", err)
	}
	defer rows.Close()

	var sessions []models.StudySession
	for rows.Next() {
		var session models.StudySession
		err := rows.Scan(
			&session.ID,
			&session.StudyActivityID,
			&session.GroupID,
			&session.CreatedAt,
			&session.GroupName,
			&session.ReviewItemsCount,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("error scanning session: %v", err)
		}
		sessions = append(sessions, session)
	}

	return sessions, total, nil
}

// GetLastStudySession retrieves the most recent study session for a group
func (r *GroupRepository) GetLastStudySession(groupID int64) (*models.StudySession, error) {
	query := `
		SELECT ss.id, ss.study_activity_id, ss.group_id, ss.created_at,
			g.name as group_name,
			COUNT(wri.id) as review_items_count
		FROM study_sessions ss
		JOIN groups g ON ss.group_id = g.id
		LEFT JOIN word_review_items wri ON ss.id = wri.study_session_id
		WHERE ss.group_id = ?
		GROUP BY ss.id
		ORDER BY ss.created_at DESC
		LIMIT 1
	`

	var session models.StudySession
	err := r.db.QueryRow(query, groupID).Scan(
		&session.ID,
		&session.StudyActivityID,
		&session.GroupID,
		&session.CreatedAt,
		&session.GroupName,
		&session.ReviewItemsCount,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, fmt.Errorf("error querying last session: %v", err)
	}

	return &session, nil
} 