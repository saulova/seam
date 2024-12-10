package commands

import (
	applicationInterfaces "github.com/saulova/seam/application/interfaces"
	domainInterfaces "github.com/saulova/seam/domain/interfaces"
	libsInterfaces "github.com/saulova/seam/libs/interfaces"
)

type RegisterGlobalMiddlewaresCommand struct {
	middlewaresProvider domainInterfaces.MiddlewaresProviderInterface
	middlewaresManager  applicationInterfaces.MiddlewaresManagerInterface
	logger              libsInterfaces.LoggerInterface
}

func NewRegisterGlobalMiddlewaresCommand(middlewaresProvider domainInterfaces.MiddlewaresProviderInterface, middlewaresManager applicationInterfaces.MiddlewaresManagerInterface, logger libsInterfaces.LoggerInterface) *RegisterGlobalMiddlewaresCommand {
	return &RegisterGlobalMiddlewaresCommand{
		middlewaresProvider: middlewaresProvider,
		middlewaresManager:  middlewaresManager,
		logger:              logger,
	}
}

func (r *RegisterGlobalMiddlewaresCommand) Execute() error {
	r.logger.Debug("registering middlewares")

	listGlobalMiddlewareOutput, err := r.middlewaresProvider.ListGlobalMiddlewares()
	if err != nil {
		r.logger.Error("list global middlewares error", err)
		return err
	}

	r.logger.Debug("global middlewares found", "output", listGlobalMiddlewareOutput)

	for _, middlewareId := range listGlobalMiddlewareOutput.GlobalMiddlewares {
		r.logger.Debug("registering global middleware", "id", middlewareId)

		err = r.middlewaresManager.RegisterGlobalMiddleware(middlewareId)
		if err != nil {
			r.logger.Error("register global middlewares error", err)
			return err
		}

		r.logger.Debug("global middleware registered", "id", middlewareId)
	}

	r.logger.Info("all middlewares registered")

	return nil
}
