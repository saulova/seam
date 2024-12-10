package setup

import (
	"github.com/saulova/seam/mediators/application"

	"github.com/saulova/seam/infra/env"
	"github.com/saulova/seam/infra/logger"
	"github.com/saulova/seam/infra/managers"
	"github.com/saulova/seam/infra/plugins"
	"github.com/saulova/seam/infra/repositories/filesystem"
	"github.com/saulova/seam/infra/server"
	"github.com/saulova/seam/libs/dependencies"
)

func SetupApplication() {
	// Env Vars
	environmentVariables := env.NewEnvironmentVariables()

	// Logger
	logger.NewLogger(environmentVariables.LogLevel)

	// Dependency Container
	dependencyContainer := dependencies.GetDependencyContainer()

	// Configs
	storagesConfigFileHandler := filesystem.NewConfigFileHandler(environmentVariables.StoragesConfigPath)
	pluginsConfigFileHandler := filesystem.NewConfigFileHandler(environmentVariables.PluginsConfigPath)
	middlewaresConfigFileHandler := filesystem.NewConfigFileHandler(environmentVariables.MiddlewaresConfigPath)
	actionsConfigFileHandler := filesystem.NewConfigFileHandler(environmentVariables.ActionsConfigPath)
	routesConfigFileHandler := filesystem.NewConfigFileHandler(environmentVariables.RoutesConfigPath)
	servicesConfigFileHandler := filesystem.NewConfigFileHandler(environmentVariables.ServicesConfigPath)
	serverConfigFileHandler := filesystem.NewConfigFileHandler(environmentVariables.ServerConfigPath)

	dependencyContainer.AddDependency(filesystem.StoragesConfigFileHandlerId, storagesConfigFileHandler)
	dependencyContainer.AddDependency(filesystem.PluginsConfigFileHandlerId, pluginsConfigFileHandler)
	dependencyContainer.AddDependency(filesystem.MiddlewaresConfigFileHandlerId, middlewaresConfigFileHandler)
	dependencyContainer.AddDependency(filesystem.ActionsConfigFileHandlerId, actionsConfigFileHandler)
	dependencyContainer.AddDependency(filesystem.RoutesConfigFileHandlerId, routesConfigFileHandler)
	dependencyContainer.AddDependency(filesystem.ServicesConfigFileHandlerId, servicesConfigFileHandler)
	dependencyContainer.AddDependency(server.ServerConfigFileHandlerId, serverConfigFileHandler)

	// Filesystem
	filesystem.NewStoragesRepository()
	filesystem.NewPluginsRepository()
	filesystem.NewMiddlewaresRepository()
	filesystem.NewActionsRepository()
	filesystem.NewRoutesRepository()
	filesystem.NewServicesRepository()

	// Plugins
	plugins.NewPluginLoader()

	// Managers
	managers.NewStoragesManager()
	managers.NewMiddlewaresManager()
	managers.NewActionsManager()
	managers.NewRoutesManager()

	// API Server
	server.NewServerConfig("Seam - API Gateway")
	server.NewServerHandler()

	// Mediators
	application.NewPluginsMediator()
	application.NewStoragesMediator()
	application.NewMiddlewaresMediator()
	application.NewServicesMediator()
	application.NewActionsMediator()
}
