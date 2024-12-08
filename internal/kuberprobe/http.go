package kuberprobe

import (
	"context"
	"crypto/rand"
	"math/big"
	"net/http"
	"time"
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
	HTTPRoutePattern = http.MethodGet + " /kuber/{" + probePlaceHolder + "}"

	probePlaceHolder = "probe"
	probeLive        = "live"
	probeReady       = "ready"
	probeStartUp     = "startup"

	maxRandomNumber = 99
)

type httpHandler struct {
	logger           Logger
	timeStartUp      time.Time
	probabilityLive  int
	probabilityReady int
}

func NewHTTPHandler(logger Logger,
	timeOutStartUpSeconds int,
	probabilityLive int,
	probabilityReady int) *httpHandler { //nolint:revive
	return &httpHandler{
		logger:           logger,
		timeStartUp:      time.Now().Add(time.Second * time.Duration(timeOutStartUpSeconds)),
		probabilityLive:  probabilityLive,
		probabilityReady: probabilityReady,
	}
}

func (h *httpHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	statusCode := http.StatusServiceUnavailable

	switch request.PathValue(probePlaceHolder) {
	case probeLive:
		if h.isLive() {
			statusCode = http.StatusOK
		}
	case probeReady:
		if h.isReady() {
			statusCode = http.StatusOK
		}
	case probeStartUp:
		if h.isStartUp() {
			statusCode = http.StatusOK
		}
	default:
		statusCode = http.StatusNotFound
	}

	responseWriter.WriteHeader(statusCode)
	_, _ = responseWriter.Write([]byte(http.StatusText(statusCode)))
}

func (h *httpHandler) isStartUp() bool {
	return time.Now().After(h.timeStartUp)
}

func (h *httpHandler) isReady() bool {
	return getRand() < h.probabilityReady
}

func (h *httpHandler) isLive() bool {
	return getRand() < h.probabilityLive
}

func getRand() int {
	num, _ := rand.Int(rand.Reader, big.NewInt(int64(maxRandomNumber)))

	return int(num.Int64())
}
