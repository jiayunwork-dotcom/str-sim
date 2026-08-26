package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	mux := New(Config{Addr: ":8080"})
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestSimilarityEndpoint(t *testing.T) {
	mux := New(Config{Addr: ":8080"})
	payload := similarityRequest{A: "kitten", B: "sitting", Algo: "levenshtein"}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/similarity", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var resp similarityResponse
	json.Unmarshal(rec.Body.Bytes(), &resp)
	if resp.Score <= 0 || resp.Score >= 1 {
		t.Errorf("expected score between 0 and 1, got %f", resp.Score)
	}
}

func TestMatchEndpoint(t *testing.T) {
	mux := New(Config{Addr: ":8080"})
	payload := matchRequest{A: "hello", B: "hello", Algo: "levenshtein", Threshold: 0.9}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/match", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
	}
	var resp matchResponse
	json.Unmarshal(rec.Body.Bytes(), &resp)
	if !resp.Match {
		t.Error("expected match=true for identical strings")
	}
}

func TestAlgorithmsEndpoint(t *testing.T) {
	mux := New(Config{Addr: ":8080"})
	req := httptest.NewRequest(http.MethodGet, "/api/algorithms", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	var resp struct {
		Algorithms []string `json:"algorithms"`
	}
	json.Unmarshal(rec.Body.Bytes(), &resp)
	if len(resp.Algorithms) == 0 {
		t.Error("expected non-empty algorithms list")
	}
}

func TestMethodNotAllowed(t *testing.T) {
	mux := New(Config{Addr: ":8080"})
	endpoints := []string{"/api/similarity", "/api/match"}
	for _, ep := range endpoints {
		req := httptest.NewRequest(http.MethodGet, ep, nil)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("%s: expected 405, got %d", ep, rec.Code)
		}
	}
}
