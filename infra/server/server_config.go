package server

import (
	"fmt"
	"time"

	"github.com/saulova/seam/infra/repositories/filesystem"
	"github.com/saulova/seam/libs/dependencies"
)

type ServerConfig struct {
	Address                      string        `yaml:"address"`
	TLS                          bool          `yaml:"tls"`
	CertFile                     string        `yaml:"certFile"`
	KeyFile                      string        `yaml:"keyFile"`
	DisableHealthCheck           bool          `yaml:"disableHealthCheck"`
	HealthCheckLiveRoute         string        `yaml:"healthCheckLiveRoute"`
	HealthCheckReadyRoute        string        `yaml:"healthCheckReadyRoute"`
	Prefork                      bool          `yaml:"prefork"`
	ServerHeader                 string        `yaml:"serverHeader"`
	StrictRouting                bool          `yaml:"strictRouting"`
	CaseSensitive                bool          `yaml:"caseSensitive"`
	UnescapePath                 bool          `yaml:"unescapePath"`
	ETag                         bool          `yaml:"eTag"`
	BodyLimit                    int           `yaml:"bodyLimit"`
	Concurrency                  int           `yaml:"concurrency"`
	ReadTimeout                  time.Duration `yaml:"readTimeout"`
	WriteTimeout                 time.Duration `yaml:"writeTimeout"`
	IdleTimeout                  time.Duration `yaml:"idleTimeout"`
	ReadBufferSize               int           `yaml:"readBufferSize"`
	WriteBufferSize              int           `yaml:"writeBufferSize"`
	CompressedFileSuffix         string        `yaml:"compressedFileSuffix"`
	ProxyHeader                  string        `yaml:"proxyHeader"`
	GETOnly                      bool          `yaml:"getOnly"`
	DisableKeepalive             bool          `yaml:"disableKeepalive"`
	DisableDefaultDate           bool          `yaml:"disableDefaultDate"`
	DisableDefaultContentType    bool          `yaml:"disableDefaultContentType"`
	DisableHeaderNormalizing     bool          `yaml:"disableHeaderNormalizing"`
	DisableStartupMessage        bool          `yaml:"disableStartupMessage"`
	AppName                      string        `yaml:"appName"`
	StreamRequestBody            bool          `yaml:"streamRequestBody"`
	DisablePreParseMultipartForm bool          `yaml:"disablePreParseMultipartForm"`
	ReduceMemoryUsage            bool          `yaml:"reduceMemoryUsage"`
	Network                      string        `yaml:"network"`
	EnableTrustedProxyCheck      bool          `yaml:"enableTrustedProxyCheck"`
	TrustedProxies               []string      `yaml:"trustedProxies"`
	EnableIPValidation           bool          `yaml:"enableIPValidation"`
	EnablePrintRoutes            bool          `yaml:"enablePrintRoutes"`
	EnableSplittingOnParsers     bool          `yaml:"enableSplittingOnParsers"`
}

type ServerFileSchema struct {
	Server ServerConfig `yaml:"server"`
}

const ServerConfigId = "infra.api.server.ServerConfig"
const ServerConfigFileHandlerId = "infra.api.server.ServerConfigFileHandlerId"

func NewServerConfig(appName string) *ServerConfig {
	dependencyContainer := dependencies.GetDependencyContainer()

	configFileHandler := dependencyContainer.GetDependency(ServerConfigFileHandlerId).(*filesystem.ConfigFileHandler)

	var instance ServerFileSchema

	err := configFileHandler.Unmarshal(&instance)
	if err != nil {
		panic(err)
	}

	if instance.Server.CertFile == "" {
		instance.Server.CertFile = "/certs/server.cert"
	}

	if instance.Server.KeyFile == "" {
		instance.Server.KeyFile = "/certs/server.key"
	}

	if instance.Server.Address == "" {
		instance.Server.Address = "0.0.0.0:8090"
	}

	if instance.Server.HealthCheckLiveRoute == "" {
		instance.Server.HealthCheckLiveRoute = "/health/live"
	}

	if instance.Server.HealthCheckReadyRoute == "" {
		instance.Server.HealthCheckReadyRoute = "/health/ready"
	}

	if instance.Server.BodyLimit == 0 {
		instance.Server.BodyLimit = 4 * 1024 * 1024
	}
	if instance.Server.Concurrency == 0 {
		instance.Server.Concurrency = 256 * 1024
	}
	if instance.Server.ReadBufferSize == 0 {
		instance.Server.ReadBufferSize = 4096
	}
	if instance.Server.WriteBufferSize == 0 {
		instance.Server.WriteBufferSize = 4096
	}

	if instance.Server.CompressedFileSuffix == "" {
		instance.Server.CompressedFileSuffix = ".seam.gz"
	}

	if instance.Server.AppName != "" {
		instance.Server.AppName = fmt.Sprintf("%s - %s", appName, instance.Server.AppName)
	}

	if instance.Server.AppName == "" {
		instance.Server.AppName = appName
	}

	if instance.Server.Network == "" {
		instance.Server.Network = "tcp4"
	}

	dependencyContainer.AddDependency(ServerConfigId, &instance.Server)

	return &instance.Server
}
