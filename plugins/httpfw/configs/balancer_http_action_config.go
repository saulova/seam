package configs

import (
	"github.com/mitchellh/mapstructure"
)

type BalancerHttpActionConfig struct {
	UpstreamEndpoints []string `mapstructure:"upstreamEndpoints"`
}

func NewBalancerHttpActionConfig(config interface{}) (*BalancerHttpActionConfig, error) {
	var instance BalancerHttpActionConfig

	err := mapstructure.Decode(config, &instance)
	if err != nil {
		return nil, err
	}

	return &instance, nil
}
