-- Create words table
CREATE TABLE IF NOT EXISTS words (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    word TEXT NOT NULL,
    translation TEXT NOT NULL,
    example TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Create groups table
CREATE TABLE IF NOT EXISTS groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Create word_groups table (many-to-many relationship)
CREATE TABLE IF NOT EXISTS word_groups (
    word_id INTEGER,
    group_id INTEGER,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (word_id, group_id),
    FOREIGN KEY (word_id) REFERENCES words(id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);

-- Create study_sessions table
CREATE TABLE IF NOT EXISTS study_sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    group_id INTEGER,
    activity_id INTEGER,
    started_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    completed_at DATETIME,
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE SET NULL
);

-- Create word_reviews table
CREATE TABLE IF NOT EXISTS word_reviews (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    word_id INTEGER NOT NULL,
    study_session_id INTEGER NOT NULL,
    review_type TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (word_id) REFERENCES words(id) ON DELETE CASCADE,
    FOREIGN KEY (study_session_id) REFERENCES study_sessions(id) ON DELETE CASCADE
);

-- Create learning_activities table
CREATE TABLE IF NOT EXISTS learning_activities (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Create activity_sessions table (many-to-many relationship)
CREATE TABLE IF NOT EXISTS activity_sessions (
    activity_id INTEGER,
    study_session_id INTEGER,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (activity_id, study_session_id),
    FOREIGN KEY (activity_id) REFERENCES learning_activities(id) ON DELETE CASCADE,
    FOREIGN KEY (study_session_id) REFERENCES study_sessions(id) ON DELETE CASCADE
); 