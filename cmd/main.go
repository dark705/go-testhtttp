package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dark705/go-testhtttp/internal/config"
	"github.com/dark705/go-testhtttp/internal/httpserver"
	"github.com/dark705/go-testhtttp/internal/httptesthandler"
	"github.com/dark705/go-testhtttp/internal/slog"
)

func main() {
	envConfig := config.GetConfigFromEnv()

	logger := slog.New(slog.Config{Level: envConfig.LogLevel})
	logger.Infof("app, version: %s", envConfig.Version)

	httpTestHandler := httphandler.NewHTTPTestHandler(logger)
	httpHostHandler := httphandler.NewHTTPHostHandler(logger)

	httpHandler := http.NewServeMux()
	httpHandler.Handle(httphandler.HTTPTestRoutePattern, httpTestHandler)
	httpHandler.Handle(httphandler.HTTPHostRoutePattern, httpHostHandler)

	httpServer := httpserver.NewServer(httpserver.Config{
		Name:                          "test",
		HTTPListenPort:                envConfig.HTTPPort,
		RequestHeaderMaxBytes:         envConfig.HTTPRequestHeaderMaxSize,
		ReadHeaderTimeoutMilliseconds: envConfig.HTTPRequestReadHeaderTimeoutMilliseconds,
	}, logger, httpHandler)

	httpServer.Run()
	defer httpServer.Stop()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	logger.Infof("got signal from OS: %v. shutdown...", <-osSignals)
}
