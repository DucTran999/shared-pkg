package backoff

import (
	"time"
)

// ConstantBackoff is a backoff strategy where each retry interval is the same.
//
// It's useful when you want predictable, fixed-delay retries.
type constantBackoff struct {
	Interval time.Duration
}

// NewConstantBackoff creates a new constant backoff strategy with optional configuration.
//
// It uses a default interval of 500ms unless overridden with WithBase().
// Jitter and other exponential/linear options are ignored.
//
// Example:
//
//	b := NewConstantBackoff(
//	    WithBase(1 * time.Second),
//	)
func NewConstantBackoff(opts ...BackoffOption) *constantBackoff {
	// Default config
	cfg := Config{
		Base: 500 * time.Millisecond,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	return &constantBackoff{
		Interval: cfg.Base,
	}
}

// Next returns the constant interval regardless of the attempt number.
// The attempt number is ignored in this strategy.
func (c *constantBackoff) Next(_ int) time.Duration {
	return c.Interval
}
