package main

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/cors/middlewares"
)

type CORSPlugin struct{}

func NewPlugin(appDependencyContainer *dependencies.DependencyContainer) interfaces.PluginInterface {
	dependencies.SetDependencyContainer(appDependencyContainer)

	instance := &CORSPlugin{}

	return instance
}

func (s *CORSPlugin) PluginBootstrap(config interface{}) {
	middlewares.NewCORSMiddleware()
}
