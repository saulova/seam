---
title: Basic Auth
description: A documentation page about the basic auth plugin.
tableOfContents: false
---

The Basic Auth plugin provides a middleware for basic HTTP authentication. It protects routes by requiring clients to provide a valid username and password in the `Authorization` header of their requests.

### Middlewares

#### plugins.basicauth.middlewares.BasicAuthMiddleware

- **Purpose**:  
  This middleware is responsible for applying Basic HTTP Authentication to the routes where it is configured. It validates the `Authorization` header, verifies the provided credentials against the defined users, and injects the authenticated username and password into the request context.

- **Usage**:  
  When configured in a route's middleware list, it ensures that only authenticated users can access the route. If authentication fails, the middleware responds with a `401 Unauthorized` status and halts further request processing.

- **Configuration Integration**:  
  The middleware is configured through the plugin and referenced by name in the route's middleware list.

##### Middleware Configuration

The plugin uses the following configuration structure to set up the BasicAuthMiddleware:

- **Realm**  
  The authentication realm displayed in the `WWW-Authenticate` response header when authentication fails. Default: `"Restricted"`.

- **Users**  
  A mapping of usernames to passwords. These credentials are used to validate incoming requests.  
  Example:

  ```yaml title="middlewares.yaml"
  middlewares:
    # ...
    BasicAuthMiddleware:
      use: plugins.basicauth.middlewares.BasicAuthMiddleware
      config:
        # ...
        users:
          admin: "password123"
          user: "userpass"
  ```

- **ContextUsernameKey**  
  The key used to store the authenticated username in the request context (locals). Default: `"username"`.

- **ContextPasswordKey**  
  The key used to store the authenticated password in the request context (locals). Default: `"password"`.

##### Example Configuration

Below is an example of how to configure the BasicAuth plugin:

```yaml title="middlewares.yaml"
middlewares:
  # ...
  BasicAuthMiddlewareName:
    use: plugins.basicauth.middlewares.BasicAuthMiddleware
    config:
      realm: "Restricted"
      users:
        admin: "password123"
        user: "userpass"
      contextUsernameKey: "username"
      contextPasswordKey: "password"
```

##### How It Works

1. **Request Validation**: The middleware intercepts incoming requests to routes protected by Basic Auth. It checks the `Authorization` header for credentials.

1. **Authentication Challenge**: If no valid credentials are provided, the server responds with:

   - HTTP status: `401 Unauthorized`
   - Header:
     ```http
     WWW-Authenticate: Basic realm="Restricted"
     ```

1. **Context Injection**: Upon successful authentication, the middleware injects the username and password into the request context, making them accessible to subsequent middlewares or actions.

##### Using with Routes

To use BasicAuthMiddleware with a specific route, configure the route to include the middleware:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/basicauth.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  BasicAuthMiddleware:
    use: plugins.basicauth.middlewares.BasicAuthMiddleware
    config:
      realm: "Restricted"
      users:
        admin: "password123"
        user: "userpass"
```

```yaml title="routes.yaml"
routes:
  # ...
  - path: "/secure"
    method: "GET"
    middlewares:
      - BasicAuthMiddleware
    action: HttpForwardToInternalSecureService
```

### Best Practices

- **Limit User Access**: Only define users who need access to specific resources.
- **Combine with Other Security Layers**: Use BasicAuth as part of a broader security strategy, such as TLS for encrypted communication.
