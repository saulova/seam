---
title: Actions
description: A documentation page about actions.
tableOfContents: false
---

Actions components defines the final operation executed within a route. It determines what happens to a request once all middlewares and route-specific processing are completed. Actions defining the endpoint's behavior and its interaction with external resources.

### HTTP Forward Action

The HTTP Forward action routes the incoming request to a specified external HTTP service. It acts as a proxy, sending the request to the target URL and returning the service's response to the client.

**Use Case**: Integrating with backend services or APIs.

Example Configuration:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/httpfw.so
```

```yaml title="actions.yaml"
actions:
  # ...
  ActionName:
    use: plugins.httpfw.actions.HttpAction
    config:
      upstreamEndpoint: http://backend-service:8080
```

### Balancer Forward Action

The Balancer Forward action is used to distribute requests across multiple backend services based on a round-robin load balancing strategy.

**Use Case**: Ensuring high availability and load distribution across services.

Example Configuration:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/httpfw.so
```

```yaml title="actions.yaml"
action:
  # ...
  ActionName:
    use: plugins.httpfw.actions.BalancerHttpAction
    config:
      upstreamEndpoints:
        - http://backend-service-1:8080
        - http://backend-service-2:8080
        - http://backend-service-3:8080
```

### How Actions Work

- **Sequential Flow**: Actions are triggered only after all route-specific middlewares and preprocessing are complete.
- **Single Responsibility**: Each route can have only one action. For complex scenarios, delegate tasks to external services or use middleware.

### Best Practices

- **Meaningful Action Names**: Use clear and descriptive names for actions that reflect their purpose.
