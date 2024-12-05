package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dark705/go-testhtttp/internal/config"
	"github.com/dark705/go-testhtttp/internal/httpserver"
	"github.com/dark705/go-testhtttp/internal/httptesthandler"
	"github.com/dark705/go-testhtttp/internal/prometheus"
	"github.com/dark705/go-testhtttp/internal/slog"
	promhttpmetrics "github.com/slok/go-http-metrics/metrics/prometheus"
	promhttpmiddleware "github.com/slok/go-http-metrics/middleware"
	promhttpmiddlewarestd "github.com/slok/go-http-metrics/middleware/std"
)

func main() {
	envConfig := config.GetConfigFromEnv()

	logger := slog.New(slog.Config{Level: envConfig.LogLevel})
	logger.Infof("app, version: %s", envConfig.Version)

	prometheusServer := prometheus.NewServer(prometheus.Config{HTTPListenPort: envConfig.PrometheusPort}, logger)
	prometheusServer.Run()
	defer prometheusServer.Stop()

	httpTestHandler := httphandler.NewHTTPTestHandler(logger)
	httpHostHandler := httphandler.NewHTTPHostHandler(logger)

	httpHandler := http.NewServeMux()
	httpHandler.Handle(httphandler.HTTPTestRoutePattern, httpTestHandler)
	httpHandler.Handle(httphandler.HTTPHostRoutePattern, httpHostHandler)

	prometheusMiddlewareHandler := promhttpmiddleware.New(promhttpmiddleware.Config{
		Recorder: prometheus.NewFilterRecorder(
			promhttpmetrics.NewRecorder(promhttpmetrics.Config{}), []string{}),
	})

	httpHandlerWithMetric := promhttpmiddlewarestd.Handler("", prometheusMiddlewareHandler, httpHandler)

	httpServer := httpserver.NewServer(httpserver.Config{
		Name:                          "test",
		HTTPListenPort:                envConfig.HTTPPort,
		RequestHeaderMaxBytes:         envConfig.HTTPRequestHeaderMaxSize,
		ReadHeaderTimeoutMilliseconds: envConfig.HTTPRequestReadHeaderTimeoutMilliseconds,
	}, logger, httpHandlerWithMetric)

	httpServer.Run()
	defer httpServer.Stop()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	logger.Infof("got signal from OS: %v. shutdown...", <-osSignals)
}
