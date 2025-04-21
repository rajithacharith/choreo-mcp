package choreoclient

import (
	"net/http"
	"time"
)

// NewHTTPClient creates and returns a new HTTP client with default settings.
func NewHTTPClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}
