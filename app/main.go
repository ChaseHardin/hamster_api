package main

import (
	"github.com/gorilla/mux"
	"hamster_api/app/features/hello_world"
	"log"
	"net/http"
)

func main () {
	router := mux.NewRouter()

	router.HandleFunc("/greet", hello_world.GreetingHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":4020", router))
}
