package configs

import (
	"github.com/mitchellh/mapstructure"
)

type RequestIdMiddlewareConfig struct {
	HeaderName string `mapstructure:"headerName"`
	ContextKey string `mapstructure:"contextKey"`
}

func NewRequestIdMiddlewareConfig(config interface{}) (*RequestIdMiddlewareConfig, error) {
	var instance RequestIdMiddlewareConfig

	err := mapstructure.Decode(config, &instance)
	if err != nil {
		return nil, err
	}

	if instance.HeaderName == "" {
		instance.HeaderName = "X-Request-ID"
	}

	if instance.ContextKey == "" {
		instance.ContextKey = "requestId"
	}

	return &instance, nil
}
