package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Consumer App Started ...")
	http.HandleFunc("/subscription", handleMessage)
	http.ListenAndServe(":8080", nil)
}

func handleMessage(writer http.ResponseWriter, request *http.Request) {

	log.Println("Event received ...")
	byteBody, _ := io.ReadAll(request.Body)
	log.Println(string(byteBody))
	time.Sleep(time.Second / 2)
}
