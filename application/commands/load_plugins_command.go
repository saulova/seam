package commands

import (
	applicationInterfaces "github.com/saulova/seam/application/interfaces"
	domainInterfaces "github.com/saulova/seam/domain/interfaces"
	libsInterfaces "github.com/saulova/seam/libs/interfaces"
)

type LoadPluginsCommand struct {
	pluginsProvider domainInterfaces.PluginsProviderInterface
	PluginLoader    applicationInterfaces.PluginLoaderInterface
	logger          libsInterfaces.LoggerInterface
}

func NewLoadPluginsCommand(pluginsProvider domainInterfaces.PluginsProviderInterface, PluginLoader applicationInterfaces.PluginLoaderInterface, logger libsInterfaces.LoggerInterface) *LoadPluginsCommand {
	return &LoadPluginsCommand{
		pluginsProvider: pluginsProvider,
		PluginLoader:    PluginLoader,
		logger:          logger,
	}
}

func (l *LoadPluginsCommand) Execute() error {
	l.logger.Debug("loading plugins")

	listPluginsOutput, err := l.pluginsProvider.ListPlugins()
	if err != nil {
		l.logger.Error("list plugins error", err)
		return err
	}

	l.logger.Debug("plugins found", "output", listPluginsOutput)

	for _, plugin := range listPluginsOutput.Plugins {
		l.logger.Debug("loading plugin", plugin)

		loadedPlugin, err := l.PluginLoader.LoadPlugin(plugin.Path)
		if err != nil {
			l.logger.Error("load plugin error", err)
			return err
		}

		loadedPlugin.PluginBootstrap(plugin.Config)

		l.logger.Debug("plugin loaded", plugin)
	}

	l.logger.Info("all plugins loaded")

	return nil
}
