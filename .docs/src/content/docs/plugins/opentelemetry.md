---
title: Open Telemetry
description: A documentation page about the open telemetry plugin.
tableOfContents: false
---

The OpenTelemetry Plugin integrates tracing, metrics, and logging capabilities into your application, leveraging the OpenTelemetry framework. It provides a plugin for configuring the tracing provider and a middleware for capturing traces at the request level.

### Plugin Configuration

This configuration sets up the tracing provider, which determines how traces are recorded and exported. If no provider is specified, a **default provider** is used.

#### Configuration Parameters

- **Provider**  
  The name of the tracing provider to use. If left empty, the default provider is used.  
  Example: `"customProvider"`  
  Default: `""` (uses the default provider).

- **ProviderConfig**  
  Custom configuration for the specified provider. Its structure depends on the chosen provider.  
  Example: For OTLP, this could include endpoint settings.

- **ContextTracerKey**  
  The context key used to store the tracer instance in the request context for later use.  
  Example: `"tracer"`  
  Default: None.

#### Default Tracer Provider

If no custom provider is specified, the default provider is used with the following configuration parameters:

- **ServiceName**  
  The name of the service emitting the traces.  
  Example: `"my-service"`  
  Default: `""`.

- **UseOTLP**  
  Whether to use OpenTelemetry Protocol (OTLP) for exporting traces.  
  Example: `true`  
  Default: `false`.

- **DisableTraces**  
  Disables trace collection entirely when set to `true`.  
  Default: `false`.

- **DisableMetrics**  
  Disables metric collection when set to `true`.  
  Default: `false`.

- **DisableLogs**  
  Disables log collection when set to `true`.  
  Default: `false`.

#### Example Plugin Configuration

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/opentelemetry.so
    config:
      provider: ""
      providerConfig:
        serviceName: "example-service"
        useOTLP: true
      contextTracerKey: "tracer"
```

### Middlewares

#### plugins.opentelemetry.middlewares.OpenTelemetryMiddleware

This middleware integrates tracing into individual routes, ensuring that requests are traced and metrics are captured at the middleware level.

##### Configuration Parameters

- **TracerName**  
  The name of the tracer to use for the middleware. This name identifies the tracer instance that will record spans for the requests.
  Example: `"example-server"`

##### How It Works

1. Attaches a new span to each request that passes through the middleware.
2. Links the span to the configured tracer.
3. Ensures the span's lifecycle matches the request lifecycle, closing it at the end of processing.

#### Example Configuration

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/opentelemetry.so
    config:
      provider: ""
      providerConfig:
        serviceName: "example-service"
        useOTLP: true
        disableMetrics: false
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  OpenTelemetryMiddleware:
    use: plugins.opentelemetry.middlewares.OpenTelemetryMiddleware
    config:
      tracerName: "internal-service-endpoint"
```

```yaml title="routes.yaml"
routes:
  # ...
  - path: "/traced-endpoint"
    method: "GET"
    middlewares:
      - OpenTelemetryMiddleware
    action: HttpForwardToService
```

### Best Practices

- **Service Name**: Always specify a meaningful `serviceName` to differentiate traces from different services in a distributed system.
- **Combine with Metrics**: Leverage the full power of OpenTelemetry by enabling both tracing and metrics for comprehensive observability.
- **Use OTLP**: Prefer OTLP as the transport protocol for better integration with OpenTelemetry-compatible backends like Jaeger or Prometheus.
