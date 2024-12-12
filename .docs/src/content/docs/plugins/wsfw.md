---
title: Websocket Forward
description: A documentation page about the websocket plugin.
tableOfContents: false
---

The WebSocket Forward Plugin enables applications to handle WebSocket connections efficiently, supporting both middleware for connection validation and an action for routing WebSocket traffic to upstream services. It is ideal for real-time communication use cases such as chat applications, live updates, and collaborative tools.

### Middlewares

#### plugins.wsfw.middlewares.CheckWebSocketUpgradeMiddleware

- **Purpose**:  
  This middleware ensures that incoming requests are valid WebSocket upgrade requests. It intercepts requests and checks for the appropriate WebSocket headers before allowing the request to proceed.

- **Usage**:  
  Include this middleware in your route configuration to validate WebSocket upgrade requests. If a request is not a valid WebSocket upgrade, the middleware will block it and respond accordingly.

- **Configuration Integration**:  
  This middleware does not require specific configuration parameters but must be referenced in the route's middleware list.

##### How It Works

1. Validates the `Upgrade` header to ensure it contains the value `"websocket"`.
2. Verifies that the connection type supports WebSocket (`Connection: Upgrade` header).
3. Rejects non-WebSocket upgrade requests with an appropriate HTTP error response.

### Actions

#### plugins.wsfw.actions.WebSocketAction

- **Purpose**:  
  This action handles WebSocket connections by routing them to an upstream WebSocket endpoint. It supports features such as ping-pong heartbeats, inactivity detection, and upstream communication.

- **Usage**:  
  Attach this action to a route to enable WebSocket communication. It forwards the WebSocket connection to the specified upstream endpoint while maintaining heartbeat intervals and monitoring connection activity.

- **Configuration Integration**:  
  The action is configured through the plugin, requiring several parameters to define timeout behaviors and the upstream WebSocket endpoint.

##### Configuration Parameters

- **PongTimeout**  
  The maximum time to wait for a pong response to a ping.  
  Example: `30s`  
  Default: `0` (no timeout).

- **PingInterval**  
  The interval between sending ping messages to keep the connection alive.  
  Example: `10s`  
  Default: `0` (no pings).

- **MaxInactivity**  
  The maximum duration of inactivity before the connection is considered stale and is closed.  
  Example: `60s`  
  Default: `0` (no inactivity checks).

- **UpstreamEndpoint**  
  The URL of the upstream WebSocket server to which connections will be routed.  
  Example: `"ws://localhost:8080/ws"`  
  Default: None (must be configured).

### Example Configuration

##### Validating WebSocket Upgrades:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/wsfw.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  CheckWebSocketUpgrade:
    use: plugins.wsfw.middlewares.CheckWebSocketUpgradeMiddleware
```

```yaml title="actions.yaml"
actions:
  # ...
  ForwardWebSocket:
    use: plugins.wsfw.actions.WebSocketAction
    config:
      pongTimeout: 30s
      pingInterval: 10s
      maxInactivity: 60s
      upstreamEndpoint: "ws://localhost:8080/ws"
```

```yaml title="routes.yaml"
routes:
  # ...
  - path: "/ws/chat"
    method: "GET"
    middlewares:
      - CheckWebSocketUpgrade
    action: ForwardWebSocket
```

### Best Practices

- **Heartbeat Management**: Use `pongTimeout` and `pingInterval` configurations to keep WebSocket connections healthy and detect inactive peers.
- **Connection Limits**: Combine WebSocket routes with rate-limiting middleware to prevent resource exhaustion.
- **Secure Communication**: Use secure WebSocket endpoints (`wss://`) to ensure encrypted communication, especially in production environments.
