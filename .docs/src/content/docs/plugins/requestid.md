---
title: Request ID
description: A documentation page about the request id plugin.
tableOfContents: false
---

The Request ID Plugin provides a middleware to generate and manage unique identifiers for each HTTP request. These identifiers are critical for tracing requests through your system, especially in distributed environments.

### Middlewares

#### plugins.requestid.middlewares.RequestIdMiddleware

- **Purpose**:  
  This middleware ensures that each incoming request has a unique identifier (Request ID), either generating one or extracting it from a specified header. The Request ID is made available in the request context for use by subsequent middlewares or actions.

- **Usage**:  
  Attach this middleware to routes where you want to track requests using unique IDs. The generated or extracted Request ID can be logged, forwarded to downstream services, or used for debugging and monitoring.

- **Configuration Integration**:  
  The middleware is configured through the plugin and referenced by name in the route’s middleware list.

##### Middleware Configuration

The following options are available for configuring the RequestIdMiddleware:

- **HeaderName**  
  Specifies the name of the HTTP header to check for an existing Request ID. If the header is not present, the middleware generates a new ID. Default: `"X-Request-ID"`.

- **ContextKey**  
  The key used to store the Request ID in the request context (locals). Subsequent middlewares or actions can retrieve the Request ID using this key. Default: `"requestId"`.

##### Example Configuration

Below is an example of how to configure the Request ID Plugin:

```yaml title="middlewares.yaml"
middlewares:
  RequestIdMiddleware:
    use: plugins.requestid.middlewares.RequestIdMiddleware
    config:
      headerName: "X-Request-ID"
      contextKey: "requestId"
```

##### How It Works

1. **Header Check**:  
   The middleware checks the incoming request for a header specified by `headerName`. If the header is present, its value is used as the Request ID.

2. **ID Generation**:  
   If the header is absent, the middleware generates a new unique Request ID (e.g., a UUID).

3. **Context Injection**:  
   The Request ID is stored in the request context under the key defined by `contextKey`. This makes it accessible to other middlewares and actions.

##### Using as Global Middleware

To use request id as global middleware, configure the middleware file to include the middleware as global middleware:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/requestid.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  RequestIdMiddleware:
    use: plugins.requestid.middlewares.RequestIdMiddleware
    config:
      headerName: "X-Request-ID"
  # ...
globalMiddlewares:
  - RequestIdMiddleware
```

##### Using with Routes

To enable Request ID generation for a route, configure it as follows:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/requestid.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  RequestIdMiddleware:
    use: plugins.requestid.middlewares.RequestIdMiddleware
    config:
      headerName: "X-Request-ID"
```

```yaml title="routes.yaml"
routes:
  # ...
  - path: "/api/resource"
    method: "GET"
    middlewares:
      - RequestIdMiddleware
    action: HttpForwardToInternalService
```

### Best Practices

- **Consistent Header Usage**: Standardize the use of a Request ID header across your services (e.g. `"X-Request-ID"`).
- **Traceability**: Pass the Request ID to downstream services and include it in logs for end-to-end tracing.
- **Combine with Logger Middleware**: Pair this middleware with the Logger Plugin to include Request IDs in log entries, aiding in debugging and correlation.
