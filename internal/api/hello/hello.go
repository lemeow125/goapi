package hello

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Hello world!")
	})
}