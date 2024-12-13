package middlewares

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type CheckWebSocketUpgradeMiddleware struct{}

const CheckWebSocketUpgradeMiddlewareId = "plugins.wsfw.middlewares.CheckWebSocketUpgradeMiddleware"

func NewCheckWebSocketUpgradeMiddleware() interfaces.BuilderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	instance := &CheckWebSocketUpgradeMiddleware{}

	dependencyContainer.AddDependency(CheckWebSocketUpgradeMiddlewareId, instance)

	return instance
}

func (g *CheckWebSocketUpgradeMiddleware) Build(config interface{}) (interface{}, error) {
	middlewareFunc := func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("wsAllowed", true)
			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	}

	return middlewareFunc, nil
}
