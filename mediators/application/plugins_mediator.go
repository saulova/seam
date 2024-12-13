package application

import (
	"github.com/saulova/seam/infra/plugins"
	"github.com/saulova/seam/infra/repositories/filesystem"

	"github.com/saulova/seam/application/commands"
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
)

type PluginsMediator struct {
	loadPluginsCommand *commands.LoadPluginsCommand
	logger             interfaces.LoggerInterface
}

const PluginsMediatorId = "mediators.application.PluginsMediator"

func NewPluginsMediator() *PluginsMediator {
	dependencyContainer := dependencies.GetDependencyContainer()

	pluginsRepository := dependencyContainer.GetDependency(filesystem.PluginsRepositoryId).(*filesystem.PluginsRepository)
	pluginLoader := dependencyContainer.GetDependency(plugins.PluginLoaderId).(*plugins.PluginLoader)
	logger := dependencyContainer.GetDependency(interfaces.LoggerInterfaceId).(interfaces.LoggerInterface)

	loadPluginsCommand := commands.NewLoadPluginsCommand(pluginsRepository, pluginLoader, logger)

	instance := &PluginsMediator{
		loadPluginsCommand: loadPluginsCommand,
		logger:             logger,
	}

	dependencyContainer.AddDependency(PluginsMediatorId, instance)

	return instance
}

func (p *PluginsMediator) LoadPlugins() error {
	p.logger.Debug("executing plugins mediator load plugins")

	err := p.loadPluginsCommand.Execute()
	if err != nil {
		return err
	}

	p.logger.Debug("plugins mediator load plugins executed")

	return nil
}
