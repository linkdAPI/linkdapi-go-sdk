package linkdapi

import (
	"context"
	"time"
)

// Config holds configuration options for the LinkdAPI client.
type Config struct {
	// BaseURL is the base URL for the API (default: "https://linkdapi.com")
	BaseURL string

	// Timeout is the request timeout (default: 30 seconds)
	Timeout time.Duration

	// MaxRetries is the maximum number of retry attempts (default: 3)
	MaxRetries int

	// RetryDelay is the initial delay between retries (default: 1 second)
	// Note: Delay increases exponentially with each retry
	RetryDelay time.Duration

	// Context is the context to use for all requests (default: context.Background())
	// Set this if you need custom timeout or cancellation behavior
	Context context.Context
}

// DefaultConfig returns a Config with default values.
func DefaultConfig() *Config {
	return &Config{
		BaseURL:    "https://linkdapi.com",
		Timeout:    30 * time.Second,
		MaxRetries: 3,
		RetryDelay: 1 * time.Second,
	}
}
