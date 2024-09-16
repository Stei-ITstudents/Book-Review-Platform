-- create booknook batabase if doesn't exist
CREATE DATABASE IF NOT EXISTS booknook;
-- use booknook database
USE booknook;
-- books table
CREATE TABLE IF NOT EXISTS books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    published_date DATE,
    genre VARCHAR(100),
    summary TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- review table
CREATE TABLE IF NOT EXISTS reviews (
    id INT AUTO_INCREMENT PRIMARY KEY,
    book_id INT,
    user_id VARCHAR(255),
    rating INT CHECK (rating BETWEEN 1 AND 5),
    review TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
);

-- initial data
INSERT IGNORE INTO books (title, author, published_date, genre, summary)
VALUES
    ('100 Page Python Intro: Beyond the Basic Stuff with Python', 'Sundeep Agarwal', '2019-01-01', 'Programming, Python', 'A concise introduction to Python programming.'),
    ('JavaScript Notes for Professionals', 'Stack Overflow Community', '2021-01-01', 'Programming, JavaScript', 'A comprehensive reference guide to JavaScript.'),
    ('Essential Node.js', 'Kevin Skoglund', '2016-01-01', 'Programming, Node.js', 'A practical introduction to Node.js.'),
    ('Adaptive Web Design', 'Ethan Marcotte', '2010-01-01', 'Web Design, Responsive Design', 'A guide to creating websites that adapt to different screen sizes.')
;

-- Insert values into reviews
INSERT INTO reviews (book_id, user_id, rating, review)
VALUES
    ((SELECT id FROM books WHERE title = '100 Page Python Intro: Beyond the Basic Stuff with Python' LIMIT 1), 1, 5, 'A concise and well-explained guide to Python for beginners.'),
    ((SELECT id FROM books WHERE title = 'JavaScript Notes for Professionals' LIMIT 1), 2, 4, 'Covers advanced topics in JavaScript, ideal for intermediate learners.'),
    ((SELECT id FROM books WHERE title = 'Essential Node.js' LIMIT 1), 3, 5, 'A focused guide to Node.js covering core concepts and practical applications.'),
    ((SELECT id FROM books WHERE title = 'Adaptive Web Design' LIMIT 1), 4, 4, 'A practical guide to creating responsive web designs that adapt to different screen sizes.'),
    ((SELECT id FROM books WHERE title = 'Adaptive Web Design' LIMIT 1), 5, 5, 'Explores adaptive design techniques for creating accessible websites, with insights for enhancing user experience across devices.')
;