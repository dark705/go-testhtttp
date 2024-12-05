package prometheus

import (
	"context"
	"strings"
	"time"

	"github.com/slok/go-http-metrics/metrics"
)

func NewFilterRecorder(recorder metrics.Recorder, handlerIDPrefixFilter []string) *filterRecorder { //nolint:revive
	return &filterRecorder{recorder: recorder, handlerIDPrefixFilter: handlerIDPrefixFilter}
}

type filterRecorder struct {
	recorder              metrics.Recorder
	handlerIDPrefixFilter []string
}

func (r *filterRecorder) ObserveHTTPRequestDuration(ctx context.Context, props metrics.HTTPReqProperties, duration time.Duration) {
	props.ID = r.filterID(props.ID)
	r.recorder.ObserveHTTPRequestDuration(ctx, props, duration)
}

func (r *filterRecorder) ObserveHTTPResponseSize(ctx context.Context, props metrics.HTTPReqProperties, sizeBytes int64) {
	props.ID = r.filterID(props.ID)
	r.recorder.ObserveHTTPResponseSize(ctx, props, sizeBytes)
}

func (r *filterRecorder) AddInflightRequests(ctx context.Context, props metrics.HTTPProperties, quantity int) {
	props.ID = r.filterID(props.ID)
	r.recorder.AddInflightRequests(ctx, props, quantity)
}

func (r *filterRecorder) filterID(id string) string { //nolint:varnamelen
	for _, prefix := range r.handlerIDPrefixFilter {
		if strings.HasPrefix(id, prefix) {
			return prefix
		}
	}

	return id
}
