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
	Send() error
}

type alertManager struct {
	alertEndpoint string
	httpClient    http.Client
}

func NewAlertManager(host string) *alertManager {
	once.Do(func() {
		amInst = &alertManager{
			alertEndpoint: fmt.Sprintf("%s/api/v2/alerts", host),
		}
	})

	return amInst
}

func (am *alertManager) Send(opts ...Options) error {
	var cfg options

	for _, opt := range opts {
		opt(&cfg)
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
	// Print JSON you're sending
	fmt.Println("Request JSON:", string(jsonByte))

	req, postErr := http.NewRequest(http.MethodPost, am.alertEndpoint, bytes.NewBuffer(jsonByte))
	if postErr != nil {
		return postErr
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Output response
	fmt.Println("Status:", resp.Status)
	return nil
}
