package httpclient

import "context"

// HttpClient defines a simple interface for performing HTTP requests with context support.
//
// Example usage:
//
//	var client HttpClient = NewClient()
//	data, err := client.Get(ctx, "http://example.com")
type HttpClient interface {
	// Get performs an HTTP GET request to the specified URL using the provided context.
	Get(ctx context.Context, url string) (*response, error)
}
