package Server

import (
	"net/http"
	"fmt"
)

func initRoutes(){

	http.HandleFunc("/", index)

	http.HandleFunc("/countries", func (w http.ResponseWriter, r *http.Request){
		switch r.Method{
			case http.MethodGet:
				getCountries(w, r)
			case http.MethodPost:
				addCountry(w, r)
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "method not allowed")
		}
	})
}