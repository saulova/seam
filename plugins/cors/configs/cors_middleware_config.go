package configs

import (
	"github.com/mitchellh/mapstructure"
)

type CORSMiddlewareConfig struct {
	AllowOrigins     string `mapstructure:"allowOrigins"`
	AllowMethods     string `mapstructure:"allowMethods"`
	AllowHeaders     string `mapstructure:"allowHeaders"`
	AllowCredentials bool   `mapstructure:"allowCredentials"`
	ExposeHeaders    string `mapstructure:"exposeHeaders"`
	MaxAge           int    `mapstructure:"maxAge"`
}

func NewCORSMiddlewareConfig(config interface{}) (*CORSMiddlewareConfig, error) {
	var instance CORSMiddlewareConfig

	err := mapstructure.Decode(config, &instance)
	if err != nil {
		return nil, err
	}

	if instance.AllowOrigins == "" {
		instance.AllowOrigins = "*"
	}

	if instance.AllowMethods == "" {
		instance.AllowMethods = "GET,POST,HEAD,PUT,DELETE,PATCH"
	}

	return &instance, nil
}
