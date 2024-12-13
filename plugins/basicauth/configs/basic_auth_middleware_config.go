package configs

import (
	"github.com/mitchellh/mapstructure"
)

type BasicAuthMiddlewareConfig struct {
	Realm              string            `mapstructure:"realm"`
	Users              map[string]string `mapstructure:"users"`
	ContextUsernameKey string            `mapstructure:"contextUsernameKey"`
	ContextPasswordKey string            `mapstructure:"contextPasswordKey"`
}

func NewBasicAuthMiddlewareConfig(config interface{}) (*BasicAuthMiddlewareConfig, error) {
	var instance BasicAuthMiddlewareConfig

	err := mapstructure.Decode(config, &instance)
	if err != nil {
		return nil, err
	}

	if instance.Realm == "" {
		instance.Realm = "Restricted"
	}

	if instance.ContextUsernameKey == "" {
		instance.ContextUsernameKey = "basicAuthUsername"
	}

	if instance.ContextPasswordKey == "" {
		instance.ContextPasswordKey = "basicAuthPassword"
	}

	return &instance, nil
}
