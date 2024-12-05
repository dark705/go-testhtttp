package slog

import (
	"context"
	"fmt"
	"log/slog"

	"go.opentelemetry.io/otel/trace"
)

const (
	otelFieldTraceID = "traceID"
	otelFieldSpanID  = "spanID"
)

type OTelHandler struct {
	slog.Handler
}

func (h *OTelHandler) Handle(ctx context.Context, record slog.Record) error {
	span := trace.SpanFromContext(ctx)
	if span.IsRecording() {
		record.AddAttrs(
			slog.String(otelFieldTraceID, span.SpanContext().TraceID().String()),
			slog.String(otelFieldSpanID, span.SpanContext().SpanID().String()))
	}

	return fmt.Errorf("slog, OTelHandler.Handle, err: %w", h.Handler.Handle(ctx, record))
}
