package linkdapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client is the main LinkdAPI client for interacting with the LinkdAPI service.
//
// This client provides:
//   - Automatic retry mechanism for failed requests
//   - Connection pooling for improved performance
//   - Comprehensive error handling
//   - Built-in timeout and cancellation support
//
// The Client is safe for concurrent use by multiple goroutines.
type Client struct {
	apiKey      string
	baseURL     string
	httpClient  *http.Client
	maxRetries  int
	retryDelay  time.Duration
	timeout     time.Duration
	ctx         context.Context // Context for all requests
}

// NewClient creates a new LinkdAPI client with default configuration.
//
// Example:
//
//	client := linkdapi.NewClient("your_api_key")
//	defer client.Close()
func NewClient(apiKey string) *Client {
	return NewClientWithConfig(apiKey, DefaultConfig())
}

// NewClientWithConfig creates a new LinkdAPI client with custom configuration.
//
// Example:
//
//	config := &linkdapi.Config{
//	    Timeout:    60 * time.Second,
//	    MaxRetries: 5,
//	}
//	client := linkdapi.NewClientWithConfig("your_api_key", config)
//	defer client.Close()
func NewClientWithConfig(apiKey string, config *Config) *Client {
	if config == nil {
		config = DefaultConfig()
	}

	// Use provided context or default to background context
	ctx := config.Context
	if ctx == nil {
		ctx = context.Background()
	}

	return &Client{
		apiKey:     apiKey,
		baseURL:    strings.TrimRight(config.BaseURL, "/"),
		maxRetries: config.MaxRetries,
		retryDelay: config.RetryDelay,
		timeout:    config.Timeout,
		ctx:        ctx,
		httpClient: &http.Client{
			Timeout: config.Timeout,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     90 * time.Second,
			},
		},
	}
}

// Close closes the HTTP client and releases resources.
func (c *Client) Close() {
	if c.httpClient != nil {
		c.httpClient.CloseIdleConnections()
	}
}

// getHeaders returns the default headers for API requests.
func (c *Client) getHeaders() map[string]string {
	return map[string]string{
		"X-linkdapi-apikey": c.apiKey,
		"Accept":            "application/json",
		"Content-Type":      "application/json",
		"User-Agent":        "LinkdAPI-Go-Client/1.0",
	}
}

// sendRequest sends an HTTP request with retry logic using the client's context.
func (c *Client) sendRequest(method, endpoint string, params map[string]string) (map[string]any, error) {
	ctx := c.ctx
	requestURL := fmt.Sprintf("%s/%s", c.baseURL, strings.TrimLeft(endpoint, "/"))

	// Add query parameters
	if len(params) > 0 {
		urlParams := url.Values{}
		for key, value := range params {
			urlParams.Add(key, value)
		}
		requestURL = fmt.Sprintf("%s?%s", requestURL, urlParams.Encode())
	}

	var lastErr error
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		// Check if context is already cancelled
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// Create request
		req, err := http.NewRequestWithContext(ctx, method, requestURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}

		// Add headers
		for key, value := range c.getHeaders() {
			req.Header.Set(key, value)
		}

		// Send request
		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = err
			if attempt < c.maxRetries {
				time.Sleep(c.retryDelay * time.Duration(attempt+1))
				continue
			}
			return nil, fmt.Errorf("request failed after %d attempts: %w", c.maxRetries+1, err)
		}

		// Read response body
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			lastErr = err
			if attempt < c.maxRetries {
				time.Sleep(c.retryDelay * time.Duration(attempt+1))
				continue
			}
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}

		// Check status code
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			lastErr = fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
			if attempt < c.maxRetries {
				time.Sleep(c.retryDelay * time.Duration(attempt+1))
				continue
			}
			return nil, lastErr
		}

		// Parse JSON response
		var result map[string]any
		if err := json.Unmarshal(body, &result); err != nil {
			return nil, fmt.Errorf("failed to parse JSON response: %w", err)
		}

		return result, nil
	}

	return nil, lastErr
}

// Helper functions for building parameter maps

// stringParam adds a string parameter to the params map if the value is not empty.
func stringParam(params map[string]string, key, value string) {
	if value != "" {
		params[key] = value
	}
}

// intParam adds an integer parameter to the params map.
func intParam(params map[string]string, key string, value int) {
	params[key] = fmt.Sprintf("%d", value)
}

// boolParam adds a boolean parameter to the params map if the pointer is not nil.
func boolParam(params map[string]string, key string, value *bool) {
	if value != nil {
		if *value {
			params[key] = "true"
		} else {
			params[key] = "false"
		}
	}
}

// sliceParam adds a slice parameter to the params map as comma-separated values.
func sliceParam(params map[string]string, key string, values []string) {
	if len(values) > 0 {
		params[key] = strings.Join(values, ",")
	}
}
