-- Insert test groups
INSERT INTO groups (name) VALUES
('Basic Greetings'),
('Common Phrases'),
('Numbers');

-- Insert test words
INSERT INTO words (arabic, transliteration, english, parts) VALUES
('مرحبا', 'marhaba', 'hello', '{"parts": [{"arabic": "م", "transliteration": ["m"]}, {"arabic": "ر", "transliteration": ["r"]}, {"arabic": "ح", "transliteration": ["h"]}, {"arabic": "ب", "transliteration": ["b"]}, {"arabic": "ا", "transliteration": ["a"]}]}'),
('كيف حالك', 'kayf halak', 'how are you', '{"parts": [{"arabic": "ك", "transliteration": ["k"]}, {"arabic": "ي", "transliteration": ["y"]}, {"arabic": "ف", "transliteration": ["f"]}, {"arabic": " ", "transliteration": [" "]}, {"arabic": "ح", "transliteration": ["h"]}, {"arabic": "ا", "transliteration": ["a"]}, {"arabic": "ل", "transliteration": ["l"]}, {"arabic": "ك", "transliteration": ["k"]}]}'),
('شكرا', 'shukran', 'thank you', '{"parts": [{"arabic": "ش", "transliteration": ["sh"]}, {"arabic": "ك", "transliteration": ["k"]}, {"arabic": "ر", "transliteration": ["r"]}, {"arabic": "ا", "transliteration": ["a"]}, {"arabic": "ن", "transliteration": ["n"]}]}');

-- Link words to groups
INSERT INTO word_groups (word_id, group_id) VALUES
(1, 1), -- مرحبا -> Basic Greetings
(2, 1), -- كيف حالك -> Basic Greetings
(3, 2); -- شكرا -> Common Phrases

-- Insert test study activity
INSERT INTO study_activities (name, thumbnail_url, description) VALUES
('Vocabulary Quiz', 'https://example.com/quiz.jpg', 'Practice your vocabulary with flashcards');

-- Insert test study session
INSERT INTO study_sessions (study_activity_id, group_id, created_at, end_time) VALUES
(1, 1, datetime('now', '-1 hour'), datetime('now', '-30 minutes'));

-- Insert test word reviews
INSERT INTO word_review_items (word_id, study_session_id, correct, created_at) VALUES
(1, 1, 1, datetime('now', '-45 minutes')), -- Correct review for مرحبا
(2, 1, 0, datetime('now', '-40 minutes')), -- Wrong review for كيف حالك
(3, 1, 1, datetime('now', '-35 minutes')); -- Correct review for شكرا 