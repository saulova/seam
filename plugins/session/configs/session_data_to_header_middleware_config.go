package configs

import (
	"github.com/saulova/seam/libs/dependencies"

	"github.com/mitchellh/mapstructure"
)

type MapSessionToHeaderMiddlewareConfig struct {
	Headers map[string]string `mapstructure:"headers"`
}

const MapSessionToHeaderMiddlewareConfigId = "plugins.session.configs.MapSessionToHeaderMiddlewareConfig"

func NewMapSessionToHeaderMiddlewareConfig(config interface{}) (*MapSessionToHeaderMiddlewareConfig, error) {
	dependencyContainer := dependencies.GetDependencyContainer()

	var instance MapSessionToHeaderMiddlewareConfig

	err := mapstructure.Decode(config, &instance)
	if err != nil {
		return nil, err
	}

	dependencyContainer.AddDependency(MapSessionToHeaderMiddlewareConfigId, &instance)

	return &instance, nil
}
