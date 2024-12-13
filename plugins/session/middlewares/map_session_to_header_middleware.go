package middlewares

import (
	"fmt"

	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/session/configs"

	"github.com/gofiber/fiber/v2"
)

type MapSessionToHeaderMiddleware struct {
	logger interfaces.LoggerInterface
}

const MapSessionToHeaderMiddlewareId = "plugins.session.middlewares.MapSessionToHeaderMiddleware"

func NewMapSessionToHeaderMiddleware() interfaces.BuilderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	logger := dependencyContainer.GetDependency(interfaces.LoggerInterfaceId).(interfaces.LoggerInterface)

	instance := &MapSessionToHeaderMiddleware{
		logger: logger,
	}

	dependencyContainer.AddDependency(MapSessionToHeaderMiddlewareId, instance)

	return instance
}

func (m *MapSessionToHeaderMiddleware) Build(config interface{}) (interface{}, error) {
	sessionDataToHeaderMiddlewareConfig, err := configs.NewMapSessionToHeaderMiddlewareConfig(config)
	if err != nil {
		return nil, err
	}

	middlewareFunc := func(ctx *fiber.Ctx) error {
		sessionDataMap := ctx.Locals("session")

		if sessionDataMap == nil {
			m.logger.Error("missing session")

			return ctx.Next()
		}

		for headerKey, sessionKey := range sessionDataToHeaderMiddlewareConfig.Headers {
			sessionData := sessionDataMap.(map[string]interface{})[sessionKey]

			var value string

			switch v := sessionData.(type) {
			case string:
				value = sessionData.(string)
			case int:
				value = string(sessionData.(int))
			case nil:
				value = ""
			default:
				m.logger.Error(fmt.Sprintf("unexpected type %T", v))
			}

			ctx.Request().Header.Set(headerKey, value)
		}

		return ctx.Next()
	}

	return middlewareFunc, nil
}
