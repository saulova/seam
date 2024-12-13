package main

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/helmet/middlewares"
)

type HelmetPlugin struct{}

func NewPlugin(appDependencyContainer *dependencies.DependencyContainer) interfaces.PluginInterface {
	dependencies.SetDependencyContainer(appDependencyContainer)

	instance := &HelmetPlugin{}

	return instance
}

func (s *HelmetPlugin) PluginBootstrap(config interface{}) {
	middlewares.NewHelmetMiddleware()
}
