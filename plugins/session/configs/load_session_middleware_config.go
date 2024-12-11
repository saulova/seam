package configs

import (
	"time"

	"github.com/saulova/seam/libs/dependencies"

	"github.com/mitchellh/mapstructure"
)

type LoadSessionMiddlewareConfig struct {
	Storage               string        `mapstructure:"storage"`
	KeyLookup             string        `mapstructure:"keyLookup"`
	CookieHTTPOnly        bool          `mapstructure:"cookieHTTPOnly"`
	CookieSecure          bool          `mapstructure:"cookieSecure"`
	CookieSameSite        string        `mapstructure:"cookieSameSite"`
	CookieSessionOnly     bool          `mapstructure:"cookieSessionOnly"`
	CookieExpiration      time.Duration `mapstructure:"cookieExpiration"`
	DisableAutoRenew      bool          `mapstructure:"disableAutoRenew"`
	AutoRenewAfter        time.Duration `mapstructure:"autoRenewAfter"`
	DisableSessionForward bool          `mapstructure:"disableSessionForward"`
}

const LoadSessionMiddlewareConfigId = "plugins.session.configs.LoadSessionMiddlewareConfig"

func NewLoadSessionMiddlewareConfig(config interface{}) (*LoadSessionMiddlewareConfig, error) {
	dependencyContainer := dependencies.GetDependencyContainer()

	var instance LoadSessionMiddlewareConfig

	err := mapstructure.Decode(config, &instance)
	if err != nil {
		return nil, err
	}

	if instance.KeyLookup == "" {
		instance.KeyLookup = "cookie:session_id"
	}

	if instance.AutoRenewAfter == time.Duration(0) {
		instance.AutoRenewAfter = 30 * time.Minute
	}

	if instance.CookieSameSite == "" {
		instance.CookieSameSite = "Lax"
	}

	if instance.CookieExpiration == time.Duration(0) {
		instance.CookieExpiration = 2 * time.Hour
	}

	dependencyContainer.AddDependency(LoadSessionMiddlewareConfigId, &instance)

	return &instance, nil
}
