package middlewares

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/cors/configs"

	corsMw "github.com/gofiber/fiber/v2/middleware/cors"
)

type CORSMiddleware struct{}

const CORSMiddlewareId = "plugins.cors.middlewares.CORSMiddleware"

func NewCORSMiddleware() interfaces.BuilderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	instance := &CORSMiddleware{}

	dependencyContainer.AddDependency(CORSMiddlewareId, instance)

	return instance
}

func (c *CORSMiddleware) createCORSConfig(config *configs.CORSMiddlewareConfig) (corsMw.Config, error) {
	corsConfig := corsMw.Config{
		AllowOrigins:     config.AllowOrigins,
		AllowMethods:     config.AllowMethods,
		AllowHeaders:     config.AllowHeaders,
		AllowCredentials: config.AllowCredentials,
		ExposeHeaders:    config.ExposeHeaders,
		MaxAge:           config.MaxAge,
	}

	return corsConfig, nil
}

func (c *CORSMiddleware) Build(config interface{}) (interface{}, error) {
	corsMiddlewareConfig, err := configs.NewCORSMiddlewareConfig(config)
	if err != nil {
		return nil, err
	}

	corsConfig, err := c.createCORSConfig(corsMiddlewareConfig)
	if err != nil {
		return nil, err
	}

	middlewareFunc := corsMw.New(corsConfig)

	return middlewareFunc, nil
}
