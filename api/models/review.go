package models

import (
	"net/http"
	"encoding/json"
	"github.com/CristyNel/booknook/api/internal/database"
)

// represents a book review
type Review struct {
	ID         int    `json:"id"`
	BookID     int    `json:"book_id"`
	UserID     int    `json:"user_id"`
	Rating     int    `json:"rating"`
	ReviewText string `json:"review_text"`
}

// fetches reviews
func GetReviewsByBookID(bookID int) ([]Review, error) {
	rows, err := database.DB.Query(
		`SELECT id, book_id, user_id, rating, review_text 
		FROM reviews WHERE book_id = ?`, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []Review
	for rows.Next() {
		var review Review
		if err := rows.Scan(&review.ID, &review.BookID, &review.UserID, &review.Rating, &review.ReviewText); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

// inserts a new review
func AddReview(review *Review) error {
	_, err := database.DB.Exec(
		`INSERT INTO reviews (book_id, user_id, rating, review_text) 
		VALUES (?, ?, ?, ?)`,
		review.BookID, review.UserID, review.Rating, review.ReviewText)
	return err
}

// review creation
func CreateReview(w http.ResponseWriter, r *http.Request) {
	var review Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Example validation (optional)
	if review.Rating < 1 || review.Rating > 5 {
		http.Error(w, "Rating must be between 1 and 5", http.StatusBadRequest)
		return
	}

	if err := AddReview(&review); err != nil {
		http.Error(w, "Failed to add review", http.StatusInternalServerError)
		return
	}

	// Return the created review object or ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(review); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
