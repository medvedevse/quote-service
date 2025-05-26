package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"quote-service/internal/handler"
)

func TestGetQuotes(t *testing.T) {
	req, err := http.NewRequest("GET", "/quotes", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.GetQuotes)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %d", w.Code)
	}
}
