package alertmanager

import (
	"strings"
)

type alertManager struct {
	host string

	alertName   string
	level       string
	summary     string
	description string
}

func NewAlertManager(host string) (*alertManager, error) {
	hostCleaned := strings.Trim(host, " ")
	if len(host) == 0 {
		return nil, ErrEmptyHost
	}

	return &alertManager{
		host: hostCleaned,
	}, nil
}

func (am *alertManager) Alert() *alertManager {
	return am
}
