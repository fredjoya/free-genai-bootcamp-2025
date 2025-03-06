 #!/bin/bash

BASE_URL="http://localhost:8080/api"
echo "Testing API endpoints..."

# Function to make curl requests and format output
test_endpoint() {
    local method=$1
    local endpoint=$2
    local data=$3
    echo -e "\n=== Testing $method $endpoint ==="
    if [ -n "$data" ]; then
        curl -X $method "${BASE_URL}${endpoint}" \
            -H "Content-Type: application/json" \
            -d "$data" \
            -w "\nStatus: %{http_code}\n"
    else
        curl -X $method "${BASE_URL}${endpoint}" \
            -w "\nStatus: %{http_code}\n"
    fi
}

# Dashboard endpoints
test_endpoint "GET" "/dashboard/last_study_session"
test_endpoint "GET" "/dashboard/study_progress"
test_endpoint "GET" "/dashboard/quick-stats"

# Words endpoints
test_endpoint "GET" "/words"
test_endpoint "GET" "/words/1"

# Groups endpoints
test_endpoint "GET" "/groups"
test_endpoint "GET" "/groups/1"
test_endpoint "GET" "/groups/1/words"
test_endpoint "GET" "/groups/1/study_sessions"

# Study activities endpoints
test_endpoint "GET" "/study_activities/1"
test_endpoint "GET" "/study_activities/1/study_sessions"
test_endpoint "POST" "/study_activities" '{"group_id": 1, "study_activity_id": 1}'

# Study sessions endpoints
test_endpoint "GET" "/study_sessions"
test_endpoint "GET" "/study_sessions/1"
test_endpoint "GET" "/study_sessions/1/words"
test_endpoint "POST" "/study_sessions/1/review" '{"word_id": 1, "correct": true}'

# Reset endpoints
test_endpoint "POST" "/reset_history"
test_endpoint "POST" "/full_reset" 