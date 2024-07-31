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

type CarbonFree struct {
	NumberOfTrees int `json:"numberOfTrees"`
}

type PlantedTreeEvent struct {
	Id int `json:"id"`
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
		//time.Sleep(time.Second / 2)
		handler.EventsProcessed++
	}
}

func main() {
	log.Println("Producer App Started ...")
	// Create a new client for DAPR using the SDK
	client, err := dapr.NewClient()
	if err != nil {
		log.Fatalln("Error to create instace of DAPR Client: ", err)
	}
	defer client.Close()

	handler := &Handler{Client: client, EventsProcessed: 1}
	http.Handle("/plant", handler)
	http.ListenAndServe(":8081", nil)
}

func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	var carbon CarbonFree
	_ = json.NewDecoder(request.Body).Decode(&carbon)
	plantTree(carbon.NumberOfTrees, handler)

}
