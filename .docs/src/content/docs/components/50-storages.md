---
title: Storages
description: A documentation page about storages.
tableOfContents: false
---

Storages are external storage mechanisms used by plugins, middlewares and actions to store, retrieve, or manage data. They enable functionalities such as session management, caching, or state persistence, providing a flexible way to integrate storage solutions into the system.

Example Configuration:

```yaml title="plugins.yaml"
plugins:
  - path: /plugins/valkey.so
```

```yaml title="storages.yaml"
storages:
  ValkeyStorageName:
    use: plugins.valkey.storages.ValkeyStorage
    config:
      prefix: cool-prefix
      host: valkey-host
      port: 6379
```

### How Storages Are Used

Storages are referenced by middlewares or actions that require persistent or temporary data handling. For example:

- **Session Middleware**: Stores user session data in a specified storage (e.g., Redis).
- **Rate Limiting Middleware**: Uses storage to track request counts for each client.
- **Caching**: Stores frequently accessed data to reduce load on backend services.

### Use Case

#### Session Management

Middlewares can use a storage to load user sessions. For example:

```yaml title="plugins.yaml"
plugins:
  - path: /plugins/valkey.so
  - path: /plugins/session.so
    config:
      storage: ValkeySessionStorage
      cookieHTTPOnly: true
      cookieSecure: true
      cookieSameSite: true
      cookieSessionOnly: true
      cookieExpiration: 2h
```

```yaml title="storages.yaml"
storages:
  ValkeySessionStorage:
    use: plugins.valkey.storages.ValkeyStorage
    config:
      prefix: session
      host: valkey
      port: 6379
```

```yaml title="middlewares.yaml"
middlewares:
  SessionMiddleware:
    use: plugins.session.middlewares.LoadSessionMiddleware
```

This configuration tells the session plugin to use `ValkeySessionStorage` for storing session data.

### Best Practices

- **Meaningful Storage Names**: Use clear and descriptive names for storage that reflect their purpose.
- **Use Dedicated Storages**: Assign separate storages for distinct purposes (e.g., one for sessions, another for caching) to avoid data conflicts.
- **Scalability**: Choose storages that can scale with your application's needs, especially in distributed environments.
