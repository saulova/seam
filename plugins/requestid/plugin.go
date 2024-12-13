package main

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/requestid/middlewares"
)

type RequestIdPlugin struct{}

func NewPlugin(appDependencyContainer *dependencies.DependencyContainer) interfaces.PluginInterface {
	dependencies.SetDependencyContainer(appDependencyContainer)

	instance := &RequestIdPlugin{}

	return instance
}

func (s *RequestIdPlugin) PluginBootstrap(config interface{}) {
	middlewares.NewRequestIdMiddleware()
}
