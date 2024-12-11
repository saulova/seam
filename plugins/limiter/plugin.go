package main

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/limiter/middlewares"
)

type LimiterPlugin struct{}

func NewPlugin(appDependencyContainer *dependencies.DependencyContainer) interfaces.PluginInterface {
	dependencies.SetDependencyContainer(appDependencyContainer)

	instance := &LimiterPlugin{}

	return instance
}

func (s *LimiterPlugin) PluginBootstrap(config interface{}) {
	middlewares.NewLimiterMiddleware()
}
