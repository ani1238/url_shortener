package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ani1238/url_shortener/redisdb"
	"github.com/gorilla/mux"
)

func init() {
	// Initialize the Redis client
	redisdb.InitializeRedisClient()

	// Set the initial value for the shortened URL count
	if err := redisdb.SetInitialShortenedURLCount(); err != nil {
		log.Fatalf("Failed to set initial shortened URL count in Redis: %v", err)
	}
}

func main() {
	r := mux.NewRouter()
	us := NewURLShortener()

	r.HandleFunc("/shortenurl", us.shortenURL).Methods(http.MethodPost)
	r.HandleFunc("/metrics/topdomains", us.getTopDomains)
	r.HandleFunc("/{id:[a-zA-Z0-9]+}", us.redirectLongURL)

	http.Handle("/", r)

	fmt.Println("URL Shortener started on :8080")
	http.ListenAndServe(":8080", nil)
}
