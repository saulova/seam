package dtos

import (
	"github.com/saulova/seam/domain/entities/routes"
)

type ListRoutesOutput struct {
	Routes map[string]*routes.RouteEntity
}
type FindRouteByIdOutput struct {
	Route *routes.RouteEntity
}
