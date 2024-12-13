package commands

import (
	"github.com/saulova/seam/domain/entities/services"

	applicationInterfaces "github.com/saulova/seam/application/interfaces"
	domainInterfaces "github.com/saulova/seam/domain/interfaces"
	libsInterfaces "github.com/saulova/seam/libs/interfaces"
)

type LoadServiceRoutesCommand struct {
	routesProvider domainInterfaces.RoutesProviderInterface
	routesManager  applicationInterfaces.RoutesManagerInterface
	logger         libsInterfaces.LoggerInterface
}

func NewLoadServiceRoutesCommand(routesProvider domainInterfaces.RoutesProviderInterface, routesManager applicationInterfaces.RoutesManagerInterface, logger libsInterfaces.LoggerInterface) *LoadServiceRoutesCommand {
	return &LoadServiceRoutesCommand{
		routesProvider: routesProvider,
		routesManager:  routesManager,
		logger:         logger,
	}
}

func (l *LoadServiceRoutesCommand) Execute(service *services.ServiceEntity) error {
	l.logger.Debug("loading service routes", service)

	for _, routeId := range service.RoutesIds {
		l.logger.Debug("loading service route", "serviceId", service.Id, "routeId", routeId)

		findRouteByIdOutput, err := l.routesProvider.FindRouteById(routeId)
		if err != nil {
			l.logger.Error("find route by id error", err)
			return err
		}

		l.logger.Debug("route found by id", "serviceId", service.Id, "output", findRouteByIdOutput)

		err = l.routesManager.RegisterRoute(service, findRouteByIdOutput.Route)
		if err != nil {
			l.logger.Error("register route error", err)
			return err
		}

		l.logger.Debug("route registered", "serviceId", service.Id, "routeId", routeId)
	}

	l.logger.Info("all service routes loaded", "serviceId", service.Id)

	return nil
}
