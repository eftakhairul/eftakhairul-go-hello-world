package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Okay bool `json:"okay"`
}

func main() {
	log.Printf("Starting server at port 8080\n")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", rootHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/health/live", healthHandler)
	http.HandleFunc("/health/ready", healthHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Response{Okay: true})
	return
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
