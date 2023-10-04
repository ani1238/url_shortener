package main

import "fmt"
import "sync"
import "net/http"

type URLShortener struct {
	mutex    sync.Mutex
	mappings map[string]string
	baseURL  string
}

func NewURLShortener() *URLShortener {
	return &URLShortener{
		mappings: make(map[string]string),
		baseURL:  "http://localhost:8080/",
	}
}

func (us *URLShortener) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	longURL := r.FormValue("url")

	us.mutex.Lock()
	defer us.mutex.Unlock()

	shortURL := generateShortURL()
	us.mappings[shortURL] = longURL

	fmt.Fprintf(w, "Shortened URL: %s%s\n", us.baseURL, shortURL)
}

func (us *URLShortener) Redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:] // Remove leading '/'

	us.mutex.Lock()
	defer us.mutex.Unlock()

	longURL, exists := us.mappings[shortURL]
	if !exists {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func generateShortURL() string {
	return "abc123"
}
