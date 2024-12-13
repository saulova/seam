package configs

import (
	"github.com/mitchellh/mapstructure"

	"time"
)

type WebSocketActionConfig struct {
	PongTimeout       time.Duration `mapstructure:"pongTimeout"`
	PingInterval      time.Duration `mapstructure:"pongInterval"`
	MaxInactivity     time.Duration `mapstructure:"maxInactivity"`
	UpstreamEndpoint  string        `mapstructure:"upstreamEndpoint"`
	HandshakeTimeout  time.Duration `mapstructure:"handshakeTimeout"`
	Subprotocols      []string      `mapstructure:"subprotocols"`
	Origins           []string      `mapstructure:"origins"`
	ReadBufferSize    int           `mapstructure:"readBufferSize"`
	WriteBufferSize   int           `mapstructure:"writeBufferSize"`
	EnableCompression bool          `mapstructure:"enableCompression"`
}

func NewWebSocketActionConfig(config interface{}) (*WebSocketActionConfig, error) {
	var instance WebSocketActionConfig

	err := mapstructure.Decode(config, &instance)
	if err != nil {
		return nil, err
	}

	if instance.PongTimeout == 0 {
		instance.PongTimeout = 1 * time.Second
	}

	if instance.PingInterval == 0 {
		instance.PingInterval = 15 * time.Second
	}

	if instance.MaxInactivity == 0 {
		instance.MaxInactivity = 30 * time.Second
	}

	return &instance, nil
}
