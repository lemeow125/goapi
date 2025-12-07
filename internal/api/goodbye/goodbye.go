package goodbye

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Goodbye(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Goodbye %s!", name)
}

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/goodbye/{name}", Goodbye).Methods("GET")
}