-- Create library schema
CREATE SCHEMA IF NOT EXISTS library;

-- Create books table
CREATE TABLE IF NOT EXISTS library.books (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    publish_year INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at TIMESTAMP
);

-- Create index for title column
CREATE INDEX idx_books_title
ON library.books (title);

-- Create index for author column
CREATE INDEX idx_books_author
ON library.books (author);

-- insert books data as seeder
INSERT INTO library.books (title, author, publish_year) VALUES 
('To Kill a Mockingbird', 'Harper Lee', 1960),
('1984', 'George Orwell', 1949),
('Pride and Prejudice', 'Jane Austen', 1813),
('The Great Gatsby', 'F. Scott Fitzgerald', 1925),
('Moby-Dick', 'Herman Melville', 1851),
('War and Peace', 'Leo Tolstoy', 1869),
('The Catcher in the Rye', 'J.D. Salinger', 1951),
('The Hobbit', 'J.R.R. Tolkien', 1937),
('Fahrenheit 451', 'Ray Bradbury', 1953),
('The Lord of the Rings', 'J.R.R. Tolkien', 1954);