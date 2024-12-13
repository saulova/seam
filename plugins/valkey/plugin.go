package main

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/valkey/storages"
)

type ValkeyPlugin struct{}

func NewPlugin(appDependencyContainer *dependencies.DependencyContainer) interfaces.PluginInterface {
	dependencies.SetDependencyContainer(appDependencyContainer)

	instance := &ValkeyPlugin{}

	return instance
}

func (s *ValkeyPlugin) PluginBootstrap(config interface{}) {
	storages.NewValkeyStorage()
}
