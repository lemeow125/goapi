package books

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type Book struct {
    ID     int    `db:"id" json:"id"`
    Title  string `db:"title" json:"title"`
    Author string `db:"author" json:"author"`
}

type CreateBook struct {
    Title  string `db:"title" json:"title"`
    Author string `db:"author" json:"author"`
}

type ErrorResponse struct {
    Error string `json:"error"`
}

func SetupRoutes(r *mux.Router, db *sqlx.DB) {
	// GET Books
	r.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		log.Println("/books")

		rows, err := db.Query("SELECT * FROM Books")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}

		defer rows.Close()

        books := []Book{}
        for rows.Next() {
            var book Book
            err := rows.Scan(&book.ID, &book.Title, &book.Author)
            if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
				return
            }
            books = append(books, book)
        }

        json.NewEncoder(w).Encode(books)

	}).Methods("GET")

	// GET Book
	r.HandleFunc("/books/{title}", func(w http.ResponseWriter, r *http.Request){
        w.Header().Set("Content-Type", "application/json")
		log.Println("/books/{title}")

		vars := mux.Vars(r)
		title := vars["title"]

		rows, err := db.Query("SELECT * FROM Books where Title = ?", title)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}

		defer rows.Close()

		books := []Book{}
        for rows.Next() {
            var book Book
            err := rows.Scan(&book.ID, &book.Title, &book.Author)
            if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
				return
            }
            books = append(books, book)
        }

        json.NewEncoder(w).Encode(books)

	}).Methods("GET")

	// POST Book
	r.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		var createBook CreateBook
		
		err := json.NewDecoder(r.Body).Decode(&createBook)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}

		// DB Transaction
		tx, err := db.Begin()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			return
		}
		
		// Transaction Rollback
		var txErr error
		defer func() {
			if txErr != nil {
				tx.Rollback()
			}
		}()

		// INSERT with RETURNING - use proper placeholders and scan all fields
		var book Book
		txErr = tx.QueryRow(
			"INSERT INTO Books(title, author) VALUES($1, $2) RETURNING id, title, author", 
			createBook.Title, createBook.Author,
		).Scan(&book.ID, &book.Title, &book.Author)
		
		if txErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: txErr.Error()})
			return
		}
		
		// Commit Transaction
		txErr = tx.Commit()
		if txErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Error: txErr.Error()})
			return
		}

		// Response
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(book)
	}).Methods("POST")
}