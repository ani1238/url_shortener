package main

import (
	"fmt"
	"net/http"

	"github.com/ani1238/url_shortener/redisdb"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	us := NewURLShortener()

	r.HandleFunc("/shorten", us.shortenURL).Methods(http.MethodPost)
	r.HandleFunc("/{id:[a-zA-Z0-9]+}", us.redirectLongURL)

	// Initialize the Redis client.
	redisdb.InitializeRedisClient()

	http.Handle("/", r)

	fmt.Println("URL Shortener started on :8080")
	http.ListenAndServe(":8080", nil)
}
