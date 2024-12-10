package managers

import (
	"path"

	"github.com/gofiber/fiber/v2"
	"github.com/saulova/seam/application/interfaces"
	"github.com/saulova/seam/domain/entities/routes"
	"github.com/saulova/seam/domain/entities/services"
	"github.com/saulova/seam/infra/server"
	"github.com/saulova/seam/libs/dependencies"
)

type RoutesManager struct {
	middlewaresManager *MiddlewaresManager
	actionsManager     *ActionsManager
}

const RoutesManagerId = "infra.api.RoutesManager"

func NewRoutesManager() interfaces.RoutesManagerInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	middlewaresManager := dependencyContainer.GetDependency(MiddlewaresManagerId).(*MiddlewaresManager)
	actionsManager := dependencyContainer.GetDependency(ActionsManagerId).(*ActionsManager)

	instance := &RoutesManager{
		middlewaresManager: middlewaresManager,
		actionsManager:     actionsManager,
	}

	dependencyContainer.AddDependency(RoutesManagerId, instance)

	return instance
}

func (r *RoutesManager) RegisterRoute(service *services.ServiceEntity, route *routes.RouteEntity) error {
	routeRegister := server.NewRouteRegister()

	routePath := path.Join(service.GatewayBasePath, route.GatewayPath)

	routeRegister.UsePath(routePath)
	routeRegister.UseMethods(route.Methods)

	middlewaresIds := append(service.MiddlewaresIds, route.MiddlewaresIds...)
	middlewares := make([]func(ctx *fiber.Ctx) error, 0)

	for _, middlewareId := range middlewaresIds {
		middlewares = append(middlewares, r.middlewaresManager.GetMiddleware(middlewareId).(func(ctx *fiber.Ctx) error))
	}

	routeRegister.UseMiddlewares(middlewares)

	if route.Action != "" {
		action := r.actionsManager.GetAction(route.Action).(func(ctx *fiber.Ctx) error)

		routeRegister.UseAction(action)
	}

	routeRegister.Register()

	return nil
}
