package api

import (
	"fmt"
	"net/http"
)

func Run() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Hello world!")
	})
	fmt.Printf("Starting server...")
	http.ListenAndServe(":8000", nil)
}