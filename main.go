package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HealthResponse struct {
	Nama      string `json:"nama"`
	NRP       string `json:"nrp"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Uptime    string `json:"uptime"`
}

var startTime time.Time

func healthHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startTime).String()

	response := HealthResponse{
		Nama:      "Adriel Mahira Dharma",
		NRP:       "5025241097",
		Status:    "UP",
		Timestamp: time.Now().Format(time.RFC3339),
		Uptime:    uptime,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	startTime = time.Now()

	http.HandleFunc("/health", healthHandler)

	port := ":3000"
	fmt.Printf("Server running on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
