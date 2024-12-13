package configs

import (
	"github.com/mitchellh/mapstructure"
)

type DefaultTracerProviderConfig struct {
	ServiceName    string `mapstructure:"serviceName"`
	UseOTLP        bool   `mapstructure:"useOTLP"`
	DisableTraces  bool   `mapstructure:"disableTraces"`
	DisableMetrics bool   `mapstructure:"disableMetrics"`
	DisableLogs    bool   `mapstructure:"disableLogs"`
}

func NewDefaultTracerProviderConfig(config interface{}) (*DefaultTracerProviderConfig, error) {
	var instance DefaultTracerProviderConfig

	err := mapstructure.Decode(config, &instance)
	if err != nil {
		return nil, err
	}

	if instance.ServiceName == "" {
		instance.ServiceName = "seam-api-gateway"
	}

	return &instance, nil
}
