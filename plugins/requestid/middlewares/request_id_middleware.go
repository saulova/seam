package middlewares

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/requestid/configs"

	reqId "github.com/gofiber/fiber/v2/middleware/requestid"
)

type RequestIdMiddleware struct{}

const RequestIdMiddlewareId = "plugins.requestid.middlewares.RequestIdMiddleware"

func NewRequestIdMiddleware() interfaces.BuilderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	instance := &RequestIdMiddleware{}

	dependencyContainer.AddDependency(RequestIdMiddlewareId, instance)

	return instance
}

func (c *RequestIdMiddleware) createRequestIdConfig(config *configs.RequestIdMiddlewareConfig) (reqId.Config, error) {
	requestIdConfig := reqId.Config{
		Header:     config.HeaderName,
		ContextKey: config.ContextKey,
	}

	return requestIdConfig, nil
}

func (c *RequestIdMiddleware) Build(config interface{}) (interface{}, error) {
	requestIdMiddlewareConfig, err := configs.NewRequestIdMiddlewareConfig(config)
	if err != nil {
		return nil, err
	}

	requestIdConfig, err := c.createRequestIdConfig(requestIdMiddlewareConfig)
	if err != nil {
		return nil, err
	}

	middlewareFunc := reqId.New(requestIdConfig)

	return middlewareFunc, nil
}
