package slog

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

const (
	LevelDebug = slog.Level(-4)
	LevelInfo  = slog.Level(0)
	LevelWarn  = slog.Level(4)
	LevelError = slog.Level(8)
	LevelFatal = slog.Level(10)
)

var LevelNames = map[slog.Leveler]string{
	LevelDebug: "debug",
	LevelInfo:  "info",
	LevelWarn:  "warning",
	LevelError: "error",
	LevelFatal: "fatal",
}

type Config struct {
	Level string
}

type logger struct {
	slogLogger *slog.Logger
}

func New(c Config) *logger { //nolint:revive
	var level slog.Level

	err := level.UnmarshalText([]byte(c.Level))
	if err != nil {
		level = slog.LevelInfo
	}

	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(_ []string, attr slog.Attr) slog.Attr {
			if attr.Key == slog.LevelKey {
				lvl, _ := attr.Value.Any().(slog.Level)
				levelLabel, exists := LevelNames[lvl]
				if !exists {
					levelLabel = lvl.String()
				}
				attr.Value = slog.StringValue(levelLabel)
			}

			return attr
		},
	})

	slogLogger := slog.New(&OTelHandler{jsonHandler})

	return &logger{slogLogger}
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.slogLogger.Log(context.Background(), LevelDebug, fmt.Sprintf(format, args...))
}

func (l *logger) DebugfContext(ctx context.Context, format string, args ...interface{}) { //nolint:goprintffuncname
	l.slogLogger.Log(ctx, LevelDebug, fmt.Sprintf(format, args...))
}

func (l *logger) Infof(format string, args ...any) {
	l.slogLogger.Log(context.Background(), LevelInfo, fmt.Sprintf(format, args...))
}

func (l *logger) InfofContext(ctx context.Context, format string, args ...interface{}) { //nolint:goprintffuncname
	l.slogLogger.Log(ctx, LevelInfo, fmt.Sprintf(format, args...))
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.slogLogger.Log(context.Background(), LevelWarn, fmt.Sprintf(format, args...))
}

func (l *logger) WarnfContext(ctx context.Context, format string, args ...interface{}) { //nolint:goprintffuncname
	l.slogLogger.Log(ctx, LevelWarn, fmt.Sprintf(format, args...))
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.slogLogger.Log(context.Background(), LevelError, fmt.Sprintf(format, args...))
}

func (l *logger) ErrorfContext(ctx context.Context, format string, args ...interface{}) { //nolint:goprintffuncname
	l.slogLogger.Log(ctx, LevelError, fmt.Sprintf(format, args...))
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.slogLogger.Log(context.Background(), LevelFatal, fmt.Sprintf(format, args...))
	os.Exit(1)
}

func (l *logger) FatalfContext(ctx context.Context, format string, args ...interface{}) { //nolint:goprintffuncname
	l.slogLogger.Log(ctx, LevelFatal, fmt.Sprintf(format, args...))
	os.Exit(1)
}
