package configs

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

type LoggerMiddlewareConfig struct {
	Format        string        `mapstructure:"format"`
	TimeZone      string        `mapstructure:"timeZone"`
	TimeInterval  time.Duration `mapstructure:"timeInterval"`
	DisableColors bool          `mapstructure:"disableColors"`
}

func NewLoggerMiddlewareConfig(config interface{}) (*LoggerMiddlewareConfig, error) {
	var instance LoggerMiddlewareConfig

	err := mapstructure.Decode(config, &instance)
	if err != nil {
		return nil, err
	}

	if instance.Format == "" {
		instance.Format = "${pid} | ${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n"
	}

	if instance.TimeZone == "" {
		instance.TimeZone = "UTC"
	}

	if instance.TimeInterval == time.Duration(0) {
		instance.TimeInterval = 1 * time.Millisecond
	}

	return &instance, nil
}
