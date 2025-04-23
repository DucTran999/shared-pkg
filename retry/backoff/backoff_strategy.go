// Package backoff provides configurable backoff strategies for retry logic.
package backoff

import "time"

// BackoffStrategy defines the interface for backoff strategies.
// It calculates the next delay duration based on the retry attempt number.
type BackoffStrategy interface {
	Next(attempt int) time.Duration
}

// BackoffOption represents a functional option for configuring a backoff strategy.
type BackoffOption func(*Config)

// Config holds shared configuration options used by different backoff strategies.
type Config struct {
	Base   time.Duration // Base delay (used in exponential backoff)
	Step   time.Duration // Step increment (used in linear backoff)
	Factor float64       // Multiplication factor (used in exponential backoff)
	Cap    time.Duration // Maximum cap for delay
	Jitter bool          // If true, apply jitter (randomization) to delay
}

// WithBase sets the base delay duration.
// Commonly used in exponential and constant backoff strategies.
func WithBase(d time.Duration) BackoffOption {
	return func(c *Config) {
		c.Base = d
	}
}

// WithJitter enables or disables jitter.
// Jitter randomizes delay to avoid retry storms or contention.
func WithJitter(enabled bool) BackoffOption {
	return func(c *Config) {
		c.Jitter = enabled
	}
}

// WithStep sets the step increment duration.
// Used in linear backoff strategies.
func WithStep(d time.Duration) BackoffOption {
	return func(c *Config) {
		c.Step = d
	}
}

// WithFactor sets the exponential growth factor.
// Used in exponential backoff strategies.
func WithFactor(f float64) BackoffOption {
	return func(c *Config) {
		c.Factor = f
	}
}

// WithCap sets the maximum cap for delay.
// Useful to prevent retry delays from growing indefinitely.
func WithCap(cap time.Duration) BackoffOption {
	return func(c *Config) {
		c.Cap = cap
	}
}
