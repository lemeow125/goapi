package goodbye

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/goodbye/{name}", func(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		name := vars["name"]
		fmt.Fprintf(w, "Goodbye %s!", name)
	})
}