package main

import (
	"github.com/dark705/go-testhtttp/internal/config"
	"github.com/dark705/go-testhtttp/internal/httpserver"
	"github.com/dark705/go-testhtttp/internal/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	envConfig := config.GetConfigFromEnv()

	logger := slog.New(slog.Config{Level: envConfig.LogLevel})
	logger.Infof("app, version: %s", envConfig.Version)

	httpHandler := http.NewServeMux()

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

	logger.Infof("got signal from OS: %v. Shutdown...", <-osSignals)
}
