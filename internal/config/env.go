package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/dark705/go-testhtttp/internal/helpers"
)

type EnvConfig struct {
	Version  string `env:"VERSION" envDefault:"version_not_set"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`

	HTTPPort                                 string `env:"HTTP_PORT" envDefault:"8000"`
	HTTPRequestHeaderMaxSize                 int    `env:"HTTP_REQUEST_HEADER_MAX_SIZE" envDefault:"10000"`
	HTTPRequestReadHeaderTimeoutMilliseconds int    `env:"HTTP_REQUEST_READ_HEADER_TIMEOUT_MILLISECONDS" envDefault:"2000"`

	PrometheusPort string `env:"PROMETHEUS_PORT" envDefault:"9000"`
}

func GetConfigFromEnv() *EnvConfig {
	config := &EnvConfig{}
	err := env.Parse(config)
	helpers.FailOnError(err, "fail get config from Env")

	return config
}
