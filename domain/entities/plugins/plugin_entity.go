package plugins

import (
	"errors"
)

type PluginEntity struct {
	Path   string
	Config interface{}
}

func validate(input *PluginEntityInput) error {
	if input.Path == "" {
		return errors.New("invalid plugin path")
	}

	return nil
}

func NewPluginEntity(input *PluginEntityInput) (*PluginEntity, error) {
	err := validate(input)
	if err != nil {
		return nil, err
	}

	instance := &PluginEntity{
		Path:   input.Path,
		Config: input.Config,
	}

	return instance, nil
}
