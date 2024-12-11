package main

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/jwt/middlewares"
)

type JWTPlugin struct{}

func NewPlugin(appDependencyContainer *dependencies.DependencyContainer) interfaces.PluginInterface {
	dependencies.SetDependencyContainer(appDependencyContainer)

	instance := &JWTPlugin{}

	return instance
}

func (s *JWTPlugin) PluginBootstrap(config interface{}) {
	middlewares.NewJWTValidationMiddleware()
}
