package application

import (
	"github.com/saulova/seam/infra/managers"
	"github.com/saulova/seam/infra/repositories/filesystem"

	"github.com/saulova/seam/application/commands"
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
)

type ServicesMediator struct {
	loadServicesCommand *commands.LoadServicesCommand
	logger              interfaces.LoggerInterface
}

const ServicesMediatorId = "mediators.application.ServicesMediator"

func NewServicesMediator() *ServicesMediator {
	dependencyContainer := dependencies.GetDependencyContainer()

	servicesRepository := dependencyContainer.GetDependency(filesystem.ServicesRepositoryId).(*filesystem.ServicesRepository)
	routesRepository := dependencyContainer.GetDependency(filesystem.RoutesRepositoryId).(*filesystem.RoutesRepository)
	routesManager := dependencyContainer.GetDependency(managers.RoutesManagerId).(*managers.RoutesManager)
	logger := dependencyContainer.GetDependency(interfaces.LoggerInterfaceId).(interfaces.LoggerInterface)

	loadServiceRoutesCommand := commands.NewLoadServiceRoutesCommand(routesRepository, routesManager, logger)
	loadServicesCommand := commands.NewLoadServicesCommand(servicesRepository, loadServiceRoutesCommand, logger)

	instance := &ServicesMediator{
		loadServicesCommand: loadServicesCommand,
		logger:              logger,
	}

	dependencyContainer.AddDependency(ServicesMediatorId, instance)

	return instance
}

func (s *ServicesMediator) LoadServices() error {
	s.logger.Debug("executing services mediator load services")

	err := s.loadServicesCommand.Execute()
	if err != nil {
		return err
	}

	s.logger.Debug("services mediator load services executed")

	return nil
}
