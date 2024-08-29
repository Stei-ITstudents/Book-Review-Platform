-- book table
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
INSERT INTO books (title, author, published_date, genre, summary)
VALUES
    ('The Great Gatsby', 'F. Scott Fitzgerald', '1925-04-10', 'Fiction', 'A novel set in the Roaring Twenties.'),
    ('To Kill a Mockingbird', 'Harper Lee', '1960-07-11', 'Fiction', 'A story of racial injustice in the Deep South.'),
    ('1984', 'George Orwell', '1949-06-08', 'Dystopian', 'A dystopian social science fiction novel.');

-- Insert values into reviews
INSERT INTO reviews (book_id, user_id, rating, review)
VALUES
    (1, 'user123', 5, 'An amazing read with deep themes.'),
    (2, 'user456', 4, 'A powerful novel with strong characters.'),
    (3, 'user789', 5, 'A chilling portrayal of a dystopian future.');
    