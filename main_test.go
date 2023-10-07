package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

var us = NewURLShortener()

func TestGetTopDomains(t *testing.T) {
	req, err := http.NewRequest("GET", "/metrics/topdomains", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(us.getTopDomains)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Add more assertions based on the expected response
}

func TestShortenURL(t *testing.T) {
	payload := []byte(`{"url":"http://example.com"}`)
	req, err := http.NewRequest("POST", "/shorten", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(us.shortenURL)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Add assertions based on the expected response
}

func TestRedirectToLongURL(t *testing.T) {
	req, err := http.NewRequest("GET", "/shortened-url", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(us.redirectLongURL)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusFound {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusFound)
	}

	// Add more assertions based on the expected response
}
