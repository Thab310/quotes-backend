package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

var quotes = []Quote{
	{Text: "The best way to predict the future is to create it.", Author: "Peter Drucker"},
	{Text: "Technology is best when it brings people together.", Author: "Matt Mullenweg"},
	{Text: "Innovation distinguishes between a leader and a follower.", Author: "Steve Jobs"},
}

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func getQuotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/api/quotes", enableCORS(getQuotes))

	log.Printf("Server starting on port %s!", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
