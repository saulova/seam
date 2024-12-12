---
title: Limiter
description: A documentation page about the limiter plugin.
tableOfContents: false
---

The Limiter Plugin provides a middleware to control the rate of requests to your application, helping prevent abuse and ensuring resource fairness. It uses a token-bucket algorithm to manage the maximum number of allowed connections within a specified time window.

### Middlewares

#### plugins.limiter.middlewares.LimiterMiddleware

- **Purpose**:  
  This middleware enforces rate-limiting rules on routes where it is configured. It keeps track of requests and restricts access based on the defined limits.

- **Usage**:  
  Add the middleware to the desired routes to control the rate of requests. If a client exceeds the maximum allowed connections, the middleware responds with an appropriate status code (e.g., `429 Too Many Requests`) and halts further processing.

- **Configuration Integration**:  
  The middleware is configured through the plugin and referenced by name in the route's middleware list.

##### Middleware Configuration

The following options are available for configuring the LimiterMiddleware:

- **Storage**  
  Specifies the storage backend for keeping track of request counts (e.g., Valkey). This must be defined.

- **MaxConnections**  
  The maximum number of requests allowed within the defined expiration window. Default: `5`.

- **Expiration**  
  The time window for limiting requests, specified as a duration (e.g., `60s`, `2m`). Default: `60s`.

- **SkipFailedRequests**  
  If `true`, failed requests (e.g., those with status `>= 400`) are not counted against the rate limit. Default: `false`.

- **SkipSuccessfulRequests**  
  If `true`, successful requests (e.g., status `< 400`) are not counted against the rate limit. Default: `false`.

- **SlidingWindow**  
  If `true`, the limiter uses a sliding window algorithm instead of a fixed window, ensuring smoother enforcement. Default: `false`.

##### Example Configuration

Below is an example of how to configure the Limiter Plugin:

```yaml title="middlewares.yaml"
middlewares:
  LimiterMiddleware:
    use: plugins.limiter.middlewares.LimiterMiddleware
    config:
      storage: ValkeySessionStorage
      maxConnections: 50
      expiration: 60s
      skipFailedRequests: false
      skipSuccessfulRequests: false
      slidingWindow: true
```

##### How It Works

1. **Request Tracking**:  
   Each incoming request is checked against the stored request count in the configured storage backend.

2. **Limit Enforcement**:  
   If the request count exceeds the `maxConnections` within the specified `expiration` window:

   - The middleware responds with an HTTP status `429 Too Many Requests`.
   - A `Retry-After` header is sent, indicating when the client can retry.

3. **Sliding Window Option**:  
   When enabled, the middleware uses a sliding window algorithm to distribute requests more evenly across time, rather than resetting limits at the end of each time window.

##### Using as Global Middleware

To use rate limiting as global middleware, configure the middleware file to include the middleware as global middleware:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/limiter.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  LimiterMiddleware:
    use: plugins.limiter.middlewares.LimiterMiddleware
    config:
      storage: ValkeySessionStorage
      maxConnections: 20
      expiration: 30s
  # ...
globalMiddlewares:
  - LimiterMiddleware
```

##### Using with Routes

To apply rate limiting to a route, configure it as follows:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/limiter.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  LimiterMiddleware:
    use: plugins.limiter.middlewares.LimiterMiddleware
    config:
      storage: ValkeySessionStorage
      maxConnections: 20
      expiration: 30s
```

```yaml title="routes.yaml"
routes:
  # ...
  - path: "/api/resource"
    method: "GET"
    middlewares:
      - LimiterMiddleware
    action: HttpForwardToInternalService
```

### Best Practices

- **Use Appropriate Storage**: Choose a reliable and scalable storage backend (e.g., Valkey) for tracking request counts.
- **Adjust Limits Per Use Case**: Set different rate limits for sensitive routes, public APIs, or internal services based on expected traffic.
- **Combine with Other Plugins**: Use the Limiter Plugin alongside authentication or other middleware to ensure secure and fair access.
