package backoff

import (
	"math/rand"
	"time"
)

// LinearBackoff implements a linear backoff strategy where the delay increases
// by a fixed step per retry attempt, with optional jitter and maximum cap.
type linearBackoff struct {
	Step   time.Duration // Step is the fixed duration added per attempt (e.g., 2s, 5s)
	Cap    time.Duration // Cap limits the maximum delay to prevent excessive waiting
	Jitter bool          // Jitter, if true, adds randomness to the delay to avoid retry storms
}

// NewLinearBackoff creates a new linear backoff strategy with optional configuration overrides.
//
// It starts with a default step of 500ms, a cap of 5s, and jitter is enabled.
// Options like WithStep, WithCap, and WithJitter can be passed to customize the behavior.
//
// Example:
//
//	b := NewLinearBackoff(
//	    WithStep(1 * time.Second),
//	    WithCap(10 * time.Second),
//	    WithJitter(false),
//	)
func NewLinearBackoff(opts ...BackoffOption) *linearBackoff {
	// Default config
	cfg := Config{
		Step:   500 * time.Millisecond,
		Cap:    5 * time.Second,
		Jitter: true,
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	return &linearBackoff{
		Step:   cfg.Step,
		Cap:    cfg.Cap,
		Jitter: cfg.Jitter,
	}
}

// Next calculates the delay duration for a given retry attempt using linear backoff logic.
func (l *linearBackoff) Next(attempt int) time.Duration {
	base := time.Duration(attempt) * l.Step
	if l.Cap > 0 && base > l.Cap {
		base = l.Cap
	}

	if l.Jitter {
		half := int64(base / 2)
		// #nosec G404 -- math/rand is fine here for jitter in backoff
		return time.Duration(half + rand.Int63n(half))
	}

	return base
}
