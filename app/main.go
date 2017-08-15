package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hamsterApi/app/features/helloWorld"
	"github.com/hamsterApi/app/features/places"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/greet", helloWorld.GreetingHandler).Methods("GET")
	router.HandleFunc("/places", places.GetPlaces).Methods("GET")

	log.Fatal(http.ListenAndServe(":4020", router))
}
