package configs

import (
	"github.com/mitchellh/mapstructure"
)

type SigningKeyConfig struct {
	Algorithm      string `mapstructure:"algorithm"`
	Key            string `mapstructure:"key"`
	KeyAsByteSlice []byte
}

type JWTMiddlewareConfig struct {
	SigningKey SigningKeyConfig `mapstructure:"signingKey"`
	JwkUrls    []string         `mapstructure:"jwkUrls"`
}

func NewJWTMiddlewareConfig(config interface{}) (*JWTMiddlewareConfig, error) {
	var instance JWTMiddlewareConfig

	err := mapstructure.Decode(config, &instance)
	if err != nil {
		return nil, err
	}

	if instance.SigningKey.Key != "" {
		instance.SigningKey.KeyAsByteSlice = []byte(instance.SigningKey.Key)
	}

	return &instance, nil
}
