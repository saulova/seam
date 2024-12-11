package main

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/httpfw/actions"
)

type HttpForwardPlugin struct{}

func NewPlugin(appDependencyContainer *dependencies.DependencyContainer) interfaces.PluginInterface {
	dependencies.SetDependencyContainer(appDependencyContainer)

	instance := &HttpForwardPlugin{}

	return instance
}

func (s *HttpForwardPlugin) PluginBootstrap(config interface{}) {
	actions.NewBalancerHttpAction()
	actions.NewHttpAction()
}
