package dtos

import (
	"github.com/saulova/seam/domain/entities/services"
)

type ListServicesOutput struct {
	Services map[string]*services.ServiceEntity
}
