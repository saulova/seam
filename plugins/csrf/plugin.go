package main

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/csrf/middlewares"
)

type CSRFPlugin struct{}

func NewPlugin(appDependencyContainer *dependencies.DependencyContainer) interfaces.PluginInterface {
	dependencies.SetDependencyContainer(appDependencyContainer)

	instance := &CSRFPlugin{}

	return instance
}

func (s *CSRFPlugin) PluginBootstrap(config interface{}) {
	middlewares.NewCSRFMiddleware()
}
