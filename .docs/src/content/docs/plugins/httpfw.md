---
title: HTTP Forward
description: A documentation page about the HTTP Forward plugin.
tableOfContents: false
---

The HTTP Forward Plugin provides two actions for forwarding HTTP requests to upstream services. These actions allow seamless integration of load balancing or single-endpoint forwarding into your application's routing layer.

### Actions

#### plugins.httpfw.actions.HttpAction

- **Purpose**:  
  This action forwards HTTP requests to a single upstream endpoint. It is designed for use cases where requests need to be routed directly to a specific service or server.

- **Usage**:  
  Attach this action to a route to forward incoming requests to the defined upstream endpoint. It is suitable for scenarios such as API gateways, proxies, or microservices architectures with a one-to-one mapping of routes to services.

- **Configuration Integration**:  
  The action is configured through the plugin by specifying the `upstreamEndpoint`. The route's `action` field references this action by name.

##### Configuration Parameters

- **UpstreamEndpoint**  
  The URL of the upstream service to which requests will be forwarded.  
  Example: `"http://localhost:8080"`.  
  Default: None (must be configured).

##### Example Configuration

Below is an example of how to configure the HTTP Forward Plugin:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/httpfw.so
```

```yaml title="actions.yaml"
actions:
  # ...
  ForwardToService:
    use: plugins.httpfw.actions.HttpAction
    config:
      upstreamEndpoint: "http://localhost:8080"
```

```yaml title="routes.yaml"
routes:
  # ...
  - path: "/api/resource"
    method: "GET"
    action: ForwardToService
```

#### plugins.httpfw.actions.BalancerHttpAction

- **Purpose**:  
  This action forwards HTTP requests to one of several upstream endpoints using a load-balancing strategy. It is ideal for distributing requests across multiple service instances.

- **Usage**:  
  Attach this action to a route to enable load-balanced forwarding to multiple upstream endpoints. It is suitable for high-availability and scalable architectures.

- **Configuration Integration**:  
  The action is configured through the plugin by specifying a list of `upstreamEndpoints`. The route's `action` field references this action by name.

##### Configuration Parameters

- **UpstreamEndpoints**  
  A list of URLs representing the upstream services. Requests are distributed among these services.  
  Example:
  ```yaml
  upstreamEndpoints:
    - "http://localhost:8081"
    - "http://localhost:8082"
  ```

##### Example Configuration

Load-Balanced Forwarding with `BalancerHttpAction`:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/httpfw.so
```

```yaml title="actions.yaml"
actions:
  # ...
  LoadBalancedForward:
    use: plugins.httpfw.actions.BalancerHttpAction
    config:
      upstreamEndpoints:
        - "http://localhost:8081"
        - "http://localhost:8082"
        - "http://localhost:8083"
```

```yaml title="routes.yaml"
routes:
  # ...
  - path: "/api/scale"
    method: "POST"
    action: LoadBalancedForward
```

### Best Practices

- **Health Checks**: Regularly monitor the health of upstream endpoints to ensure availability and minimize downtime.
- **Connection Pooling**: Optimize connections to upstream services to reduce latency and resource usage.
- **Combine with Middleware**: Use the HTTP Forward Plugin with authentication, rate limiting middlewares for a robust and secure routing solution.
