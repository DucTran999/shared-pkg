package alertmanager

type AlertManager interface {
	Alert() error
}
