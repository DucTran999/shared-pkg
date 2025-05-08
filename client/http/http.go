package http

import (
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

type HttpClient interface {
	Get(url string) ([]byte, error)
}

type httpClient struct {
	goHttp *http.Client
}

func NewClient() *httpClient {
	return &httpClient{
		goHttp: http.DefaultClient,
	}
}

func (h *httpClient) Get(url string) ([]byte, error) {
	resp, err := h.goHttp.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http get failed: %w", err)
	}

	defer func() {
		if cErr := resp.Body.Close(); cErr != nil {
			log.Warn().Err(cErr).Str("url", url).Msg("failed to close response body")
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}
