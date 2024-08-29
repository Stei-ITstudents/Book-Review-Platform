package routes

import (
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "booknook/api/models"
    "booknook/api/utils"
)

// initializes the routes for the application
func DefineRoutes(router *mux.Router) {
    router.HandleFunc("/books", getBooks).Methods("GET")
    router.HandleFunc("/books/{id}", getBook).Methods("GET")
    router.HandleFunc("/books", createBook).Methods("POST")
    router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
    router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
}

// adds a health check endpoint to the router
func HealthCheckEndpoint(router *mux.Router) {
    router.HandleFunc("/health", healthCheck).Methods("GET")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
    books, err := models.GetAllBooks()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    utils.WriteJSON(w, http.StatusOK, books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }
    book, err := models.GetBookByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    utils.WriteJSON(w, http.StatusOK, book)
}

func createBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    if err := utils.ReadJSON(r, &book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := models.AddBook(&book); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    utils.WriteJSON(w, http.StatusCreated, book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }
    var book models.Book
    if err := utils.ReadJSON(r, &book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    book.ID = id
    if err := models.UpdateBook(&book); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    utils.WriteJSON(w, http.StatusOK, book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid book ID", http.StatusBadRequest)
        return
    }
    if err := models.DeleteBook(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}

