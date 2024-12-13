package configs

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

type LimiterMiddlewareConfig struct {
	Storage                string        `mapstructure:"storage"`
	MaxConnections         int           `mapstructure:"maxConnections"`
	Expiration             time.Duration `mapstructure:"expiration"`
	SkipFailedRequests     bool          `mapstructure:"skipFailedRequests"`
	SkipSuccessfulRequests bool          `mapstructure:"skipSuccessfulRequests"`
	SlidingWindow          bool          `mapstructure:"slidingWindow"`
}

func NewLimiterMiddlewareConfig(config interface{}) (*LimiterMiddlewareConfig, error) {
	var instance LimiterMiddlewareConfig

	err := mapstructure.Decode(config, &instance)
	if err != nil {
		return nil, err
	}

	if instance.MaxConnections == 0 {
		instance.MaxConnections = 5
	}

	if instance.Expiration == time.Duration(0) {
		instance.Expiration = 1 * time.Minute
	}

	return &instance, nil
}
