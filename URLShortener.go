package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/ani1238/url_shortener/redisdb"
	"github.com/gorilla/mux"
)

var (
	mutex      sync.Mutex
	counter    int64 = 0
	base64Char       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type URLShortener struct {
	baseURL string
}

func NewURLShortener() *URLShortener {
	return &URLShortener{
		baseURL: "http://localhost:8080/",
	}
}

func responseJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (us *URLShortener) shortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if input.URL == "" {
		http.Error(w, "No URL received", http.StatusBadRequest)
		return
	}

	// Generate a unique ID for the short URL.
	id, cnt := generateShortURLID()

	// Store the long URL in Redis.
	if err := redisdb.AddToRedis("counter", cnt, 365*24*time.Hour); err != nil {
		http.Error(w, "Failed to store counter in Redis", http.StatusInternalServerError)
		return
	}

	if err := redisdb.AddToRedis(id, input.URL, 24*time.Hour); err != nil {
		http.Error(w, "Failed to store URL in Redis", http.StatusInternalServerError)
		return
	}

	responseJSON(w, r, map[string]string{
		"shortened_url": us.baseURL + id,
		"error":         "",
	})
}

func (us *URLShortener) redirectLongURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Retrieve the long URL from Redis.
	longURL, err := redisdb.GetFromRedis(id)
	if err != nil {
		http.Error(w, "Not a valid ID", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

// Generate a unique short URL ID.
func generateShortURLID() (string, int64) {
	mutex.Lock()
	defer mutex.Unlock()

	counter++

	id := convertToBase64(counter)
	return id, counter
}

func convertToBase64(num int64) string {
	var encoded []byte
	for num > 0 {
		remainder := num % 64
		encoded = append(encoded, base64Char[remainder])
		num = num / 64
	}
	// Reverse the encoded characters to get the correct base64 representation.
	length := len(encoded)
	for i := 0; i < length/2; i++ {
		encoded[i], encoded[length-i-1] = encoded[length-i-1], encoded[i]
	}
	return string(encoded)
}
