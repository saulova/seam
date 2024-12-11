package actions

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/helpers"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/httpfw/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

type BalancerHttpAction struct{}

const BalancerHttpActionId = "plugins.httpfw.actions.BalancerHttpAction"

func NewBalancerHttpAction() interfaces.BuilderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	instance := &BalancerHttpAction{}

	dependencyContainer.AddDependency(BalancerHttpActionId, instance)

	return instance
}

func (h *BalancerHttpAction) Build(config interface{}) (interface{}, error) {
	balancerHttpActionConfig, err := configs.NewBalancerHttpActionConfig(config)
	if err != nil {
		return nil, err
	}

	forwardFunc := func(ctx *fiber.Ctx) error {
		requestUrl := ctx.Request().URI().String()

		urls := make([]string, 0, len(balancerHttpActionConfig.UpstreamEndpoints))

		for _, endpoint := range balancerHttpActionConfig.UpstreamEndpoints {
			url, err := helpers.GetTargetUrl(requestUrl, ctx.Path(), endpoint)
			if err != nil {
				return err
			}

			urls = append(urls, url)
		}

		return proxy.BalancerForward(urls)(ctx)
	}

	return forwardFunc, nil
}
