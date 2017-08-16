package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hamsterApi/app/features/helloWorld"
	"github.com/hamsterApi/app/features/places"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/greet", helloWorld.GreetingHandler).Methods("GET")
	router.HandleFunc("/places", places.GetPlaces).Methods("GET")

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "4020"
	}

	fmt.Println("Listening at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
