-- Insert test words
INSERT INTO words (word, translation, example) VALUES
('hello', 'hola', 'Hello, how are you?'),
('world', 'mundo', 'Hello, world!'),
('goodbye', 'adiós', 'Goodbye, see you later!'),
('thank you', 'gracias', 'Thank you very much!'),
('please', 'por favor', 'Please, help me.');

-- Insert test groups
INSERT INTO groups (name, description) VALUES
('Basic Phrases', 'Common everyday phrases'),
('Greetings', 'Different ways to greet people'),
('Numbers', 'Counting from 1 to 10');

-- Insert word-group relationships
INSERT INTO word_groups (word_id, group_id) VALUES
(1, 1), -- hello in Basic Phrases
(1, 2), -- hello in Greetings
(2, 1), -- world in Basic Phrases
(3, 1), -- goodbye in Basic Phrases
(3, 2), -- goodbye in Greetings
(4, 1), -- thank you in Basic Phrases
(5, 1); -- please in Basic Phrases

-- Insert learning activities
INSERT INTO learning_activities (name, description) VALUES
('Flashcards', 'Review words using flashcards'),
('Multiple Choice', 'Choose the correct translation'),
('Writing Practice', 'Practice writing the words');

-- Insert test study session
INSERT INTO study_sessions (group_id, started_at, completed_at) VALUES
(1, datetime('now', '-1 hour'), datetime('now', '-30 minutes'));

-- Insert word reviews
INSERT INTO word_reviews (word_id, study_session_id, review_type) VALUES
(1, 1, 'correct'),
(2, 1, 'correct'),
(3, 1, 'incorrect'); 