package services

import (
	"errors"
)

type ServiceEntity struct {
	Id              string
	MiddlewaresIds  []string
	GatewayBasePath string
	RoutesIds       []string
}

func validate(input *ServiceEntityInput) error {
	if len(input.RoutesIds) == 0 {
		return errors.New("missing routes")
	}

	return nil
}

func NewServiceEntity(input *ServiceEntityInput) (*ServiceEntity, error) {
	err := validate(input)
	if err != nil {
		return nil, err
	}

	instance := &ServiceEntity{
		Id:              input.Id,
		MiddlewaresIds:  input.MiddlewaresIds,
		GatewayBasePath: input.GatewayBasePath,
		RoutesIds:       input.RoutesIds,
	}

	return instance, nil
}
