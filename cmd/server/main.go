// cmd/server/main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"MME/internal/repository"
	"MME/internal/service"
)

func main() {
	port := ":8080"
	
	// For testing, change this to your actual downstream API endpoint (e.g., a mock receiver or a database hook)
	downstreamURL := "https://httpbin.org/post" 

	fmt.Printf("🚀 Microservice booting up. Forwarding data to: %s\n", downstreamURL)

	// Dependency Injection
	extractor := service.NewExtractorService()
	dsClient := repository.NewDownstreamClient(downstreamURL)
	handler := service.NewHttpHandler(extractor, dsClient)

	http.HandleFunc("/extract", handler.ExtractHandler)

	log.Fatal(http.ListenAndServe(port, nil))
}