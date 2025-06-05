package httpclient

import (
	"net/http"
	"time"
)

// Option is a functional option type for configuring the httpClient.
type Option func(*httpClient)

// WithTimeout returns an Option that sets the timeout duration for HTTP requests made by the client.
// This controls the maximum amount of time a request can take before being canceled.
//   - Defaults to 5 seconds if not specified or if an invalid value is provided.
//
// Example usage:
//
//	client := NewClient(WithTimeout(10 * time.Second))
func WithTimeout(timeout time.Duration) Option {
	return func(c *httpClient) {
		// If timeout is invalid ignore it and use the default timeout.
		if timeout <= 0 {
			return
		}
		c.client.Timeout = timeout
	}
}

// WithMaxIdleConns sets the MaxIdleConns on the underlying *http.Transport.
// This controls the maximum number of idle connections across all hosts.
//   - If n is invalid (<= 0), it will ignore it and use the default value of 100.
//
// Example usage:
//
//	client := NewClient(WithMaxIdleConns(50))
func WithMaxIdleConns(n int) Option {
	return func(c *httpClient) {
		// If n is invalid (<= 0), ignore it and use the default value. (100)
		if n <= 0 {
			return
		}

		if c.client.Transport != nil {
			if transport, ok := c.client.Transport.(*http.Transport); ok {
				transport.MaxIdleConns = n
			}
		}
	}
}

// WithIdleConnTimeout sets the maximum amount of time an idle (keep-alive) connection
// will remain idle before closing. This option modifies the IdleConnTimeout field
// of the underlying http.Transport.
//
// Example usage:
//
//	client := NewClient(WithIdleConnTimeout(90 * time.Second))
func WithIdleConnTimeout(d time.Duration) Option {
	return func(c *httpClient) {
		// If d is invalid (<= 0), ignore it and use the default value. (90 seconds)
		if d <= 0 {
			return
		}

		if c.client.Transport != nil {
			if transport, ok := c.client.Transport.(*http.Transport); ok {
				transport.IdleConnTimeout = d
			}
		}
	}
}

// WithMaxConnsPerHost sets the maximum number of concurrent connections per host
// for the underlying http.Transport. This helps limit or control parallelism
// when communicating with a specific service.
//
// If n <= 0, the option is ignored and the default is used (commonly 10).
//
// Example usage:
//
//	client := NewClient(WithMaxConnsPerHost(20))
func WithMaxConnsPerHost(n int) Option {
	return func(c *httpClient) {
		// If n is invalid (<= 0), ignore it and use the default value. (10)
		if n <= 0 {
			return
		}

		if c.client.Transport != nil {
			if transport, ok := c.client.Transport.(*http.Transport); ok {
				transport.MaxConnsPerHost = n
			}
		}
	}
}

// WithTLSHandshakeTimeout sets the maximum amount of time the client will wait
// for a TLS handshake to complete with the server. This modifies the
// TLSHandshakeTimeout field of the underlying http.Transport.
//
//   - If d <= 0, the value is ignored and the default is used (typically 5 seconds).
//
// Example usage:
//
//	client := NewClient(WithTLSHandshakeTimeout(10 * time.Second))
func WithTLSHandshakeTimeout(d time.Duration) Option {
	return func(c *httpClient) {
		// If d is invalid (<= 0), ignore it and use the default value. (5 seconds)
		if d <= 0 {
			return
		}

		if c.client.Transport != nil {
			if transport, ok := c.client.Transport.(*http.Transport); ok {
				transport.TLSHandshakeTimeout = d
			}
		}
	}
}
