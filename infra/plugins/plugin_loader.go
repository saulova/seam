package plugins

import (
	"fmt"
	"plugin"

	applicationInterfaces "github.com/saulova/seam/application/interfaces"
	"github.com/saulova/seam/libs/dependencies"
	libsInterfaces "github.com/saulova/seam/libs/interfaces"
)

type PluginLoader struct {
	dependencyContainer *dependencies.DependencyContainer
}

const PluginLoaderId = "infra.plugins.PluginLoader"

func NewPluginLoader() applicationInterfaces.PluginLoaderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	instance := &PluginLoader{
		dependencyContainer: dependencyContainer,
	}

	dependencyContainer.AddDependency(PluginLoaderId, instance)

	return instance
}

func (p *PluginLoader) LoadPlugin(pluginPath string) (libsInterfaces.PluginInterface, error) {
	plug, err := plugin.Open(pluginPath)
	if err != nil {
		return nil, fmt.Errorf("plugin error: %s", err.Error())
	}

	newPluginFunc, err := plug.Lookup("NewPlugin")
	if err != nil {
		return nil, fmt.Errorf("missing NewPlugin func for: %s", pluginPath)
	}

	pluginInstance := newPluginFunc.(func(appDependencyContainer *dependencies.DependencyContainer) libsInterfaces.PluginInterface)(p.dependencyContainer)

	return pluginInstance, nil
}
