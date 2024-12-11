package middlewares

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/basicauth/configs"

	basicAuthMw "github.com/gofiber/fiber/v2/middleware/basicauth"
)

type BasicAuthMiddleware struct{}

const BasicAuthMiddlewareId = "plugins.basicauth.middlewares.BasicAuthMiddleware"

func NewBasicAuthMiddleware() interfaces.BuilderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	instance := &BasicAuthMiddleware{}

	dependencyContainer.AddDependency(BasicAuthMiddlewareId, instance)

	return instance
}

func (c *BasicAuthMiddleware) createBasicAuthConfig(config *configs.BasicAuthMiddlewareConfig) (basicAuthMw.Config, error) {
	basicAuthConfig := basicAuthMw.Config{
		Realm:           config.Realm,
		Users:           config.Users,
		ContextUsername: config.ContextUsernameKey,
		ContextPassword: config.ContextPasswordKey,
	}

	return basicAuthConfig, nil
}

func (c *BasicAuthMiddleware) Build(config interface{}) (interface{}, error) {
	basicAuthMiddlewareConfig, err := configs.NewBasicAuthMiddlewareConfig(config)
	if err != nil {
		return nil, err
	}

	basicAuthConfig, err := c.createBasicAuthConfig(basicAuthMiddlewareConfig)
	if err != nil {
		return nil, err
	}

	middlewareFunc := basicAuthMw.New(basicAuthConfig)

	return middlewareFunc, nil
}
