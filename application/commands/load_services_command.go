package commands

import (
	domainInterfaces "github.com/saulova/seam/domain/interfaces"
	libsInterfaces "github.com/saulova/seam/libs/interfaces"
)

type LoadServicesCommand struct {
	servicesProvider         domainInterfaces.ServicesProviderInterface
	loadServiceRoutesCommand *LoadServiceRoutesCommand
	logger                   libsInterfaces.LoggerInterface
}

func NewLoadServicesCommand(servicesProvider domainInterfaces.ServicesProviderInterface, loadServiceRoutesCommand *LoadServiceRoutesCommand, logger libsInterfaces.LoggerInterface) *LoadServicesCommand {
	return &LoadServicesCommand{
		servicesProvider:         servicesProvider,
		loadServiceRoutesCommand: loadServiceRoutesCommand,
		logger:                   logger,
	}
}

func (l *LoadServicesCommand) Execute() error {
	l.logger.Debug("loading services")

	listServicesOutput, err := l.servicesProvider.ListServices()
	if err != nil {
		l.logger.Error("list services error", err)
		return err
	}

	l.logger.Debug("services found", "output", listServicesOutput)

	for _, service := range listServicesOutput.Services {
		err := l.loadServiceRoutesCommand.Execute(service)
		if err != nil {
			return err
		}
	}

	l.logger.Info("all services loaded")

	return nil
}
