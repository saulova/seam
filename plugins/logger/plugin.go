package main

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/logger/middlewares"
)

type LoggerPlugin struct{}

func NewPlugin(appDependencyContainer *dependencies.DependencyContainer) interfaces.PluginInterface {
	dependencies.SetDependencyContainer(appDependencyContainer)

	instance := &LoggerPlugin{}

	return instance
}

func (s *LoggerPlugin) PluginBootstrap(config interface{}) {
	middlewares.NewLoggerMiddleware()
}
