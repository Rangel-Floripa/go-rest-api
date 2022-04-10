package routes

import (
	"api-go-rest2/controllers"
	"log"
	"net/http"
)

func HandleResquest() {
	http.HandleFunc("/", controlles.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
