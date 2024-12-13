package main

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/opentelemetry/configs"
	"github.com/saulova/seam/plugins/opentelemetry/managers"
	"github.com/saulova/seam/plugins/opentelemetry/middlewares"
	"github.com/saulova/seam/plugins/opentelemetry/providers"
)

type OpenTelemetryPlugin struct{}

func NewPlugin(appDependencyContainer *dependencies.DependencyContainer) interfaces.PluginInterface {
	dependencies.SetDependencyContainer(appDependencyContainer)

	instance := &OpenTelemetryPlugin{}

	return instance
}

func (s *OpenTelemetryPlugin) PluginBootstrap(config interface{}) {
	configs.NewOpenTelemetryPluginConfig(config)

	providers.NewDefaultTelemetryProviders()

	openTelemetryManager := managers.NewOpenTelemetryManager()
	openTelemetryManager.InitTracerProvider()

	middlewares.NewOpenTelemetryMiddleware()
}
