package httpserver

import (
	"context"
	"errors"
	"github.com/dark705/go-testhtttp/internal/helpers"
	"net"
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

type Server struct {
	httpServer *http.Server
	logger     Logger
	config     Config
}

type Config struct {
	Name                          string
	HTTPListenInterface           string
	HTTPListenPort                string
	RequestHeaderMaxBytes         int
	ReadHeaderTimeoutMilliseconds int
}

func NewServer(config Config, logger Logger, handler http.Handler) *Server {
	return &Server{
		logger:     logger,
		config:     config,
		httpServer: &http.Server{Handler: handler, MaxHeaderBytes: config.RequestHeaderMaxBytes, ReadHeaderTimeout: time.Duration(config.ReadHeaderTimeoutMilliseconds) * time.Millisecond},
	}
}

func (s *Server) Run() {
	address := s.config.HTTPListenInterface + ":" + s.config.HTTPListenPort
	s.logger.Infof("%s HTTPServer, start on: %s", s.config.Name, address)
	listener, err := net.Listen("tcp", address)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		helpers.FailOnError(err, s.config.Name+" HTTPServer, fail open port")
	}
	go func() {
		err = s.httpServer.Serve(listener)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			helpers.FailOnError(err, s.config.Name+" HTTPServer, fail start")
		}
	}()
}

func (s *Server) Stop() {
	s.logger.Infof(s.config.Name + " HTTPServer, stop...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		s.logger.Errorf(s.config.Name + " HTTPServer, fail stop")
		cancel()

		return
	}
	s.logger.Infof(s.config.Name + "HTTPServer, success stop")
	cancel()
}
