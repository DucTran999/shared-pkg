package backoff

import (
	"math"
	"math/rand"
	"time"
)

type exponentialBackoff struct {
	Base   time.Duration // Initial delay
	Factor float64       // Growth factor
	Cap    time.Duration // Max delay
	Jitter bool          // Whether to add randomness
}

// NewExponentialBackoff creates a new exponential backoff strategy with optional configuration overrides.
//
// It uses a default base of 500ms, a cap of 5s, and jitter enabled.
// Options like WithBase, WithCap, WithJitter, and WithFactor can be used to customize behavior.
//
// Example:
//
//	b := NewExponentialBackoff(
//	    WithBase(1 * time.Second),
//	    WithFactor(2.0),
//	    WithCap(10 * time.Second),
//	    WithJitter(true),
//	)
func NewExponentialBackoff(opts ...BackoffOption) *exponentialBackoff {
	// Default config
	cfg := Config{
		Base:   500 * time.Millisecond,
		Factor: 2.0,
		Cap:    5 * time.Second,
		Jitter: true,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	return &exponentialBackoff{
		Base:   cfg.Base,
		Factor: cfg.Factor,
		Cap:    cfg.Cap,
		Jitter: cfg.Jitter,
	}
}

func (e *exponentialBackoff) Next(attempt int) time.Duration {
	delay := float64(e.Base) * math.Pow(e.Factor, float64(attempt))
	if e.Cap > 0 && time.Duration(delay) > e.Cap {
		delay = float64(e.Cap)
	}

	d := time.Duration(delay)
	if e.Jitter {
		half := int64(d / 2)
		// #nosec G404 -- math/rand is fine here for jitter in backoff
		d = time.Duration(half + rand.Int63n(half))
	}

	return d
}
