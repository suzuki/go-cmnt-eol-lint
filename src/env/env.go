package env

import (
	"context"

	"github.com/kelseyhightower/envconfig"
)

var (
	defaultPrefix = ""
	keyOfContext  = "config"
)

type Config struct {
	AllowedEOLs string `envconfig:"ALLOWED_EOLS" default:".!){}]"`
	MinWords    int    `envconfig:"MIN_WORDS" default:"5"`
}

func newConfig() *Config {
	var config Config

	if err := envconfig.Process(defaultPrefix, &config); err != nil {
		panic(err)
	}

	return &config
}

func WithConfig(ctx context.Context) context.Context {
	return context.WithValue(ctx, keyOfContext, newConfig())
}

func GetConfig(ctx context.Context) *Config {
	v := ctx.Value(keyOfContext)
	config := v.(*Config)

	return config
}
