package main

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/wsfw/actions"
	"github.com/saulova/seam/plugins/wsfw/middlewares"
	"github.com/saulova/seam/plugins/wsfw/ws"
)

type WebSocketForwardPlugin struct{}

func NewPlugin(appDependencyContainer *dependencies.DependencyContainer) interfaces.PluginInterface {
	dependencies.SetDependencyContainer(appDependencyContainer)

	instance := &WebSocketForwardPlugin{}

	return instance
}

func (s *WebSocketForwardPlugin) PluginBootstrap(config interface{}) {
	ws.NewWebSocketForward()
	middlewares.NewCheckWebSocketUpgradeMiddleware()
	actions.NewWebSocketAction()
}
