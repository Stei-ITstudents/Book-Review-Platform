package handlers

import (
	"booknook/api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// handles review creation
func CreateReview(w http.ResponseWriter, r *http.Request) {
	var review models.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := models.AddReview(&review); err != nil {
		http.Error(w, "Failed to add review", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// handles retrieving reviews
func GetReviews(w http.ResponseWriter, r *http.Request) {
	bookID, err := parseBookID(r)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	reviews, err := models.GetReviewsByBookID(bookID)
	if err != nil {
		http.Error(w, "Failed to retrieve reviews", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(reviews); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// extracts book ID from request
func parseBookID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return 0, fmt.Errorf("missing book ID")
	}
	return strconv.Atoi(id)
}
