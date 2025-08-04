package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAddHandler_ValidInput(t *testing.T) {
	form := url.Values{}
	form.Add("a", "3")
	form.Add("b", "5")

	req := httptest.NewRequest(http.MethodPost, "/api/add", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", rr.Code)
	}

	var res Result
	if err := json.NewDecoder(rr.Body).Decode(&res); err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	if res.Sum != 8 {
		t.Errorf("Expected sum 8, got %d", res.Sum)
	}
}

func TestAddHandler_InvalidInput(t *testing.T) {
	form := url.Values{}
	form.Add("a", "abc")
	form.Add("b", "5")

	req := httptest.NewRequest(http.MethodPost, "/api/add", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("Expected status 400, got %d", rr.Code)
	}
}
