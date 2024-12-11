package actions

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/helpers"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/httpfw/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

type HttpAction struct{}

const HttpActionId = "plugins.httpfw.actions.HttpAction"

func NewHttpAction() interfaces.BuilderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	instance := &HttpAction{}

	dependencyContainer.AddDependency(HttpActionId, instance)

	return instance
}

func (h *HttpAction) Build(config interface{}) (interface{}, error) {
	httpActionConfig, err := configs.NewHttpActionConfig(config)
	if err != nil {
		return nil, err
	}

	forwardFunc := func(ctx *fiber.Ctx) error {
		uri := ctx.Request().URI()

		requestUrl := uri.String()

		url, err := helpers.GetTargetUrl(requestUrl, ctx.Path(), httpActionConfig.UpstreamEndpoint)
		if err != nil {
			return err
		}

		return proxy.Forward(url)(ctx)
	}

	return forwardFunc, nil
}
