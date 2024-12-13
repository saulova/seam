---
title: Server
description: A documentation page about server.
tableOfContents: false
---

The Server defines the core settings for how the application listens for and handles incoming requests. It includes configurations for network, performance, security, and routing behavior.

### Configuration Fields

Here’s a breakdown of the server configuration options:

#### General Settings

- **Address**  
  The address and port where the server listens for requests. Default: `0.0.0.0:8090`.

- **TLS**  
  Enables TLS (HTTPS) for secure communication. Default: `false`.

- **CertFile**  
  Path to the TLS certificate file, required if `TLS` is `true`. Default: `/certs/server.cert`.

- **KeyFile**  
  Path to the TLS key file, required if `TLS` is `true`. Default: `/certs/key.cert`.

- **AppName**  
  The name of the application.

#### Health Check

- **DisableHealthCheck**  
  Disables built-in health check routes. Default: `false`.

- **HealthCheckLiveRoute**  
  Defines the route for the "liveness" health check. Default: `/health/live`.

- **HealthCheckReadyRoute**  
  Defines the route for the "readiness" health check. Default: `/health/ready`.

#### Performance and Optimization

- **Prefork**  
  Enables preforking, where the server forks multiple OS processes for better multi-core CPU utilization. Default: `false`.

- **Concurrency**  
  Sets the maximum number of concurrent requests handled by the server. Default: `252144`.

- **BodyLimit**  
  Maximum size (in bytes) for request bodies. Prevents excessive memory usage from large payloads. Default: `4194304`.

- **ReduceMemoryUsage**  
  Optimizes internal memory usage at the cost of additional CPU overhead. Default: `false`.

- **StreamRequestBody**  
  Streams request bodies directly to handlers instead of loading them entirely into memory. Default: `false`.

- **CompressedFileSuffix**  
  File suffix used for serving pre-compressed static files. Default: `.seam.gz`.

#### Routing Behavior

- **StrictRouting**  
  Enables strict routing; distinguishes between `/route` and `/route/`. Default: `false`.

- **CaseSensitive**  
  Enables case-sensitive routing; distinguishes between `/Route` and `/route`. Default: `false`.

- **UnescapePath**  
  Automatically unescapes URL paths. Default: `.seam.gz`. Default: `false`.

- **GETOnly**  
  Restricts the server to handle only `GET` requests. Default: `false`.

#### Security

- **EnableTrustedProxyCheck**  
  Validates if incoming requests are from trusted proxies. Default: `false`.

- **TrustedProxies**  
  A list of IPs or CIDR blocks considered as trusted proxies. Default: `[]`.

- **EnableIPValidation**  
  Enables validation of client IPs in requests. Default: `false`.

#### Timeout Settings

- **ReadTimeout**  
  Maximum duration for reading the entire request. Default: `0s`.

- **WriteTimeout**  
  Maximum duration for writing the response. Default: `0s`.

- **IdleTimeout**  
  Maximum duration for keeping an idle connection open. Default: `0s`.

#### Headers and Defaults

- **ServerHeader**  
  Custom value for the `Server` HTTP header.

- **ETag**  
  Automatically adds ETag headers to responses for caching purposes. Default: `false`.

- **DisableDefaultDate**  
  Removes the default `Date` header from responses. Default: `false`.

- **DisableDefaultContentType**  
  Removes the default `Content-Type` header from responses. Default: `false`.

- **DisableHeaderNormalizing**  
  Prevents automatic normalization of header names to canonical form. Default: `false`.

#### Debugging and Logging

- **DisableStartupMessage**  
  Suppresses the startup message in logs. Default: `false`.

- **EnablePrintRoutes**  
  Prints all routes and their configurations during server startup. Default: `false`.

#### Buffer Settings

- **ReadBufferSize**  
  Sets the size of the read buffer for incoming requests. Default: `4096`.

- **WriteBufferSize**  
  Sets the size of the write buffer for outgoing responses. Default: `4096`.

#### Miscellaneous

- **DisableKeepalive**  
  Disables HTTP keep-alive connections. Default: `false`.

- **DisablePreParseMultipartForm**  
  Prevents automatic parsing of multipart form data. Default: `false`.

- **EnableSplittingOnParsers**  
  Enables splitting of input on request parsers for special cases. Default: `false`.

- **Network**  
  Defines the network type (e.g., `tcp` or `udp`) used by the server. Default: `tcp4`.

### Minimal Configuration

HTTP:

```yaml title="server.yaml"
server:
  address: "0.0.0.0:8090" # or 80
  prefork: true # production
```

HTTPS:

```yaml title="server.yaml"
server:
  address: "0.0.0.0:4433" # or 443
  tls: true
  certFile: "./certs/server.crt"
  keyFile: "./certs/server.key"
  prefork: true # production
```

### Best Practices

- **Use TLS**: Always enable TLS in production for secure communication.
- **Set Limits**: Define reasonable values for `BodyLimit`, `ReadTimeout`, and `WriteTimeout` to prevent resource exhaustion.
- **Enable Health Checks**: Use `HealthCheckLiveRoute` and `HealthCheckReadyRoute` for better monitoring and fault detection.
- **Optimize Performance**: Use `Prefork` for multi-core CPUs and tune `Concurrency` based on your workload.
- **Harden Security**: Validate trusted proxies and IPs if your application is behind a reverse proxy.
