package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	if err := initDB(); err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	initKafka()

	r := mux.NewRouter()
	r.HandleFunc("/messages", createMessageHandler).Methods("POST")
	r.HandleFunc("/stats", getStatsHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
