package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"str-sim/internal/sim"
)

type Config struct {
	Addr string
}

func New(cfg Config) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/api/similarity", handleSimilarity)
	mux.HandleFunc("/api/match", handleMatch)
	mux.HandleFunc("/api/algorithms", handleAlgorithms)
	return mux
}

func ListenAndServe(cfg Config) error {
	mux := New(cfg)
	return http.ListenAndServe(cfg.Addr, mux)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

type similarityRequest struct {
	A    string `json:"a"`
	B    string `json:"b"`
	Algo string `json:"algo"`
}

type similarityResponse struct {
	A     string  `json:"a"`
	B     string  `json:"b"`
	Algo  string  `json:"algo"`
	Score float64 `json:"score"`
}

func handleSimilarity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpError(w, http.StatusMethodNotAllowed, "POST required")
		return
	}
	var req similarityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	if req.A == "" && req.B == "" {
		httpError(w, http.StatusBadRequest, "a and b are required")
		return
	}
	if req.Algo == "" {
		req.Algo = "levenshtein"
	}
	score, err := sim.Similarity(req.A, req.B, req.Algo)
	if err != nil {
		httpError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, similarityResponse{A: req.A, B: req.B, Algo: req.Algo, Score: score})
}

type matchRequest struct {
	A         string  `json:"a"`
	B         string  `json:"b"`
	Algo      string  `json:"algo"`
	Threshold float64 `json:"threshold"`
}

type matchResponse struct {
	A         string  `json:"a"`
	B         string  `json:"b"`
	Algo      string  `json:"algo"`
	Score     float64 `json:"score"`
	Match     bool    `json:"match"`
	Threshold float64 `json:"threshold"`
}

func handleMatch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httpError(w, http.StatusMethodNotAllowed, "POST required")
		return
	}
	var req matchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	if req.A == "" && req.B == "" {
		httpError(w, http.StatusBadRequest, "a and b are required")
		return
	}
	if req.Algo == "" {
		req.Algo = "levenshtein"
	}
	if req.Threshold <= 0 {
		req.Threshold = 0.8
	}
	score, err := sim.Similarity(req.A, req.B, req.Algo)
	if err != nil {
		httpError(w, http.StatusBadRequest, err.Error())
		return
	}
	matched, _ := sim.Match(req.A, req.B, req.Algo, req.Threshold)
	writeJSON(w, http.StatusOK, matchResponse{
		A: req.A, B: req.B, Algo: req.Algo,
		Score: score, Match: matched, Threshold: req.Threshold,
	})
}

func handleAlgorithms(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]interface{}{"algorithms": sim.Algorithms()})
}

func httpError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(v)
}

func ParsePort(addr string) int {
	parts := strings.Split(addr, ":")
	if len(parts) < 2 {
		return 0
	}
	p, _ := strconv.Atoi(parts[len(parts)-1])
	return p
}

func FormatAddr(addr string) string {
	port := ParsePort(addr)
	if port == 0 {
		return addr
	}
	return fmt.Sprintf("http://0.0.0.0:%d", port)
}
