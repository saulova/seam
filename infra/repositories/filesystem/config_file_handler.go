package filesystem

import (
	"fmt"
	"os"

	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"

	"gopkg.in/yaml.v3"
)

type ConfigFileHandler struct {
	logger     interfaces.LoggerInterface
	configPath string
	configFile []byte
	loaded     bool
}

const ConfigFileHandlerId = "libs.filesystem.ConfigFileHandler"

func NewConfigFileHandler(configPath string) *ConfigFileHandler {
	dependencyContainer := dependencies.GetDependencyContainer()

	logger := dependencyContainer.GetDependency(interfaces.LoggerInterfaceId).(interfaces.LoggerInterface)

	instance := &ConfigFileHandler{
		logger:     logger,
		configPath: configPath,
		configFile: []byte{},
		loaded:     false,
	}

	dependencyContainer.AddDependency(ConfigFileHandlerId, instance)

	return instance
}

func (c *ConfigFileHandler) loadConfigFile() error {
	c.logger.Debug("loading config file", "configPath", c.configPath)

	configFile, err := os.ReadFile(c.configPath)
	if err != nil {
		c.logger.Error("can not read config file", err)
		return fmt.Errorf("can not read config file: %w", err)
	}

	c.configFile = configFile

	c.loaded = true

	c.logger.Debug("config file loaded", "configPath", c.configPath, "configFile", configFile)

	return nil
}

func (c *ConfigFileHandler) Unmarshal(out interface{}) error {
	if !c.loaded {
		c.loadConfigFile()
	}

	c.logger.Debug("unmarshal config file", "configPath", c.configPath)

	err := yaml.Unmarshal(c.configFile, out)
	if err != nil {
		c.logger.Error("can not unmarshal config file", err)
		return err
	}

	c.logger.Debug("config file decoded", out)

	return nil
}
