package hello

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Hello %s!", name)
}

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/hello/{name}", Hello).Methods("GET")
}