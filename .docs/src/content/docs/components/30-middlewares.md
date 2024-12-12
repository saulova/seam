---
title: Middlewares
description: A documentation page about middlewares.
tableOfContents: false
---

Middlewares components allows you to define reusable processing layers that are executed before an action is performed. Middlewares handle with cross-cutting concerns such as authentication, logging, rate limiting, and request transformation, acting as a pipeline through which every request passes.

### Authentication Middleware

Verifies the identity of the requestor using methods like API keys, JWT, or OAuth tokens.

**Use Case**: Protecting routes that require secure access.

Example Configuration:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/jwt.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  MiddlewareName:
    use: plugins.jwt.middlewares.JWTValidationMiddleware
    config:
      signingKey:
        algorithm: RS256
        key: S3cr3tP@ssw0rd
```

or

```yaml title="middlewares.yaml"
middlewares:
  # ...
  MiddlewareName:
    use: plugins.jwt.middlewares.JWTValidationMiddleware
    config:
      jwkUrls:
        - http://auth-server:8080/.well-known/jwks
```

### Rate Limiting Middleware

Controls the number of requests a client can make in a specific time period. Prevents abuse and ensures fair resource distribution.

**Use Case**: Throttling client requests to APIs.

Example Configuration:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/limiter.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  MiddlewareName:
    use: plugins.limiter.middlewares.LimiterMiddleware
    config:
      maxConnections: 5
      expiration: 1m
```

### Logging Middleware

Captures request and response details for monitoring and debugging.

**Use Case**: Recording transaction logs for auditing or troubleshooting.

Example Configuration:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/logger.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  1_MiddlewareName:
    use: plugins.logger.middlewares.LoggerMiddleware
    anyRequest: true
```

### CORS Middleware

Enables Cross-Origin Resource Sharing (CORS), allowing requests from different origins.

**Use Case**: APIs exposed to frontend applications hosted on other domains.

Example Configuration:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/cors.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  2_MiddlewareName:
    use: plugins.cors.middlewares.CORSMiddleware
    anyRequest: true
    config:
      allowOrigins: "*"
      allowMethods: "GET,POST,PUT,DELETE"
```

### Best Practices

- **Meaningful Middleware Names**: Use clear and descriptive names for middlewares that reflect their purpose.
- **Global Middlewares**: Use global middlewares for common functionality (e.g., logging, CORS).
- **Route-specific Middlewares**: Apply route-specific middlewares for security-sensitive or resource-intensive tasks.
