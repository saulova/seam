package interfaces

import (
	"github.com/saulova/seam/domain/entities/routes"
	"github.com/saulova/seam/domain/entities/services"
)

type RoutesManagerInterface interface {
	RegisterRoute(service *services.ServiceEntity, route *routes.RouteEntity) error
}
