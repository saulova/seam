---
title: Routes
description: A documentation page about routes.
tableOfContents: false
---

Routes are building blocks of the system, defining how incoming requests are processed and routed to the appropriate action. Each route specifies a path, HTTP methods, middlewares, and an action that determines the request's final destination or response.

### Structure of a Route

A route typically includes the following components:

- **Gateway Path**: The URL pattern that matches incoming requests (e.g., `/api/v1/resource`).
- **Methods**: The HTTP methods allowed for the route (e.g., `GET`, `POST`, `PUT`).
- **Middlewares**: Optional processing steps applied before the action.
- **Action**: The final operation that handles the request (e.g., forwarding to a backend, serving static files).

Example Route Configuration:

```yaml title="routes.yaml"
routes:
  # ...
  RouteName:
    gatewayPath: /api/v1/path
    methods: [GET, POST]
    middlewares:
      - MiddlewareName
    action: ActionName
```

### Route Components in Detail

#### 1. **Gateway Path**

Defines the URL pattern to match incoming requests. Paths can include static or dynamic segments:

- **Static Path**: Matches exact URLs.  
  Example: `/api/v1/users`
- **Dynamic Path**: Uses placeholders for variable segments.  
  Example: `/api/v1/users/:id` (matches `/api/v1/users/123`).
- **Wildcard Path**: Matching paths with variable length.
  Example: `/api/v1/*` (matches `/api/v1/users/123`).

Example:

```yaml title="routes.yaml"
routes:
  # ...
  RouteName:
    # ...
    gatewayPath: /api/v1/users
    # ...
```

#### 2. **Methods**

Specifies which HTTP methods are allowed.

Example:

```yaml title="routes.yaml"
routes:
  # ...
  RouteName:
    # ...
    methods: [GET, POST]
    # ...
```

#### 3. **Middlewares**

Defines processing steps applied to requests before the action is executed.

Example:

```yaml title="routes.yaml"
routes:
  # ...
  RouteName:
    # ...
    middlewares:
      - AuthMiddleware
    # ...
```

#### 4. **Action**

Determines what happens after all middlewares are processed. Actions define the route's final behavior, such as forwarding requests.

Example:

```yaml title="routes.yaml"
routes:
  # ...
  RouteName:
    # ...
    action: UserServiceForwardAction
```

### Best Practices

- **Meaningful Route Names**: Use clear and descriptive names for routes that reflect their purpose.
- **Organize Routes Logically**: Group routes by functionality (e.g., `/api/v1/users`, `/api/v1/products`).
- **Minimize Middleware Usage**: Use only necessary middlewares to maintain performance.
- **Use Specific Paths**: Avoid overly generic routes that might lead to unintended matches.
- **Document Each Route**: Provide clear descriptions of the purpose and expected behavior for every route.
