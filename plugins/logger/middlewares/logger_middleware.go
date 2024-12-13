package middlewares

import (
	"time"

	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/logger/configs"

	loggerMw "github.com/gofiber/fiber/v2/middleware/logger"
)

type LoggerMiddleware struct{}

const LoggerMiddlewareId = "plugins.logger.middlewares.LoggerMiddleware"

func NewLoggerMiddleware() interfaces.BuilderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	instance := &LoggerMiddleware{}

	dependencyContainer.AddDependency(LoggerMiddlewareId, instance)

	return instance
}

func (c *LoggerMiddleware) createLoggerConfig(config *configs.LoggerMiddlewareConfig) (loggerMw.Config, error) {
	loggerConfig := loggerMw.Config{
		Format:        config.Format,
		TimeZone:      config.TimeZone,
		TimeInterval:  config.TimeInterval,
		DisableColors: config.DisableColors,
		TimeFormat:    time.RFC3339,
	}

	return loggerConfig, nil
}

func (c *LoggerMiddleware) Build(config interface{}) (interface{}, error) {
	loggerMiddlewareConfig, err := configs.NewLoggerMiddlewareConfig(config)
	if err != nil {
		return nil, err
	}

	loggerConfig, err := c.createLoggerConfig(loggerMiddlewareConfig)
	if err != nil {
		return nil, err
	}

	middlewareFunc := loggerMw.New(loggerConfig)

	return middlewareFunc, nil
}
