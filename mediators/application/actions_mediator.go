package application

import (
	"github.com/saulova/seam/infra/managers"
	"github.com/saulova/seam/infra/repositories/filesystem"

	"github.com/saulova/seam/application/commands"
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
)

type ActionsMediator struct {
	loadActionsCommand *commands.LoadActionsCommand
	logger             interfaces.LoggerInterface
}

const ActionsMediatorId = "mediators.application.ActionsMediator"

func NewActionsMediator() *ActionsMediator {
	dependencyContainer := dependencies.GetDependencyContainer()

	actionsRepository := dependencyContainer.GetDependency(filesystem.ActionsRepositoryId).(*filesystem.ActionsRepository)
	actionsManager := dependencyContainer.GetDependency(managers.ActionsManagerId).(*managers.ActionsManager)
	logger := dependencyContainer.GetDependency(interfaces.LoggerInterfaceId).(interfaces.LoggerInterface)

	loadActionsCommand := commands.NewLoadActionsCommand(actionsRepository, actionsManager, logger)

	instance := &ActionsMediator{
		loadActionsCommand: loadActionsCommand,
		logger:             logger,
	}

	dependencyContainer.AddDependency(ActionsMediatorId, instance)

	return instance
}

func (a *ActionsMediator) LoadActions() error {
	a.logger.Debug("executing actions mediator load actions")
	err := a.loadActionsCommand.Execute()
	if err != nil {
		return err
	}

	a.logger.Debug("actions mediator load actions executed")

	return nil
}
