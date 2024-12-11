package actions

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/helpers"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/wsfw/configs"
	"github.com/saulova/seam/plugins/wsfw/ws"

	"github.com/gofiber/fiber/v2"
)

type WebSocketAction struct {
	webSocketForward *ws.WebSocketForward
}

const WebSocketActionId = "plugins.wsfw.actions.WebSocketAction"

func NewWebSocketAction() interfaces.BuilderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	webSocketForward := dependencyContainer.GetDependency(ws.WebSocketForwardId).(*ws.WebSocketForward)

	instance := &WebSocketAction{
		webSocketForward: webSocketForward,
	}

	dependencyContainer.AddDependency(WebSocketActionId, instance)

	return instance
}

func (w *WebSocketAction) Build(config interface{}) (interface{}, error) {
	webSocketActionConfig, err := configs.NewWebSocketActionConfig(config)
	if err != nil {
		return nil, err
	}

	actionFunc := func(ctx *fiber.Ctx) error {
		requestUrl := ctx.Request().URI().String()

		url, err := helpers.GetTargetUrl(requestUrl, ctx.Path(), webSocketActionConfig.UpstreamEndpoint)
		if err != nil {
			return err
		}

		return w.webSocketForward.GetForwardAction(webSocketActionConfig, url)(ctx)
	}

	return actionFunc, nil
}
