# Language Learning Portal Backend

This is the backend service for the Language Learning Portal, built with Go and SQLite3.

## Prerequisites

- Go 1.21 or later
- SQLite3
- Mage (Go task runner)

## Setup

1. Install Go:
   ```bash
   # On Ubuntu/Debian
   sudo snap install go --classic
   ```

2. Install Mage:
   ```bash
   go install github.com/magefile/mage@latest
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

## Project Structure

```
lang-portal/backend_go/
├── cmd/
│   └── server/
│       └── main.go           # Application entry point
├── internal/
│   ├── api/                 # HTTP handlers and middleware
│   ├── models/             # Data models
│   ├── repository/         # Database operations
│   └── service/           # Business logic
├── db/                    # Database related files
│   ├── migrations/        # SQL migration files
│   └── seeds/            # Seed data files
└── pkg/                  # Shared packages
```

## Database Management

The project uses Mage for database management tasks:

- Initialize database:
  ```bash
  mage dbinit
  ```

- Seed database with initial data:
  ```bash
  mage dbseed
  ```

- Reset database (remove and reinitialize):
  ```bash
  mage dbreset
  ```

## Running the Application

1. Start the server:
   ```bash
   go run cmd/server/main.go
   ```

2. The server will start on `http://localhost:8080`

## API Endpoints

### Dashboard
- GET `/api/dashboard/last_study_session`
- GET `/api/dashboard/study_progress`
- GET `/api/dashboard/quick-stats`

### Words
- GET `/api/words`
- GET `/api/words/:id`

### Groups
- GET `/api/groups`
- GET `/api/groups/:id`
- GET `/api/groups/:id/words`
- GET `/api/groups/:id/study_sessions`

### Study Sessions
- GET `/api/study_sessions`
- GET `/api/study_sessions/:id`
- GET `/api/study_sessions/:id/words`
- POST `/api/study_sessions/:word_id/review`
- POST `/api/reset_history`
- POST `/api/full_reset`

## Development

1. Add new migrations in `db/migrations/`
2. Add seed data in `db/seeds/`
3. Implement new handlers in `internal/api/handlers/`
4. Add business logic in `internal/service/`
5. Define data models in `internal/models/` 