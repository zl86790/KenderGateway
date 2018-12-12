package main

import (
	"fmt"
	"log"
	"net/http"

	"./core/services"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	log.Println(request.Body)
}

func main() {
	log.Println("starting...")
	results, err := services.ReadConfigs("./configures/endpoints")
	if err != nil {
		log.Println(err)
	}
	for e := results.Front(); e != nil; e = e.Next() {
		fmt.Println("Got", e.Value)
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
