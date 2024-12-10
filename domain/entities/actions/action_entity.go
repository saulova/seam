package actions

import (
	"errors"
	"fmt"
	"strings"
)

type ActionEntity struct {
	Id     string
	Use    string
	Config interface{}
}

func validate(input *ActionEntityInput) error {
	if input.Id == "" {
		return errors.New("invalid action id")
	}

	if input.Use == "" || (strings.HasPrefix("plugins.", input.Use) && strings.HasSuffix("Action", input.Use)) {
		return fmt.Errorf("invalid use action: '%s'", input.Use)
	}

	return nil
}

func NewActionEntity(input *ActionEntityInput) (*ActionEntity, error) {
	err := validate(input)
	if err != nil {
		return nil, err
	}

	instance := &ActionEntity{
		Id:     input.Id,
		Use:    input.Use,
		Config: input.Config,
	}

	return instance, nil
}
