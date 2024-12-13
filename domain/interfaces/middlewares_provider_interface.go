package interfaces

import (
	"github.com/saulova/seam/domain/dtos"
)

type MiddlewaresProviderInterface interface {
	ListMiddlewares() (*dtos.ListMiddlewaresOutput, error)
	ListGlobalMiddlewares() (*dtos.ListGlobalMiddlewaresOutput, error)
	FindMiddlewareById(id string) (*dtos.FindMiddlewareByIdOutput, error)
}
