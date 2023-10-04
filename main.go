package main

import (
	"fmt"
	"net/http"
)

func main() {
	us := NewURLShortener()

	http.HandleFunc("/shorten", us.ShortenURL)
	http.HandleFunc("/", us.Redirect)

	fmt.Println("URL Shortener started on :8080")
	http.ListenAndServe(":8080", nil)
}
