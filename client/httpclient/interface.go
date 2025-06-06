package httpclient

import "context"

// HTTPClient defines a simple interface for performing HTTP requests with context support.
//
// Example usage:
//
//	var client HTTPClient = NewClient()
//	data, err := client.Get(ctx, "http://example.com")
type HTTPClient interface {
	// Get performs an HTTP GET request to the specified URL using the provided context.
	Get(ctx context.Context, url string) (*response, error)
}
