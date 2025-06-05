package alertmanager

import "time"

type Labels map[string]any

type Annotations map[string]any

type options struct {
	Labels      Labels      `json:"labels"`
	Annotations Annotations `json:"annotations"`
	EndTime     time.Time   `json:"endsAt,omitempty"`
}

type Options func(*options)

func WithLabels(labels Labels) Options {
	return func(o *options) {
		o.Labels = labels
	}
}

func WithAnnotations(annotations Annotations) Options {
	return func(o *options) {
		o.Annotations = annotations
	}
}

func WithDuration(duration time.Duration) Options {
	return func(o *options) {
		o.EndTime = time.Now().Add(duration)
	}
}
