package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	books "github.com/lemeow125/goapi/internal/api/books"
	goodbye "github.com/lemeow125/goapi/internal/api/goodbye"
	hello "github.com/lemeow125/goapi/internal/api/hello"
)

func Run(db *sqlx.DB) {
	PORT := os.Getenv("BACKEND_PORT")
	r := mux.NewRouter()
	
	api := r.PathPrefix("/api/v1").Subrouter()

	hello.SetupRoutes(api)
	goodbye.SetupRoutes(api)
	books.SetupRoutes(api, db)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Up!")
	})

	fmt.Printf("Starting server...")
	http.ListenAndServe(":"+PORT, r)
}