package configs

import (
	"github.com/mitchellh/mapstructure"
)

type ValkeyStorageConfig struct {
	Prefix string `mapstructure:"prefix"`
	Host   string `mapstructure:"host"`
	Port   uint16 `mapstructure:"port"`
}

func NewValkeyStorageConfig(config interface{}) (*ValkeyStorageConfig, error) {
	var instance ValkeyStorageConfig

	err := mapstructure.Decode(config, &instance)
	if err != nil {
		return nil, err
	}

	if instance.Prefix == "" {
		instance.Prefix = "go:seam:"
	}

	return &instance, nil
}
