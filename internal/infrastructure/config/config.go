package config

import (
	"os"
	"strconv"
	"time"
)

const defaultHTTPTimeout = 15000 * time.Millisecond

type Config struct {
	Logger loggerConfig
	HTTP   httpConfig
}

type loggerConfig struct {
	Level string
}

type httpConfig struct {
	Port    string
	Timeout time.Duration
}

func NewConfig() *Config {
	var cfg = new(Config)

	cfg.Logger.Level = os.Getenv("LOG_LEVEL")
	cfg.HTTP.Port = os.Getenv("HTTP_PORT")
	cfg.HTTP.Timeout = defaultHTTPTimeout

	httpTimeoutStr := os.Getenv("HTTP_TIMEOUT")
	if httpTimeoutStr != "" {
		httpTimeout, err := strconv.Atoi(httpTimeoutStr)

		if err == nil {
			cfg.HTTP.Timeout = time.Duration(httpTimeout) * time.Millisecond
		}
	}

	return cfg
}
