package books

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	// GET Book
	r.HandleFunc("/books/{name}", func(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		name := vars["name"]
		fmt.Fprintf(w, "Book %s", name)
	}).Methods("GET")
}