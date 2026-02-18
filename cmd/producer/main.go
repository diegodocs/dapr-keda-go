package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

const (
	pubsubName       = "plant-tree-pubsub"
	pubsubTopic      = "events"
	maxRequestSize   = 1 << 20 // 1MB
	maxTreesPerBatch = 10000
	requestTimeout   = 30 * time.Second
	publishTimeout   = 5 * time.Second
)

type Handler struct {
	Client          dapr.Client
	EventsProcessed int
}

type PlantCommand struct {
	NumberOfTrees int `json:"numberOfTrees"`
}

type PlantedTreeEvent struct {
	Id int `json:"id"`
}

func main() {
	log.Println("Producer App Started ...")
	// Create a new client for DAPR using the SDK
	client, err := dapr.NewClient()
	if err != nil {
		log.Fatalln("Error to create instance of DAPR Client: ", err)
	}
	defer client.Close()

	handler := &Handler{Client: client, EventsProcessed: 1}
	mux := http.NewServeMux()
	mux.Handle("/plant", handler)

	// Configure HTTP server with security timeouts
	server := &http.Server{
		Addr:              ":8081",
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1MB
	}

	log.Println("Server listening on :8081")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("Server failed to start:", err)
	}
}

func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Limit request body size to prevent memory exhaustion
	request.Body = http.MaxBytesReader(writer, request.Body, maxRequestSize)
	defer request.Body.Close()

	var command PlantCommand
	if err := json.NewDecoder(request.Body).Decode(&command); err != nil {
		log.Printf("Error decoding request: %v", err)
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate input
	if command.NumberOfTrees <= 0 {
		http.Error(writer, "numberOfTrees must be positive", http.StatusBadRequest)
		return
	}
	if command.NumberOfTrees > maxTreesPerBatch {
		http.Error(writer, "numberOfTrees exceeds maximum allowed", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(request.Context(), requestTimeout)
	defer cancel()

	if err := plantTree(ctx, command.NumberOfTrees, handler); err != nil {
		log.Printf("Error planting trees: %v", err)
		http.Error(writer, "Failed to process request", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	if _, err := writer.Write([]byte("Trees planted successfully")); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func plantTree(ctx context.Context, numberOfTrees int, handler *Handler) error {
	// Publish events using DAPR pubsub
	for i := 1; i <= numberOfTrees; i++ {
		// Check if context has been cancelled
		select {
		case <-ctx.Done():
			return errors.New("request cancelled or timed out")
		default:
		}

		event := PlantedTreeEvent{Id: handler.EventsProcessed}
		jsonOut, err := json.Marshal(event)
		if err != nil {
			log.Printf("Error marshaling event: %v", err)
			return err
		}

		// Use context with timeout for publishing
		publishCtx, cancel := context.WithTimeout(ctx, publishTimeout)
		err = handler.Client.PublishEvent(
			publishCtx,
			pubsubName,
			pubsubTopic,
			jsonOut)
		cancel()

		if err != nil {
			log.Printf("Error publishing event: %v", err)
			return err
		}

		log.Println("Event Published - Id:", event.Id)
		handler.EventsProcessed++
	}
	return nil
}
