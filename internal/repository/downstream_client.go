// internal/repository/downstream_client.go
package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"MME/internal/domain"
)

type DownstreamClient struct {
	targetURL string
	client    *http.Client
}

func NewDownstreamClient(targetURL string) *DownstreamClient {
	return &DownstreamClient{
		targetURL: targetURL,
		client:    &http.Client{Timeout: 5 * time.Second}, // Don't let downstream slow down your app
	}
}

// ForwardMetadata sends the finalized metadata payload downstream
func (c *DownstreamClient) ForwardMetadata(metadata *domain.ImageMetadata) error {
	// 1. Serialize the domain data back into JSON
	payload, err := json.Marshal(metadata)
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %w", err)
	}

	// 2. Dispatch the HTTP POST request to the downstream URL
	resp, err := c.client.Post(c.targetURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("downstream service unreachable: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("downstream service returned error status: %d", resp.StatusCode)
	}

	return nil
}