package managers

import (
	"log"

	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/plugins/opentelemetry/configs"
	"github.com/saulova/seam/plugins/opentelemetry/interfaces"
	"github.com/saulova/seam/plugins/opentelemetry/providers"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type OpenTelemetryManager struct {
	dependencyContainer       *dependencies.DependencyContainer
	openTelemetryPluginConfig *configs.OpenTelemetryPluginConfig
}

const OpenTelemetryManagerId = "plugins.opentelemetry.managers.OpenTelemetryManager"

func NewOpenTelemetryManager() *OpenTelemetryManager {
	dependencyContainer := dependencies.GetDependencyContainer()

	openTelemetryPluginConfig := dependencyContainer.GetDependency(configs.OpenTelemetryPluginConfigId).(*configs.OpenTelemetryPluginConfig)

	instance := &OpenTelemetryManager{
		dependencyContainer:       dependencyContainer,
		openTelemetryPluginConfig: openTelemetryPluginConfig,
	}

	dependencyContainer.AddDependency(OpenTelemetryManagerId, instance)

	return instance
}

func (o *OpenTelemetryManager) InitTracerProvider() {
	if o.openTelemetryPluginConfig.Provider != "" {
		telemetryProviders := o.dependencyContainer.GetDependency(o.openTelemetryPluginConfig.Provider).(interfaces.TelemetryProviderInterface)

		err := telemetryProviders.Register(o.openTelemetryPluginConfig.ProviderConfig)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	telemetryProviders := o.dependencyContainer.GetDependency(providers.DefaultTelemetryProvidersId).(interfaces.TelemetryProviderInterface)

	err := telemetryProviders.Register(o.openTelemetryPluginConfig.ProviderConfig)
	if err != nil {
		log.Fatal(err)
	}

}

func (o *OpenTelemetryManager) NewTracer(name string) trace.Tracer {
	return otel.Tracer(name)
}
