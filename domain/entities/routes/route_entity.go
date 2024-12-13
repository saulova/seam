package routes

import (
	"errors"
)

type RouteEntity struct {
	Id             string
	GatewayPath    string
	Methods        []string
	MiddlewaresIds []string
	UpstreamPath   string
	Action         string
}

func validate(input *RouteEntityInput) error {
	if input.GatewayPath == "" {
		return errors.New("invalid path")
	}

	if len(input.Methods) == 0 {
		return errors.New("missing methods")
	}

	if len(input.Methods) != 0 {
		for _, method := range input.Methods {
			if method == "" {
				return errors.New("invalid method")
			}
		}
	}

	return nil
}

func NewRouteEntity(input *RouteEntityInput) (*RouteEntity, error) {
	err := validate(input)
	if err != nil {
		return nil, err
	}

	instance := &RouteEntity{
		Id:             input.Id,
		GatewayPath:    input.GatewayPath,
		Methods:        input.Methods,
		MiddlewaresIds: input.MiddlewaresIds,
		Action:         input.Action,
	}

	return instance, nil
}
