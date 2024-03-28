package env

import (
	"context"

	"github.com/codingconcepts/env"
)

type configKey struct{}

type Config struct {
	Port int `env:"REPLIK_PORT" default:"9090"`
}

func SetConfig() *Config {
	e := &Config{}
	if err := env.Set(e); err != nil {
		panic(err)
	}

	return e
}

func WithConfig(ctx context.Context) context.Context {
	return context.WithValue(ctx, configKey{}, SetConfig())
}

func GetConfig(ctx context.Context) *Config {
	env, ok := ctx.Value(configKey{}).(*Config)
	if !ok {
		return nil
	}
	return env
}
