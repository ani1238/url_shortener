package main

import "fmt"
import "sync"
import "net/http"

func init() {
	// Initialize the Redis client.
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

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

	shortURL := generateShortURL()

	// Store the mapping in Redis.
	err := redisClient.Set(context.Background(), shortURL, longURL, 24*time.Hour).Err()
	if err != nil {
		http.Error(w, "Failed to store URL in Redis", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Shortened URL: %s%s\n", us.baseURL, shortURL)
}

func (us *URLShortener) Redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]

	// Retrieve the long URL from Redis.
	longURL, err := redisClient.Get(context.Background(), shortURL).Result()
	if err == redis.Nil {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, "Failed to fetch URL from Redis", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func generateShortURL() string {
	return "abc123"
}
