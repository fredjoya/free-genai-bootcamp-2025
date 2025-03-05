# Backend Technical Specs

## Business Goal: 

A language learning school wants to build a prototype of learning portal which will act as three things:
- Inventory of possible vocabulary that can be learned
- Act as a  Learning record store (LRS), providing correct and wrong score on practice vocabulary
- A unified launchpad to launch different learning apps


## Technical Requirements

- The backend will be built suing Go
- The database will be SQLite3
- The API will be built using Gin
- The API wil always return JSON
- There will be no authentication or authorization
- Everything will be treated as a single user

## Database Schema 

We have the follwoing tables:

- words -  stored vocabulary words
    - id integer
    - arabic string
    - transliteration string
    - english string
    - parts json 
- word_groups - join table for words and groups many-to-many
    - id integer
    - word_id integer
    - group_id integer
- groups - thematic groups of words
    - id integer
    - name string
- study_sessions - records of study sessions grouping word_review_items
    - id integer
    - study_activity_id integer
    - group_id integer
    - created_at datetime
- study_activites - a specific study activity linking a study session to group
    - id integer
    - study_session_id integer
    - group_id integer
    - created_at datetime
- word_review_items - a record of word practice determining if the word was correct or not
    - word_id integer
    - study_session_id integer
    - correct boolean 
    - created_at datetime

## API Endpoints

### GET /api/dashboard/last_study_session  
Returns infomrmation about the most recent study session.

#### JSON Response
```json
{
    "id": 123,
    "group_id": 456,
    "created_at": "2025-02-08T17:20:23-05:00",
    "study_activity_id": 789,
    "group_id": 456,
    "group_name": "Basic Greetings"
}
```

### GET /api/dashboard/study_progress  
Returns study progress statistics.
Please note that the frontend will determine progress bar based on total words studied and total available words. 

#### JSON Response
```json
{
  "total_words_studied": 3,
  "total_available_words": 124
}
```

### GET /api/dashboard/quick-stats  
Returns quick overview statistics. 

#### JSON Response
```json
{
  "success_rate": 80.0,
  "total_study_sessions": 4,
  "total_active_groups": 3,
  "study_streak_days": 4
}
```

### GET /api/api/study_activities/:id

#### JSON Response
```json
{
  "id": 1,
  "name": "Vocabulary Quiz",
  "thumbnail_url": "https://example.com/thumbnail.jpg",
  "description": "Practice your vocabulary with flashcards"
}
```

- GET /api/api/study_activities/:id/study_sessions

- POST /api/study_activities
    - required params: group_id, study_activity_id

- GET /api/words  
    - pagination with 100 items per page
- GET /api/words/:id  
- GET /api/groups  
    - pagination with 100 items per page
- GET /api/groups/:id  
- GET /api/groups/:id/words  
- GET /api/groups/:id/study_sessions
- GET /api/study_sessions
    -pagination with 100 items per page
- GET /api/study_sessions/:id
- GET /api/study_sessions/:id/words

- POST /api/reset_history
- POST /api/full_reset
- POST /api/study_sessions/:word_id/review
    - required params: correct