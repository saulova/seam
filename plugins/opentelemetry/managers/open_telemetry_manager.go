package managers

import (
	"github.com/saulova/seam/libs/dependencies"
	libsInterfaces "github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/opentelemetry/configs"
	openTelemetryInterfaces "github.com/saulova/seam/plugins/opentelemetry/interfaces"
	"github.com/saulova/seam/plugins/opentelemetry/providers"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type OpenTelemetryManager struct {
	dependencyContainer       *dependencies.DependencyContainer
	openTelemetryPluginConfig *configs.OpenTelemetryPluginConfig
	logger                    libsInterfaces.LoggerInterface
}

const OpenTelemetryManagerId = "plugins.opentelemetry.managers.OpenTelemetryManager"

func NewOpenTelemetryManager() *OpenTelemetryManager {
	dependencyContainer := dependencies.GetDependencyContainer()

	openTelemetryPluginConfig := dependencyContainer.GetDependency(configs.OpenTelemetryPluginConfigId).(*configs.OpenTelemetryPluginConfig)
	logger := dependencyContainer.GetDependency(libsInterfaces.LoggerInterfaceId).(libsInterfaces.LoggerInterface)

	instance := &OpenTelemetryManager{
		dependencyContainer:       dependencyContainer,
		openTelemetryPluginConfig: openTelemetryPluginConfig,
		logger:                    logger,
	}

	dependencyContainer.AddDependency(OpenTelemetryManagerId, instance)

	return instance
}

func (o *OpenTelemetryManager) InitTracerProvider() {
	if o.openTelemetryPluginConfig.Provider != "" {
		telemetryProviders := o.dependencyContainer.GetDependency(o.openTelemetryPluginConfig.Provider).(openTelemetryInterfaces.TelemetryProviderInterface)

		err := telemetryProviders.Register(o.openTelemetryPluginConfig.ProviderConfig)
		if err != nil {
			o.logger.Error("open telemetry provider register error", err)
		}

		return
	}

	telemetryProviders := o.dependencyContainer.GetDependency(providers.DefaultTelemetryProvidersId).(openTelemetryInterfaces.TelemetryProviderInterface)

	err := telemetryProviders.Register(o.openTelemetryPluginConfig.ProviderConfig)
	if err != nil {
		o.logger.Error("open telemetry default provider register error", err)
	}
}

func (o *OpenTelemetryManager) NewTracer(name string) trace.Tracer {
	return otel.Tracer(name)
}
