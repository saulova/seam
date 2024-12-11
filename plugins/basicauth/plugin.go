package main

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/basicauth/middlewares"
)

type BasicAuthPlugin struct{}

func NewPlugin(appDependencyContainer *dependencies.DependencyContainer) interfaces.PluginInterface {
	dependencies.SetDependencyContainer(appDependencyContainer)

	instance := &BasicAuthPlugin{}

	return instance
}

func (s *BasicAuthPlugin) PluginBootstrap(config interface{}) {
	middlewares.NewBasicAuthMiddleware()
}
