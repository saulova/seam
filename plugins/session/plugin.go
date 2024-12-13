package main

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/session/managers"
	"github.com/saulova/seam/plugins/session/middlewares"
)

type SessionPlugin struct{}

func NewPlugin(appDependencyContainer *dependencies.DependencyContainer) interfaces.PluginInterface {
	dependencies.SetDependencyContainer(appDependencyContainer)

	instance := &SessionPlugin{}

	return instance
}

func (s *SessionPlugin) PluginBootstrap(config interface{}) {
	managers.NewSessionManager()
	middlewares.NewLoadSessionMiddleware()
	middlewares.NewMapSessionToHeaderMiddleware()
}
