package middlewares

import (
	"fmt"

	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/session/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type MapSessionToHeaderMiddleware struct{}

const MapSessionToHeaderMiddlewareId = "plugins.session.middlewares.MapSessionToHeaderMiddleware"

func NewMapSessionToHeaderMiddleware() interfaces.BuilderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	instance := &MapSessionToHeaderMiddleware{}

	dependencyContainer.AddDependency(MapSessionToHeaderMiddlewareId, instance)

	return instance
}

func (s *MapSessionToHeaderMiddleware) Build(config interface{}) (interface{}, error) {
	sessionDataToHeaderMiddlewareConfig, err := configs.NewMapSessionToHeaderMiddlewareConfig(config)
	if err != nil {
		return nil, err
	}

	middlewareFunc := func(ctx *fiber.Ctx) error {
		sessionDataMap := ctx.Locals("session")

		if sessionDataMap == nil {
			log.Error("missing session")

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
				log.Error(fmt.Sprintf("unexpected type %T", v))
			}

			ctx.Request().Header.Set(headerKey, value)
		}

		return ctx.Next()
	}

	return middlewareFunc, nil
}
