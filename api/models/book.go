package models

import (
	"booknook/api/database"
	"fmt"
)

// represents a book entity
type Book struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Author        string  `json:"author"`
	Description   string  `json:"description"`
	CoverImage    string  `json:"cover_image"`
	AverageRating float64 `json:"average_rating"`
}

// retrieves a book by its ID
func GetBookByID(id int) (*Book, error) {
	var book Book
	query := `SELECT id, title, author, description, cover_image, average_rating 
              FROM books WHERE id = ?`
	err := database.DB.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author, &book.Description, &book.CoverImage, &book.AverageRating)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// retrieves all books from the database
func GetAllBooks() ([]Book, error) {
	rows, err := database.DB.Query(`SELECT id, title, author, description, cover_image, average_rating FROM books`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description, &book.CoverImage, &book.AverageRating); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// inserts a new book into the database
func AddBook(book *Book) error {
	_, err := database.DB.Exec(
		`INSERT INTO books (title, author, description, cover_image, average_rating) 
         VALUES (?, ?, ?, ?, ?)`,
		book.Title, book.Author, book.Description, book.CoverImage, book.AverageRating)
	return err
}

// updates an existing book in the database
func UpdateBook(book *Book) error {
	query := `UPDATE books SET title = ?, author = ?, description = ?, cover_image = ?, average_rating = ? WHERE id = ?`
	_, err := database.DB.Exec(query, book.Title, book.Author, book.Description, book.CoverImage, book.AverageRating, book.ID)
	if err != nil {
		return fmt.Errorf("could not update book: %v", err)
	}
	return nil
}

// deletes a book from the database by its ID
func DeleteBook(id int) error {
	query := "DELETE FROM books WHERE id = ?"
	_, err := database.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete book: %v", err)
	}
	return nil
}
