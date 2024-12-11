package configs

import (
	"github.com/mitchellh/mapstructure"
)

type HelmetMiddlewareConfig struct {
	XSSProtection             string `mapstructure:"xssProtection"`
	ContentTypeNosniff        string `mapstructure:"contentTypeNosniff"`
	XFrameOptions             string `mapstructure:"xFrameOptions"`
	HSTSMaxAge                int    `mapstructure:"hstsMaxAge"`
	HSTSExcludeSubdomains     bool   `mapstructure:"hstsExcludeSubdomains"`
	ContentSecurityPolicy     string `mapstructure:"contentSecurityPolicy"`
	CSPReportOnly             bool   `mapstructure:"cspReportOnly"`
	HSTSPreloadEnabled        bool   `mapstructure:"hstsPreloadEnabled"`
	ReferrerPolicy            string `mapstructure:"referrerPolicy"`
	PermissionPolicy          string `mapstructure:"permissionPolicy"`
	CrossOriginEmbedderPolicy string `mapstructure:"crossOriginEmbedderPolicy"`
	CrossOriginOpenerPolicy   string `mapstructure:"crossOriginOpenerPolicy"`
	CrossOriginResourcePolicy string `mapstructure:"crossOriginResourcePolicy"`
	OriginAgentCluster        string `mapstructure:"originAgentCluster"`
	XDNSPrefetchControl       string `mapstructure:"xdnsPrefetchControl"`
	XDownloadOptions          string `mapstructure:"xDownloadOptions"`
	XPermittedCrossDomain     string `mapstructure:"xPermittedCrossDomain"`
}

func NewHelmetMiddlewareConfig(config interface{}) (*HelmetMiddlewareConfig, error) {
	var instance HelmetMiddlewareConfig

	err := mapstructure.Decode(config, &instance)
	if err != nil {
		return nil, err
	}

	if instance.XSSProtection == "" {
		instance.XSSProtection = "0"
	}

	if instance.ContentTypeNosniff == "" {
		instance.ContentTypeNosniff = "nosniff"
	}

	if instance.XFrameOptions == "" {
		instance.XFrameOptions = "SAMEORIGIN"
	}

	if instance.ReferrerPolicy == "" {
		instance.ReferrerPolicy = "ReferrerPolicy"
	}

	if instance.CrossOriginEmbedderPolicy == "" {
		instance.CrossOriginEmbedderPolicy = "require-corp"
	}

	if instance.CrossOriginOpenerPolicy == "" {
		instance.CrossOriginOpenerPolicy = "same-origin"
	}

	if instance.CrossOriginResourcePolicy == "" {
		instance.CrossOriginResourcePolicy = "same-origin"
	}

	if instance.OriginAgentCluster == "" {
		instance.OriginAgentCluster = "?1"
	}

	if instance.XDNSPrefetchControl == "" {
		instance.XDNSPrefetchControl = "off"
	}

	if instance.XDownloadOptions == "" {
		instance.XDownloadOptions = "noopen"
	}

	if instance.XPermittedCrossDomain == "" {
		instance.XPermittedCrossDomain = "none"
	}

	return &instance, nil
}
