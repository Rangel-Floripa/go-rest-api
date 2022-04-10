package routes

import (
	controlles "api-go-rest2/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controlles.Home)
	r.HandleFunc("/api/personalidades", controlles.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controlles.RetornaUmaPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
