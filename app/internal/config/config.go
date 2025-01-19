package config

import (
	"github.com/caarlos0/env/v11"
	"log/slog"
	"sync"
)

type Config struct {
	Home string `env:"HOME"`
}

var (
	singleton sync.Once
	cfg       Config
)

func Get() Config {
	singleton.Do(func() {
		err := env.Parse(cfg)
		if err != nil {
			slog.Error("error parsing envs", "error", err)
		}
	})
	return cfg
}
