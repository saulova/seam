---
title: Logger
description: A documentation page about the logger plugin.
tableOfContents: false
---

The Logger Plugin provides a middleware to log HTTP requests and responses. It offers customizable log formats, timezone adjustments, and performance optimizations, making it ideal for monitoring and debugging your application.

### Middlewares

#### plugins.logger.middlewares.LoggerMiddleware

- **Purpose**:  
  This middleware logs incoming requests and their corresponding responses, providing detailed insights into the application’s activity.

- **Usage**:  
  Attach the middleware to routes where you want to log activity. The logged information can be customized with various configuration options, ensuring relevance to your needs.

- **Configuration Integration**:  
  The middleware is configured through the plugin and referenced by name in the route’s middleware list.

##### Middleware Configuration

The following options are available for configuring the LoggerMiddleware:

- **Format**  
  Specifies the format of the log output. It supports placeholders for various request/response properties, such as method, path, status code, and latency. Default: `"${pid} | ${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n"`.

- **TimeZone**  
  Configures the timezone for timestamps in the logs. Default: `"UTC"`.

- **TimeInterval**  
  Sets the interval at which logs are grouped or buffered, specified as a duration (e.g., `1s`, `5s`). Default: `0` (no buffering).

- **DisableColors**  
  If `true`, disables colored output in the logs, useful for environments where plain text is preferred (e.g., log files, non-interactive terminals). Default: `false`.

##### Example Configuration

Below is an example of how to configure the Logger Plugin:

```yaml title="middlewares.yaml"
middlewares:
  LoggerMiddleware:
    use: plugins.logger.middlewares.LoggerMiddleware
    config:
      timeZone: "UTC"
      timeInterval: 5s
      disableColors: true
```

##### How It Works

1. **Request Logging**:  
   Logs details of every incoming request, including the method, path, and timestamp.

2. **Response Logging**:  
   Logs the status code, response latency, and other details once the request is processed.

3. **Custom Formats**:  
   Use the `format` configuration to include or exclude specific log details.

##### Using as Global Middleware

To use logger as global middleware, configure the middleware file to include the middleware as global middleware:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/logger.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  LoggerMiddleware:
    use: plugins.logger.middlewares.LoggerMiddleware
    config:
      format: "${pid} | ${locals:requestid} | ${time} | ${status} | ${ip} | ${method} | ${path} | ${error}\n",
      timeZone: "UTC"
  # ...
globalMiddlewares:
  - RequestIdMiddleware
  - LoggerMiddleware
```

##### Using with Routes

To enable logging for specific routes, configure it as follows:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/logger.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  LoggerMiddleware:
    use: plugins.logger.middlewares.LoggerMiddleware
    config:
      format: "${pid} | ${locals:requestid} | ${time} | ${status} | ${ip} | ${method} | ${path} | ${error}\n",
      timeZone: "UTC"
```

```yaml title="routes.yaml"
routes:
  # ...
  - path: "/api/example"
    method: "GET"
    middlewares:
      - LoggerMiddleware
    action: HttpForwardToInternalService
```

### Best Practices

- **Use Specific Formats**: Customize the log format to capture relevant details while keeping logs concise.
- **Adjust for Environment**: Disable colors (`disableColors: true`) when writing logs to non-interactive environments.
- **Combine with Monitoring Tools**: Integrate logs with monitoring solutions like ELK stack, Prometheus, ClickHouse or similar tools for centralized analysis.
