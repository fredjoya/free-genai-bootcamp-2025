#!/bin/bash

BASE_URL="http://localhost:8080/api"

# Test Dashboard endpoints
echo "Testing Dashboard endpoints..."
curl "${BASE_URL}/dashboard/last_study_session"
curl "${BASE_URL}/dashboard/study_progress"
curl "${BASE_URL}/dashboard/quick-stats"

# Test Study Activities endpoints
echo "Testing Study Activities endpoints..."
curl "${BASE_URL}/study_activities/1"
curl "${BASE_URL}/study_activities/1/study_sessions"
curl -X POST "${BASE_URL}/study_activities" -d '{"group_id": 1, "study_activity_id": 1}'

# Test Words endpoints
echo "Testing Words endpoints..."
curl "${BASE_URL}/words"
curl "${BASE_URL}/words/1"

# Test Groups endpoints
echo "Testing Groups endpoints..."
curl "${BASE_URL}/groups"
curl "${BASE_URL}/groups/1"
curl "${BASE_URL}/groups/1/words"
curl "${BASE_URL}/groups/1/study_sessions"

# Test Study Sessions endpoints
echo "Testing Study Sessions endpoints..."
curl "${BASE_URL}/study_sessions"
curl "${BASE_URL}/study_sessions/1"
curl "${BASE_URL}/study_sessions/1/words"
curl -X POST "${BASE_URL}/study_sessions/1/review" -d '{"word_id": 1, "correct": true}'

# Test Reset endpoints
echo "Testing Reset endpoints..."
curl -X POST "${BASE_URL}/reset_history"
curl -X POST "${BASE_URL}/full_reset" 