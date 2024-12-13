package main

import (
	"github.com/saulova/seam/mediators/application"

	"github.com/saulova/seam/infra/server"

	"github.com/saulova/seam/entrypoint/setup"
	"github.com/saulova/seam/libs/dependencies"
)

func main() {
	setup.SetupApplication()

	dependencyContainer := dependencies.GetDependencyContainer()

	pluginsMediator := dependencyContainer.GetDependency(application.PluginsMediatorId).(*application.PluginsMediator)
	middlewaresMediator := dependencyContainer.GetDependency(application.MiddlewaresMediatorId).(*application.MiddlewaresMediator)
	actionsMediator := dependencyContainer.GetDependency(application.ActionsMediatorId).(*application.ActionsMediator)
	servicesMediator := dependencyContainer.GetDependency(application.ServicesMediatorId).(*application.ServicesMediator)

	err := pluginsMediator.LoadPlugins()
	if err != nil {
		panic(err)
	}

	err = middlewaresMediator.LoadMiddlewares()
	if err != nil {
		panic(err)
	}

	err = actionsMediator.LoadActions()
	if err != nil {
		panic(err)
	}

	err = servicesMediator.LoadServices()
	if err != nil {
		panic(err)
	}

	serverHandler := dependencyContainer.GetDependency(server.ServerHandlerId).(*server.ServerHandler)
	serverHandler.SetHealthCheckReady(true)
	serverHandler.StartServer()
}
