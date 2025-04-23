package alertmanager_test

import (
	"testing"

	"github.com/DucTran999/shared-pkg/alertmanager"
	"github.com/stretchr/testify/require"
)

func Test_SendAlert(t *testing.T) {
	am := alertmanager.NewAlertManager("http://localhost:9093")

	err := am.Send(
		alertmanager.WithLabels(alertmanager.Labels{
			"level": "critical",
		}),
		alertmanager.WithAnnotations(alertmanager.Annotations{
			"summary":     "TestAlert",
			"description": "Example alert",
		}),
	)

	require.NoError(t, err)
}
