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
- Mage is a task runner for Go
- The API wil always return JSON
- There will be no authentication or authorization
- Everything will be treated as a single user

## Directory Structure

```test
lang-portal/backend-go/
├── cmd/
│   └── server/
│       └── main.go           # Application entry point
├── internal/
│   ├── api/
│   │   ├── handlers/        # HTTP request handlers
│   │   │   ├── dashboard.go
│   │   │   ├── groups.go
│   │   │   ├── study.go
│   │   │   └── words.go
│   │   ├── middleware/      # HTTP middleware
│   │   └── router.go        # Router setup
│   ├── models/             # Data models
│   │   ├── word.go
│   │   ├── group.go
│   │   └── study.go
│   ├── repository/         # Database operations
│   │   ├── word.go
│   │   ├── group.go
│   │   └── study.go
│   └── service/           # Business logic
│       ├── word.go
│       ├── group.go
│       └── study.go
├── db/                    # Database related files
│   ├── migrations/        # SQL migration files
│   │   ├── 0001_init.sql
│   │   └── 0002_create_tables.sql
│   └── seeds/            # Seed data files
│       └── words.json
├── pkg/                  # Shared packages
│   ├── database/        # Database connection and utilities
│   └── pagination/      # Pagination helpers
├── magefile.go          # Mage task definitions
├── go.mod              # Go module file
├── go.sum              # Go dependencies checksum
└── README.md           # Project documentation
```


## Database Schema 

Our database will be a single sqlite3 database called `words.db` that will be in the root of the project folder of `backend_go`.

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

### GET /api/api/study_activities/:id/study_sessions
    - pagination with 100 items per page
#### JSON Response
```json
{
  "items": [
    {
      "id": 123,
      "activity_name": "Vocabulary Quiz",
      "group_name": "Basic Greetings",
      "start_time": "2025-02-08T17:20:23-05:00",
      "end_time": "2025-02-08T17:30:23-05:00",
      "review_items_count": 20
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "total_items": 100,
    "items_per_page": 20
  }
}
```

### POST /api/study_activities
#### Request Params
    - group_id integer
    - study_activity_id integer

#### JSON Response
```json
{
  "id": 124,
  "group_id": 123
}
```

### GET /api/words  
    - pagination with 100 items per page

#### JSON Response
```json
{
  "items": [
    {
      "arabic": "مرحبا",
      "transliteration": "marhaba",
      "english": "hello",
      "correct_count": 5,
      "wrong_count": 2
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "total_items": 500,
    "items_per_page": 100
  }
}
```

### GET /api/words/:id  
#### JSON Response
```json
{
  "id": 1,
  "arabic": "مرحبا",
  "transliteration": "marhaba",
  "english": "hello",
  "correct_count": 5,
  "wrong_count": 2,
  "groups": [
    {
      "id": 1,
      "name": "Basic Greetings"
    },
    {
      "id": 2,
      "name": "Common Phrases"
    }
  ]
}
```

### GET /api/groups  
    - pagination with 100 items per page
#### JSON Response
```json
{
  "items": [
    {
      "id": 1,
      "name": "Basic Greetings",
      "word_count": 25
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "total_items": 500,
    "items_per_page": 100
  }
}
```

### GET /api/groups/:id  
#### JSON Response
```json
{
  "id": 1,
  "name": "Basic Greetings",
  "total_word_count": 25
}
```

### GET /api/groups/:id/words  
#### JSON Response
```json
{
  "items": [
    {
      "arabic": "مرحبا",
      "transliteration": "marhaba",
      "english": "hello",
      "correct_count": 5,
      "wrong_count": 2
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "total_items": 500,
    "items_per_page": 100
  }
}
```

### GET /api/groups/:id/study_sessions
#### JSON Response
```json
{
  "items": [
    {
      "id": 123,
      "activity_name": "Vocabulary Quiz",
      "group_name": "Basic Greetings",
      "start_time": "2025-02-08T17:20:23-05:00",
      "end_time": "2025-02-08T17:30:23-05:00",
      "review_items_count": 20
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "total_items": 100,
    "items_per_page": 20
  }
}
```

### GET /api/study_sessions
    -pagination with 100 items per page
#### JSON Response
```json
{
  "items": [
    {
      "id": 123,
      "activity_name": "Vocabulary Quiz",
      "group_name": "Basic Greetings",
      "start_time": "2025-02-08T17:20:23-05:00",
      "end_time": "2025-02-08T17:30:23-05:00",
      "review_items_count": 20
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "total_items": 100,
    "items_per_page": 100
  }
}
```

### GET /api/study_sessions/:id
#### JSON Response
```json
{
  "id": 123,
  "activity_name": "Vocabulary Quiz",
  "group_name": "Basic Greetings",
  "start_time": "2025-02-08T17:20:23-05:00",
  "end_time": "2025-02-08T17:30:23-05:00",
  "review_items_count": 20
}
```

### GET /api/study_sessions/:id/words
#### JSON Response
```json
{
  "items": [
    {
      "arabic": "مرحبا",
      "transliteration": "marhaba",
      "english": "hello",
      "correct_count": 5,
      "wrong_count": 2,
      "review_correct": true
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "total_items": 500,
    "items_per_page": 100
  }
}
```

### POST /api/reset_history
#### JSON Response
```json
{
  "success": true,
  "message": "Study history has been reset successfully"
}
```

### POST /api/full_reset
#### JSON Response
```json
{
  "success": true,
  "message": "Database has been reset and reseeded successfully"
}
```

### POST /api/study_sessions/:word_id/review
#### Request Params
    - id (study_session_id) integer
    - word_id integer
    - correct boolean

#### Request Payload
```json
{
  "correct": true
}
```

#### Request Payload
```json
{
  "success": true,
  "word_id": 1,
  "study_session_id": 123,
  "correct": true,
  "created_at": "2025-02-08T17:33:07-05:00"
}
```

## Task Runner Tasks

Mage is a task runner for Go.
Let's list out posssible tasks we need to for our lang portal.

### Initialize Database
This task will initialize the sqlite3 database called `words.db`.

### Migrate Database
This task will run a series of migration sql files on the database.

Migrations live in the `migrations` folder.
The migration files will be run in order of their file name.
The file names should looks like this:

```sql
0001_init.sql
0002_create_words_tab;le.sql
```

### Seed Data
This task will import json files and transform them into target data for our database.

All seed filed live in the `seeds` folder.
All seed files should be loaded.

In our task, we should have DSL to specify each seed file and its expected group word name.

```json
[
  {
    "arabic": "يَدْفَع",
    "transliteration": "yadfaʿ",
    "english": "to pay",
    "parts": [
      { "arabic": "يَ", "transliteration": ["ya"] },
      { "arabic": "دْ", "transliteration": ["d"] },
      { "arabic": "فَ", "transliteration": ["fa"] },
      { "arabic": "ع", "transliteration": ["ʿ"] }
    ]
  },
  ...
]
```