package configs

import (
	"github.com/saulova/seam/libs/dependencies"

	"github.com/mitchellh/mapstructure"
)

type OpenTelemetryPluginConfig struct {
	Provider         string      `mapstructure:"provider"`
	ProviderConfig   interface{} `mapstructure:"providerConfig"`
	ContextTracerKey string      `mapstructure:"contextTracerKey"`
}

const OpenTelemetryPluginConfigId = "plugins.opentelemetry.configs.OpenTelemetryPluginConfig"

func NewOpenTelemetryPluginConfig(config interface{}) (*OpenTelemetryPluginConfig, error) {
	dependencyContainer := dependencies.GetDependencyContainer()

	var instance OpenTelemetryPluginConfig

	err := mapstructure.Decode(config, &instance)
	if err != nil {
		return nil, err
	}

	if instance.ContextTracerKey == "" {
		instance.ContextTracerKey = "openTelemetryTracer"
	}

	dependencyContainer.AddDependency(OpenTelemetryPluginConfigId, &instance)

	return &instance, nil
}
