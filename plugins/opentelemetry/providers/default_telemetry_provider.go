package providers

import (
	"context"
	"time"

	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/plugins/opentelemetry/configs"
	"github.com/saulova/seam/plugins/opentelemetry/interfaces"

	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type DefaultTelemetryProvider struct {
	config *configs.DefaultTracerProviderConfig
}

const DefaultTelemetryProvidersId = "plugins.opentelemetry.providers.DefaultTelemetryProvider"

func NewDefaultTelemetryProviders() interfaces.TelemetryProviderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	instance := &DefaultTelemetryProvider{
		config: nil,
	}

	dependencyContainer.AddDependency(DefaultTelemetryProvidersId, instance)

	return instance
}

func (d *DefaultTelemetryProvider) setConfig(config interface{}) error {
	defaultTracerProviderConfig, err := configs.NewDefaultTracerProviderConfig(config)
	if err != nil {
		return nil
	}

	d.config = defaultTracerProviderConfig

	return nil
}

func (d *DefaultTelemetryProvider) newResource() (*resource.Resource, error) {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(d.config.ServiceName),
	), nil
}

func (d *DefaultTelemetryProvider) newTraceExporter() (sdktrace.SpanExporter, error) {
	if d.config.UseOTLP {
		traceExporter, err := otlptracegrpc.New(context.Background())
		if err != nil {
			return nil, err
		}

		return traceExporter, nil
	}

	traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	return traceExporter, nil
}

func (d *DefaultTelemetryProvider) newMetricExporter() (sdkmetric.Exporter, error) {
	if d.config.UseOTLP {
		metricExporter, err := otlpmetricgrpc.New(context.Background())
		if err != nil {
			return nil, err
		}

		return metricExporter, nil
	}

	metricExporter, err := stdoutmetric.New()
	if err != nil {
		return nil, err
	}

	return metricExporter, nil
}

func (d *DefaultTelemetryProvider) newLogExporter() (sdklog.Exporter, error) {
	if d.config.UseOTLP {
		logExporter, err := otlploggrpc.New(context.Background())
		if err != nil {
			return nil, err
		}

		return logExporter, nil
	}

	logExporter, err := stdoutlog.New()
	if err != nil {
		return nil, err
	}

	return logExporter, nil
}

func (d *DefaultTelemetryProvider) newTracerProvider(res *resource.Resource) (*sdktrace.TracerProvider, error) {
	traceExporter, err := d.newTraceExporter()
	if err != nil {
		return nil, err
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithResource(res),
	)

	return tracerProvider, nil
}

func (d *DefaultTelemetryProvider) newMeterProvider(res *resource.Resource) (*sdkmetric.MeterProvider, error) {
	metricExporter, err := d.newMetricExporter()
	if err != nil {
		return nil, err
	}

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(
				metricExporter,
				sdkmetric.WithInterval(1*time.Minute),
				sdkmetric.WithProducer(runtime.NewProducer()),
			),
		),
	)

	err = runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Minute))
	if err != nil {
		return nil, err
	}

	return meterProvider, nil
}

func (d *DefaultTelemetryProvider) newLoggerProvider(res *resource.Resource) (*sdklog.LoggerProvider, error) {
	logExporter, err := d.newLogExporter()
	if err != nil {
		return nil, err
	}

	loggerProvider := sdklog.NewLoggerProvider(
		sdklog.WithResource(res),
		sdklog.WithProcessor(
			sdklog.NewBatchProcessor(logExporter),
		),
	)

	return loggerProvider, nil
}

func (d *DefaultTelemetryProvider) Register(config interface{}) error {
	d.setConfig(config)

	res, err := d.newResource()
	if err != nil {
		return err
	}

	if !d.config.DisableTraces {
		tracerProvider, err := d.newTracerProvider(res)
		if err != nil {
			return err
		}

		otel.SetTracerProvider(tracerProvider)
	}

	if !d.config.DisableMetrics {
		meterProvider, err := d.newMeterProvider(res)
		if err != nil {
			return err
		}

		otel.SetMeterProvider(meterProvider)
	}

	if !d.config.DisableLogs {
		loggerProvider, err := d.newLoggerProvider(res)
		if err != nil {
			return err
		}

		global.SetLoggerProvider(loggerProvider)
	}

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return nil
}
