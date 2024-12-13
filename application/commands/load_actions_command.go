package commands

import (
	applicationInterfaces "github.com/saulova/seam/application/interfaces"
	domainInterfaces "github.com/saulova/seam/domain/interfaces"
	libsInterfaces "github.com/saulova/seam/libs/interfaces"
)

type LoadActionsCommand struct {
	actionsProvider domainInterfaces.ActionsProviderInterface
	actionsManager  applicationInterfaces.ActionsManagerInterface
	logger          libsInterfaces.LoggerInterface
}

func NewLoadActionsCommand(actionsProvider domainInterfaces.ActionsProviderInterface, actionsManager applicationInterfaces.ActionsManagerInterface, logger libsInterfaces.LoggerInterface) *LoadActionsCommand {
	return &LoadActionsCommand{
		actionsProvider: actionsProvider,
		actionsManager:  actionsManager,
		logger:          logger,
	}
}

func (l *LoadActionsCommand) Execute() error {
	l.logger.Debug("loading actions")

	listActionsOutput, err := l.actionsProvider.ListActions()
	if err != nil {
		l.logger.Error("list actions error", err)
		return err
	}

	l.logger.Debug("actions found", "output", listActionsOutput)

	for _, action := range listActionsOutput.Actions {
		l.logger.Debug("loading action", action)

		err = l.actionsManager.LoadAction(action.Id, action.Use, action.Config)
		if err != nil {
			l.logger.Error("load action error", err)
			return err
		}

		l.logger.Debug("action loaded")
	}

	l.logger.Info("all actions loaded")

	return nil
}
