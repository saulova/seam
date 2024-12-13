---
title: CORS
description: A documentation page about the CORS plugin.
tableOfContents: false
---

The CORS plugin provides a middleware for handling Cross-Origin Resource Sharing (CORS) in your application. It enables or restricts access to resources from different origins, supporting modern web applications cross-origin requests requirements.

### Middlewares

#### plugins.cors.middlewares.CORSMiddleware

- **Purpose**:  
  This middleware handles CORS policies by setting appropriate HTTP response headers based on the configuration. It allows you to control which origins, methods, and headers are permitted for cross-origin requests.

- **Usage**:  
  When configured in a route's middleware list, it ensures the appropriate CORS headers are added to responses for cross-origin requests. It can also handle preflight requests (`OPTIONS`) based on the specified policies.

- **Configuration Integration**:  
  The middleware is configured through the plugin and referenced by name in the route's middleware list.

##### Middleware Configuration

The middleware uses the following configuration structure:

- **AllowOrigins**  
  A comma-separated list of allowed origins. Use `*` to allow all origins. Default: `false`.
  Example: `"https://example.com, https://anotherdomain.com"`.

- **AllowMethods**  
  A comma-separated list of allowed HTTP methods.  
  Example: `"GET, POST, PUT, DELETE"`.

- **AllowHeaders**  
  A comma-separated list of allowed HTTP headers in the request.  
  Example: `"Content-Type, Authorization"`.

- **AllowCredentials**  
  A boolean value indicating whether credentials (e.g., cookies, HTTP authentication) are allowed in cross-origin requests. Default: `false`.

- **ExposeHeaders**  
  A comma-separated list of response headers that are exposed to the browser.  
  Example: `"X-Custom-Header, Content-Length"`.

- **MaxAge**  
  An integer specifying the maximum time (in seconds) that the browser can cache the preflight response.

##### Example Configuration

Below is an example of how to configure the CORS plugin:

```yaml title="middlewares.yaml"
middlewares:
  # ...
  CORSMiddleware:
    use: plugins.cors.middlewares.CORSMiddleware
    config:
      allowOrigins: "https://example.com, https://anotherdomain.com"
      allowMethods: "GET, POST, PUT, DELETE"
      allowHeaders: "Content-Type, Authorization"
      allowCredentials: true
      exposeHeaders: "X-Custom-Header, Content-Length"
      maxAge: 1m
```

##### How It Works

1. **Request Inspection**:  
   The middleware inspects incoming requests and evaluates their origin, method, and headers against the configured policies.

2. **Header Injection**:  
   If the request complies with the policies, the middleware injects CORS-related headers into the response.

3. **Preflight Handling**:  
   For preflight (`OPTIONS`) requests, the middleware evaluates the headers and responds directly with the configured policies without forwarding the request to subsequent middlewares or actions.

##### Using as Global Middleware

To use CORSMiddleware as global middleware, configure the middleware file to include the middleware as global middleware:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/cors.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  CORSMiddleware:
    use: plugins.cors.middlewares.CORSMiddleware
    config:
      allowOrigins: "https://example.com, https://anotherdomain.com"
      allowMethods: "GET, POST"
      allowHeaders: "Content-Type, Authorization"
      allowCredentials: false
      exposeHeaders: "X-Custom-Header"
      maxAge: 1800
  # ...
globalMiddlewares:
  - CORSMiddleware
```

##### Using with Routes

To use CORSMiddleware with a specific route, configure the route to include the middleware:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/cors.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  CORSMiddleware:
    use: plugins.cors.middlewares.CORSMiddleware
    config:
      allowOrigins: "https://example.com, https://anotherdomain.com"
      allowMethods: "GET, POST"
      allowHeaders: "Content-Type, Authorization"
      allowCredentials: false
      exposeHeaders: "X-Custom-Header"
      maxAge: 1800
```

```yaml title="routes.yaml"
routes:
  # ...
  - path: "/api/resource"
    method: "GET"
    middlewares:
      - CORSMiddleware
    action: HttpForwardToAPIService
```

### Best Practices

- **Restrict Origins**: Avoid using `*` for `allowOrigins` in production environments. Instead, define specific origins to minimize security risks.
- **Limit Methods and Headers**: Only allow the methods and headers necessary for your application.
- **Enable Credentials Carefully**: Use `allowCredentials: true` only when required, as it allows cookies and authentication headers in cross-origin requests.
- **Cache Preflight Responses**: Use a `maxAge` value to reduce preflight request overhead for frequently accessed routes.
