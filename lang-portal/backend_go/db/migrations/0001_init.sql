-- Create groups table
CREATE TABLE IF NOT EXISTS groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Create words table
CREATE TABLE IF NOT EXISTS words (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    arabic TEXT NOT NULL,
    transliteration TEXT NOT NULL,
    english TEXT NOT NULL,
    parts JSON,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Create word_groups table (many-to-many relationship)
CREATE TABLE IF NOT EXISTS word_groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    word_id INTEGER NOT NULL,
    group_id INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (word_id) REFERENCES words(id),
    FOREIGN KEY (group_id) REFERENCES groups(id),
    UNIQUE(word_id, group_id)
);

-- Create study_activities table
CREATE TABLE IF NOT EXISTS study_activities (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    thumbnail_url TEXT,
    description TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Create study_sessions table
CREATE TABLE IF NOT EXISTS study_sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    study_activity_id INTEGER NOT NULL,
    group_id INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    end_time DATETIME,
    FOREIGN KEY (study_activity_id) REFERENCES study_activities(id),
    FOREIGN KEY (group_id) REFERENCES groups(id)
);

-- Create word_review_items table
CREATE TABLE IF NOT EXISTS word_review_items (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    word_id INTEGER NOT NULL,
    study_session_id INTEGER NOT NULL,
    correct BOOLEAN NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (word_id) REFERENCES words(id),
    FOREIGN KEY (study_session_id) REFERENCES study_sessions(id)
); 