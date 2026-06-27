// internal/service/handler.go
package service

import (
	"encoding/json"
	"net/http"
	"MME/internal/domain"
	"MME/internal/repository" // Import the new repository layer
)

type HttpHandler struct {
	extractor        *ExtractorService
	downstreamClient *repository.DownstreamClient
}

// Update constructor to take the downstream client
func NewHttpHandler(extractor *ExtractorService, dsClient *repository.DownstreamClient) *HttpHandler {
	return &HttpHandler{
		extractor:        extractor,
		downstreamClient: dsClient,
	}
}

func (h *HttpHandler) ExtractHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req domain.ExtractRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.ImageURL == "" {
		http.Error(w, "Invalid request body or missing image_url", http.StatusBadRequest)
		return
	}

	// Business logic execution
	metadata, err := h.extractor.ExtractMetadata(req.ImageURL)
	if err != nil {
		http.Error(w, "Failed to extract metadata: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// NEW: Forward the result downstream!
	if err := h.downstreamClient.ForwardMetadata(metadata); err != nil {
		// Log the error but don't fail the user request if it's non-blocking,
		// or handle it according to requirements. Here we'll notify the user.
		http.Error(w, "Metadata extracted but downstream forwarding failed: "+err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(metadata)
}