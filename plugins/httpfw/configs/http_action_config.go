package configs

import (
	"github.com/mitchellh/mapstructure"
)

type HttpActionConfig struct {
	UpstreamEndpoint string `mapstructure:"upstreamEndpoint"`
}

func NewHttpActionConfig(config interface{}) (*HttpActionConfig, error) {
	var instance HttpActionConfig

	err := mapstructure.Decode(config, &instance)
	if err != nil {
		return nil, err
	}

	return &instance, nil
}
