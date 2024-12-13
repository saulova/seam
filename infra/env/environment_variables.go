package env

import (
	"os"
)

type EnvironmentVariables struct {
	StoragesConfigPath    string
	PluginsConfigPath     string
	ServerConfigPath      string
	ServicesConfigPath    string
	MiddlewaresConfigPath string
	RoutesConfigPath      string
	ActionsConfigPath     string
	LogLevel              string
}

func NewEnvironmentVariables() *EnvironmentVariables {
	storagesConfigPath := os.Getenv("STORAGES_CONFIG_PATH")
	if storagesConfigPath == "" {
		storagesConfigPath = "/configs/storages.yaml"
	}

	pluginsConfigPath := os.Getenv("PLUGINS_CONFIG_PATH")
	if pluginsConfigPath == "" {
		pluginsConfigPath = "/configs/plugins.yaml"
	}

	servicesConfigPath := os.Getenv("SERVICES_CONFIG_PATH")
	if servicesConfigPath == "" {
		servicesConfigPath = "/configs/services.yaml"
	}

	middlewaresConfigPath := os.Getenv("MIDDLEWARES_CONFIG_PATH")
	if middlewaresConfigPath == "" {
		middlewaresConfigPath = "/configs/middlewares.yaml"
	}

	routesConfigPath := os.Getenv("ROUTES_CONFIG_PATH")
	if routesConfigPath == "" {
		routesConfigPath = "/configs/routes.yaml"
	}

	actionsConfigPath := os.Getenv("SERVER_CONFIG_PATH")
	if actionsConfigPath == "" {
		actionsConfigPath = "/configs/actions.yaml"
	}

	serverConfigPath := os.Getenv("SERVER_CONFIG_PATH")
	if serverConfigPath == "" {
		serverConfigPath = "/configs/server.yaml"
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "ERROR"
	}

	return &EnvironmentVariables{
		StoragesConfigPath:    storagesConfigPath,
		PluginsConfigPath:     pluginsConfigPath,
		ServerConfigPath:      serverConfigPath,
		ServicesConfigPath:    servicesConfigPath,
		MiddlewaresConfigPath: middlewaresConfigPath,
		RoutesConfigPath:      routesConfigPath,
		ActionsConfigPath:     actionsConfigPath,
		LogLevel:              logLevel,
	}
}
