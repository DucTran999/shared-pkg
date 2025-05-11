package logger

// Define a private type to avoid collisions
type contextKey string

const (
	// Env
	Development = "development"
	Testing     = "testing"
	Staging     = "staging"
	Production  = "production"

	// key extract from context
	RequestIDKeyCtx = contextKey("request-id")
)
