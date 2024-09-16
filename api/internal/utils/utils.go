package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/CristyNel/booknook/api/models"
	"github.com/CristyNel/booknook/api/internal/database"
)

// writes JSON response
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("JSON encoding error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// reads JSON request body
func ReadJSON(r *http.Request, data interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	return decoder.Decode(data)
}

// retrieves book by ID
func GetBook(id string) (*models.Book, error) {
	var book models.Book
	query := `SELECT id, title, author, description, cover_image, average_rating FROM books WHERE id = ?`
	err := database.DB.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author, &book.Description, &book.CoverImage, &book.AverageRating)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// retrieves all books
func GetBooks() ([]models.Book, error) {
	rows, err := database.DB.Query(`SELECT id, title, author, description, cover_image, average_rating FROM books`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description, &book.CoverImage, &book.AverageRating); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// retrieves review by ID
func GetReview(id string) (*models.Review, error) {
	var review models.Review
	query := `SELECT id, book_id, user_id, rating, review_text FROM reviews WHERE id = ?`
	err := database.DB.QueryRow(query, id).Scan(&review.ID, &review.BookID, &review.UserID, &review.Rating, &review.ReviewText)
	if err != nil {
		return nil, err
	}
	return &review, nil
}

// retrieves reviews for book
func GetReviews(bookID string) ([]models.Review, error) {
	rows, err := database.DB.Query(`SELECT id, book_id, user_id, rating, review_text FROM reviews WHERE book_id = ?`, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.ID, &review.BookID, &review.UserID, &review.Rating, &review.ReviewText); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

// handles panics
func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("Panic recovered: %v", rec)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// logs requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// extracts book ID
func GetBookID(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["id"]
}
