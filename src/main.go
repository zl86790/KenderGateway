package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./core/services"
)

var globalRequestMap map[string]services.ServiceModel

func handler(writer http.ResponseWriter, request *http.Request) {
	log.Println(request.Body)
	log.Println(request.RequestURI)
	log.Println(globalRequestMap[request.RequestURI])

	client := &http.Client{}
	config := globalRequestMap[request.RequestURI]
	targetURL := config.TargetHost + config.TargetPath
	reqest, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		log.Println(err)
	}
	response, _ := client.Do(reqest)
	defer response.Body.Close()
	var by []byte
	by, _ = ioutil.ReadAll(response.Body)
	writer.Write(by)

}

func main() {
	log.Println("starting...")
	results, err := services.ReadConfigs("./configures/endpoints")
	if err != nil {
		log.Println(err)
	}
	for key, value := range results {
		fmt.Println("Key:", key, "Value:", value)
	}
	globalRequestMap = results
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
