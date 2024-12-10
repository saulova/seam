package application

import (
	"github.com/saulova/seam/infra/managers"
	"github.com/saulova/seam/infra/repositories/filesystem"

	"github.com/saulova/seam/application/commands"
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
)

type MiddlewaresMediator struct {
	loadMiddlewaresCommand           *commands.LoadMiddlewaresCommand
	registerGlobalMiddlewaresCommand *commands.RegisterGlobalMiddlewaresCommand
	logger                           interfaces.LoggerInterface
}

const MiddlewaresMediatorId = "mediators.application.MiddlewaresMediator"

func NewMiddlewaresMediator() *MiddlewaresMediator {
	dependencyContainer := dependencies.GetDependencyContainer()

	middlewaresRepository := dependencyContainer.GetDependency(filesystem.MiddlewaresRepositoryId).(*filesystem.MiddlewaresRepository)
	middlewaresManager := dependencyContainer.GetDependency(managers.MiddlewaresManagerId).(*managers.MiddlewaresManager)
	logger := dependencyContainer.GetDependency(interfaces.LoggerInterfaceId).(interfaces.LoggerInterface)

	loadMiddlewaresCommand := commands.NewLoadMiddlewaresCommand(middlewaresRepository, middlewaresManager, logger)
	registerGlobalMiddlewaresCommand := commands.NewRegisterGlobalMiddlewaresCommand(middlewaresRepository, middlewaresManager, logger)

	instance := &MiddlewaresMediator{
		loadMiddlewaresCommand:           loadMiddlewaresCommand,
		registerGlobalMiddlewaresCommand: registerGlobalMiddlewaresCommand,
		logger:                           logger,
	}

	dependencyContainer.AddDependency(MiddlewaresMediatorId, instance)

	return instance
}

func (a *MiddlewaresMediator) LoadMiddlewares() error {
	a.logger.Debug("executing middlewares mediator load middlewares")

	err := a.loadMiddlewaresCommand.Execute()
	if err != nil {
		return err
	}

	err = a.registerGlobalMiddlewaresCommand.Execute()
	if err != nil {
		return err
	}

	a.logger.Debug("middlewares mediator load middlewares executed")

	return nil
}
