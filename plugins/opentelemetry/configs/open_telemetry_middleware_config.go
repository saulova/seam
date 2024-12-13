package configs

import (
	"github.com/mitchellh/mapstructure"
)

type OpenTelemetryMiddlewareConfig struct {
	TracerName string `mapstructure:"tracerName"`
}

func NewOpenTelemetryMiddlewareConfig(config interface{}) (*OpenTelemetryMiddlewareConfig, error) {
	var instance OpenTelemetryMiddlewareConfig

	err := mapstructure.Decode(config, &instance)
	if err != nil {
		return nil, err
	}

	if instance.TracerName == "" {
		instance.TracerName = "unknown"
	}

	return &instance, nil
}
