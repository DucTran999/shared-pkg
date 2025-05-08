package http_test

import (
	"testing"

	"github.com/DucTran999/shared-pkg/client/http"
	"github.com/stretchr/testify/require"
)

func Test_Get(t *testing.T) {
	httpClient := http.NewClient()

	_, err := httpClient.Get("https://jsonplaceholder.typicode.com/posts/1")

	require.Nil(t, err)
}
