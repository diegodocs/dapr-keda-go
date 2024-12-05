package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
)

const (
	pubsubName  = "plant-tree-pubsub"
	pubsubTopic = "events"
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
	http.Handle("/plant", handler)
	http.ListenAndServe(":8081", nil)
}

func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	var command PlantCommand
	_ = json.NewDecoder(request.Body).Decode(&command)
	plantTree(command.NumberOfTrees, handler)

}

func plantTree(numberOfTrees int, handler *Handler) {

	// Publish events using DAPR pubsub
	for i := 1; i <= numberOfTrees; i++ {

		event := PlantedTreeEvent{Id: handler.EventsProcessed}
		jsonOut, _ := json.Marshal(event)

		err := handler.Client.PublishEvent(
			context.Background(),
			pubsubName,
			pubsubTopic,
			[]byte(jsonOut))

		if err != nil {
			log.Println("Error: ", err)
		}

		log.Println("Event Published - Id:", event.Id)
		handler.EventsProcessed++
	}
}
