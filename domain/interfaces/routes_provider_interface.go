package interfaces

import (
	"github.com/saulova/seam/domain/dtos"
)

type RoutesProviderInterface interface {
	ListRoutes() (*dtos.ListRoutesOutput, error)
	FindRouteById(id string) (*dtos.FindRouteByIdOutput, error)
}
