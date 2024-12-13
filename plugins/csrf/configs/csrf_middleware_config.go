package configs

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

type CSRFMiddlewareConfig struct {
	KeyLookup         string        `mapstructure:"keyLookup"`
	CookieName        string        `mapstructure:"cookieName"`
	CookieDomain      string        `mapstructure:"cookieDomain"`
	CookiePath        string        `mapstructure:"cookiePath"`
	CookieSecure      bool          `mapstructure:"cookieSecure"`
	CookieHTTPOnly    bool          `mapstructure:"cookieHTTPOnly"`
	CookieSameSite    string        `mapstructure:"cookieSameSite"`
	CookieSessionOnly bool          `mapstructure:"cookieSessionOnly"`
	CookieExpiration  time.Duration `mapstructure:"cookieExpiration"`
	SingleUseToken    bool          `mapstructure:"singleUseToken"`
	Storage           string        `mapstructure:"storage"`
}

func NewCSRFMiddlewareConfig(config interface{}) (*CSRFMiddlewareConfig, error) {
	var instance CSRFMiddlewareConfig

	err := mapstructure.Decode(config, &instance)
	if err != nil {
		return nil, err
	}

	if instance.KeyLookup == "" {
		instance.KeyLookup = "header:X-Csrf-Token"
	}

	if instance.CookieName == "" {
		instance.CookieName = "csrf_"
	}

	if instance.CookieSameSite == "" {
		instance.CookieSameSite = "Lax"
	}

	if instance.CookieExpiration == time.Duration(0) {
		instance.CookieExpiration = 1 * time.Hour
	}

	return &instance, nil
}
