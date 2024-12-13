package filesystem

import (
	"github.com/saulova/seam/domain/dtos"
	"github.com/saulova/seam/domain/entities/plugins"
	"github.com/saulova/seam/domain/interfaces"
	"github.com/saulova/seam/libs/dependencies"
)

type PluginSchema struct {
	Path   string      `yaml:"path"`
	Config interface{} `yaml:"config"`
}

type PluginsFileSchema struct {
	Plugins []PluginSchema `yaml:"plugins"`
}

type PluginsRepository struct {
	configFileHandler *ConfigFileHandler
	cachedPlugins     []*plugins.PluginEntity
}

const PluginsRepositoryId = "infra.filesystem.PluginsRepository"
const PluginsConfigFileHandlerId = "infra.filesystem.PluginsConfigFileHandler"

func NewPluginsRepository() interfaces.PluginsProviderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	configFileHandler := dependencyContainer.GetDependency(PluginsConfigFileHandlerId).(*ConfigFileHandler)

	instance := &PluginsRepository{
		configFileHandler: configFileHandler,
	}

	dependencyContainer.AddDependency(PluginsRepositoryId, instance)

	return instance
}

func (p *PluginsRepository) ListPlugins() (*dtos.ListPluginsOutput, error) {
	if len(p.cachedPlugins) > 0 {
		output := &dtos.ListPluginsOutput{
			Plugins: p.cachedPlugins,
		}

		return output, nil
	}

	var pluginsConfigs PluginsFileSchema

	err := p.configFileHandler.Unmarshal(&pluginsConfigs)
	if err != nil {
		panic(err)
	}

	result := make([]*plugins.PluginEntity, 0, len(pluginsConfigs.Plugins))

	for _, pluginConfig := range pluginsConfigs.Plugins {
		plugin, err := plugins.NewPluginEntity(&plugins.PluginEntityInput{
			Path:   pluginConfig.Path,
			Config: pluginConfig.Config,
		})
		if err != nil {
			return nil, err
		}

		result = append(result, plugin)
	}

	output := &dtos.ListPluginsOutput{
		Plugins: result,
	}

	p.cachedPlugins = result

	return output, nil
}
