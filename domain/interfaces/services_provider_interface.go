package interfaces

import (
	"github.com/saulova/seam/domain/dtos"
)

type ServicesProviderInterface interface {
	ListServices() (*dtos.ListServicesOutput, error)
}
