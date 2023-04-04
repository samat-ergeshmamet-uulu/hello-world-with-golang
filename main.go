package main

import (
	"log"
	"net/http"
)

func HelloWorldWebServer(writer http.ResponseWriter, request *http.Request){
	writer.Write([] byte("Hello, World!!!"))
}

func main() {
	http.HandleFunc("/", HelloWorldWebServer)
	log.Println("Start HTTP server on port 8080");
	log.Fatal(http.ListenAndServe(":8080", nil))
}