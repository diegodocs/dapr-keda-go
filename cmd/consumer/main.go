package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

const (
	maxRequestSize = 10 << 20 // 10MB
)

func main() {
	log.Println("Consumer App Started ...")
	mux := http.NewServeMux()
	mux.HandleFunc("/subscription", handleMessage)

	// Configure HTTP server with security timeouts
	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1MB
	}

	log.Println("Server listening on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("Server failed to start:", err)
	}
}

func handleMessage(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Println("Event received ...")

	// Limit request body size to prevent memory exhaustion
	request.Body = http.MaxBytesReader(writer, request.Body, maxRequestSize)
	defer request.Body.Close()

	byteBody, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(writer, "Failed to read request body", http.StatusBadRequest)
		return
	}

	log.Println(string(byteBody))
	time.Sleep(time.Second / 2)

	writer.WriteHeader(http.StatusOK)
}
