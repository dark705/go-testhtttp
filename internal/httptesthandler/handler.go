package httphandler

import (
	"context"
	"net/http"
	"os"
)

type Logger interface {
	Debugf(format string, args ...any)
	DebugfContext(ctx context.Context, format string, args ...any)

	Infof(format string, args ...any)
	InfofContext(ctx context.Context, format string, args ...any)

	Warnf(format string, args ...any)
	WarnfContext(ctx context.Context, format string, args ...any)

	Errorf(format string, args ...any)
	ErrorfContext(ctx context.Context, format string, args ...any)

	Fatalf(format string, args ...any)
	FatalfContext(ctx context.Context, format string, args ...any)
}

const (
	HTTPTestRoutePattern = http.MethodGet + " /test"
	HTTPHostRoutePattern = http.MethodGet + " /host"
)

type httpTestHandler struct {
	logger Logger
}

func NewHTTPTestHandler(logger Logger) *httpTestHandler { //nolint:revive
	return &httpTestHandler{
		logger: logger,
	}
}

func (h *httpTestHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	responseWriter.WriteHeader(http.StatusOK)
	_, err := responseWriter.Write([]byte("Ok"))
	if err != nil {
		h.logger.ErrorfContext(ctx, "httphandler, httpTestHandler.ServeHTTP, responseWriter.Write, error: %s", err)
	}
}

type httpHostHandler struct {
	logger Logger
}

func NewHTTPHostHandler(logger Logger) *httpHostHandler { //nolint:revive
	return &httpHostHandler{
		logger: logger,
	}
}

func (h *httpHostHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	hostname, err := os.Hostname()
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		h.logger.ErrorfContext(ctx, "httphandler, httpHostHandler.ServeHTTP, os.Hostname, error: %s", err)

		return
	}

	responseWriter.WriteHeader(http.StatusOK)

	_, err = responseWriter.Write([]byte(hostname))
	if err != nil {
		h.logger.ErrorfContext(ctx, "httphandler, httpTestHandler.ServeHTTP, responseWriter.Write, error: %s", err)
	}
}
