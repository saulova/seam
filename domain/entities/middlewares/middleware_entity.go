package middlewares

import (
	"errors"
	"fmt"
	"strings"
)

type MiddlewareEntity struct {
	Id     string
	Use    string
	Config interface{}
}

func validate(input *MiddlewareEntityInput) error {
	if input.Id == "" {
		return errors.New("invalid middleware id")
	}

	if input.Use == "" || (strings.HasPrefix("plugins.", input.Use) && strings.HasSuffix("Middleware", input.Use)) {
		return fmt.Errorf("invalid use middleware: '%s'", input.Use)
	}

	return nil
}

func NewMiddlewareEntity(input *MiddlewareEntityInput) (*MiddlewareEntity, error) {
	err := validate(input)
	if err != nil {
		return nil, err
	}

	instance := &MiddlewareEntity{
		Id:     input.Id,
		Use:    input.Use,
		Config: input.Config,
	}

	return instance, nil
}
