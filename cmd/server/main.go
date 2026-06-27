// cmd/server/main.go
package main

import (
	"fmt"
	"log"
	"MME/internal/service"
)

func main() {
	fmt.Println("🚀 Starting Milestone 1: Local Metadata Extractor...")

	// Initialize our new service layer
	extractor := service.NewExtractorService()

	// A reliable sample image URL to test with
	sampleURL := "https://raw.githubusercontent.com/golang/go/master/doc/gopher/gophercolor.png"

	fmt.Printf("Analyzing image: %s\n", sampleURL)
	
	metadata, err := extractor.ExtractMetadata(sampleURL)
	if err != nil {
		log.Fatalf("❌ Error extracting metadata: %v", err)
	}

	// Print results out to your console
	fmt.Println("\n--- Extraction Successful ---")
	fmt.Printf("Format:       %s\n", metadata.Format)
	fmt.Printf("Width:        %d px\n", metadata.Width)
	fmt.Printf("Height:       %d px\n", metadata.Height)
	fmt.Printf("File Size:    %d bytes\n", metadata.SizeBytes)
	fmt.Println("-----------------------------")
}
