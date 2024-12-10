package commands

import (
	applicationInterfaces "github.com/saulova/seam/application/interfaces"
	domainInterfaces "github.com/saulova/seam/domain/interfaces"
	libsInterfaces "github.com/saulova/seam/libs/interfaces"
)

type LoadMiddlewaresCommand struct {
	middlewaresProvider domainInterfaces.MiddlewaresProviderInterface
	middlewaresManager  applicationInterfaces.MiddlewaresManagerInterface
	logger              libsInterfaces.LoggerInterface
}

func NewLoadMiddlewaresCommand(middlewaresProvider domainInterfaces.MiddlewaresProviderInterface, middlewaresManager applicationInterfaces.MiddlewaresManagerInterface, logger libsInterfaces.LoggerInterface) *LoadMiddlewaresCommand {
	return &LoadMiddlewaresCommand{
		middlewaresProvider: middlewaresProvider,
		middlewaresManager:  middlewaresManager,
		logger:              logger,
	}
}

func (l *LoadMiddlewaresCommand) Execute() error {
	l.logger.Debug("loading middlewares")

	listMiddlewaresOutput, err := l.middlewaresProvider.ListMiddlewares()
	if err != nil {
		l.logger.Error("list middlewares error", err)
		return err
	}

	l.logger.Debug("middlewares found", "output", listMiddlewaresOutput)

	for _, middleware := range listMiddlewaresOutput.Middlewares {
		l.logger.Debug("loading middleware", middleware)

		err = l.middlewaresManager.LoadMiddleware(middleware.Id, middleware.Use, middleware.Config)
		if err != nil {
			l.logger.Error("load middleware error", err)
			return err
		}

		l.logger.Debug("middleware loaded", middleware)
	}

	l.logger.Info("all middlewares loaded")

	return nil
}
