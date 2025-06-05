package alertmanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	amInst *alertManager
	once   sync.Once
)

type AlertManager interface {
	Send(opts ...Options) error
}

type alertManager struct {
	alertEndpoint string
	httpClient    http.Client
}

func NewAlertManager(host string) *alertManager {
	if host == "" {
		panic(ErrEmptyHost)
	}

	once.Do(func() {
		amInst = &alertManager{
			alertEndpoint: fmt.Sprintf("%s/api/v2/alerts", host),
			httpClient: http.Client{
				Timeout: time.Second * 10, // Set a timeout for the HTTP client
			},
		}
	})

	return amInst
}

func (am *alertManager) Send(opts ...Options) error {
	var cfg options

	for _, opt := range opts {
		opt(&cfg)
	}

	// Validate required fields
	if len(cfg.Labels) == 0 {
		return fmt.Errorf("alert must have at least one label")
	}

	// Ensure required labels like alertname are present
	if _, ok := cfg.Labels["alertname"]; !ok {
		return fmt.Errorf("alertname label is required")
	}

	if cfg.EndTime.IsZero() {
		cfg.EndTime = time.Now().Add(time.Second * 30)
	}

	arr := []options{cfg}

	return am.sendHttpRequest(arr)
}

func (am *alertManager) sendHttpRequest(optsPost []options) error {
	jsonByte, err := json.Marshal(optsPost)
	if err != nil {
		return err
	}

	req, postErr := http.NewRequest(http.MethodPost, am.alertEndpoint, bytes.NewBuffer(jsonByte))
	if postErr != nil {
		return postErr
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send alert: %w", err)
	}
	defer resp.Body.Close()

	// Check for non-2xx response
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("alertmanager returned non-2xx status: %s", resp.Status)
	}

	return nil
}
