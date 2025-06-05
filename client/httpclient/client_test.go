package httpclient_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/DucTran999/shared-pkg/client/httpclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Get(t *testing.T) {
	// Arrange
	c := httpclient.NewClient()

	// Act
	resp, err := c.Get(context.Background(), "https://example.com")

	// Assert
	require.NoError(t, err, "Expected no error when do request to example.com")
	require.NotNil(t, resp, "Expected non-empty response from example.com")
	require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200 OK")
	require.Equal(t, "200 OK", resp.Status, "Expected status text to match 200 OK")
	assert.NotEmpty(t, resp.Body, "Expected non-empty response body from example.com")
}
