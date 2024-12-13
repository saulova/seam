package middlewares

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/opentelemetry/configs"
	"github.com/saulova/seam/plugins/opentelemetry/managers"

	openTelemetryMw "github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
)

type OpenTelemetryMiddleware struct {
	openTelemetryPluginConfig *configs.OpenTelemetryPluginConfig
	openTelemetryManager      *managers.OpenTelemetryManager
}

const OpenTelemetryMiddlewareId = "plugins.opentelemetry.middlewares.OpenTelemetryMiddleware"

func NewOpenTelemetryMiddleware() interfaces.BuilderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	openTelemetryPluginConfig := dependencyContainer.GetDependency(configs.OpenTelemetryPluginConfigId).(*configs.OpenTelemetryPluginConfig)
	openTelemetryManager := dependencyContainer.GetDependency(managers.OpenTelemetryManagerId).(*managers.OpenTelemetryManager)

	instance := &OpenTelemetryMiddleware{
		openTelemetryPluginConfig: openTelemetryPluginConfig,
		openTelemetryManager:      openTelemetryManager,
	}

	dependencyContainer.AddDependency(OpenTelemetryMiddlewareId, instance)

	return instance
}

func (o *OpenTelemetryMiddleware) Build(config interface{}) (interface{}, error) {
	openTelemetryMiddlewareConfig, err := configs.NewOpenTelemetryMiddlewareConfig(config)
	if err != nil {
		return nil, err
	}

	tracer := o.openTelemetryManager.NewTracer(openTelemetryMiddlewareConfig.TracerName)

	middlewareFunc := func(ctx *fiber.Ctx) error {
		ctx.Locals(o.openTelemetryPluginConfig.ContextTracerKey, tracer)

		return openTelemetryMw.Middleware()(ctx)
	}

	return middlewareFunc, nil
}
