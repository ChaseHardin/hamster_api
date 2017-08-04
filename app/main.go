package main

import (
	"github.com/gorilla/mux"
	"hamster_api/app/features/helloWorld"
	"log"
	"net/http"
)

func main () {
	router := mux.NewRouter()

	router.HandleFunc("/greet", helloWorld.GreetingHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":4020", router))
}
