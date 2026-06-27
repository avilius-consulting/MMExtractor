// cmd/server/main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"MME/internal/service"
)

func main() {
	port := ":8080"
	fmt.Printf("🚀 Microservice booting up on port %s...\n", port)

	// Initialize dependencies
	extractor := service.NewExtractorService()
	handler := service.NewHttpHandler(extractor)

	// Define routes
	http.HandleFunc("/extract", handler.ExtractHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(port, nil))
}