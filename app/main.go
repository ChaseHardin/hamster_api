package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"hamsterApi/app/features/helloWorld"
)

func main () {
	router := mux.NewRouter()

	router.HandleFunc("/greet", helloWorld.GreetingHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":4020", router))
}
