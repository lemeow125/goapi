package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	goodbye "github.com/lemeow125/goapi/internal/api/goodbye"
	hello "github.com/lemeow125/goapi/internal/api/hello"
)

func Run() {
	r := mux.NewRouter()
	
	api := r.PathPrefix("/api/v1").Subrouter()

	hello.SetupRoutes(api)
	goodbye.SetupRoutes(api)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Up!")
	})

	fmt.Printf("Starting server...")
	http.ListenAndServe(":8000", r)
}