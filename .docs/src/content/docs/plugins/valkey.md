---
title: Valkey
description: A documentation page about the Valkey plugin.
tableOfContents: false
---

Valkey is a storage plugin designed to manage key-value data efficiently. It provides a configuration interface for connecting to a Valkey server and managing data storage operations. This storage can be used by various middlewares and actions within the system.

### Middlewares

#### plugins.valkey.storages.ValkeyStorage

- **Purpose**:  
  The **Valkey Storage** plugin provides a robust key-value storage solution designed for high-performance and scalable applications. Its primary purpose is to support middlewares and actions by offering a centralized and consistent storage backend. This makes it ideal for use cases such as session management, caching, and other stateful operations within the system.

- **Usage**:  
  The Valkey Storage plugin is referenced in the configurations of middlewares and actions that require persistent or temporary storage. By configuring Valkey, these components can store, retrieve, and manage data efficiently without directly interacting with the underlying storage implementation.

  For instance:

  - **Middlewares** like session management can use Valkey to persist user sessions.
  - **Actions** and **Plugins** such as caching or state handling can leverage Valkey for quick data access and updates.

  To utilize Valkey Storage, define it in the `storages` section of your configuration file and refer to it by name in the middleware or action configuration.

- **Configuration Integration**:  
  Valkey Storage integrates seamlessly into the system's configuration hierarchy. The integration involves the following steps:

  1. **Define the Storage**:  
     Add an entry in the `storages` section of your configuration file. This entry specifies the Valkey server connection details and optional key prefix.

  2. **Reference the Storage**:  
     Use the defined storage name in middleware, action or plugin configurations that require it. This ensures that the components can interact with Valkey as their storage backend.

##### Storage Configuration

The following options are available for configuring the Valkey Storage:

- **Prefix**
  A string prefix applied to all keys stored in the Valkey instance. This helps in organizing or segregating data by context. Default: `"go:seam:"`.

- **Host**
  The hostname or IP address of the Valkey server. Default: `""` (must be configured).

- **Port**
  The port number on which the Valkey server is running. Default: `0` (must be configured).

##### Example Configuration

Below is an example of how to configure Valkey Storage in your application:

```yaml title="storages.yaml"
storages:
  ValkeySessionStorage:
    use: plugins.valkey.storages.ValkeyStorage
    config:
      prefix: "session_"
      host: "valkey"
      port: 6379
```

##### How It Works

1. **Key Prefixing**:  
   All keys sent to the Valkey server are prefixed with the configured `prefix`. This prevents key collisions when multiple applications or modules use the same storage.

2. **Server Communication**:  
   The plugin connects to the Valkey server using the specified `host` and `port`. It then facilitates read and write operations for middlewares or actions that rely on the storage.

3. **Integration with Middlewares and Actions**:  
   Valkey can serve as the backend for session management, caching, or other stateful middlewares and actions. Middleware or action configurations reference the storage by name.

##### Using Valkey with a Middleware

To enable Valkey storage for a middleware, configure it as follows:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/valkey.so
```

```yaml title="storages.yaml"
storages:
  # ...
  ValkeySessionStorage:
    use: plugins.valkey.storages.ValkeyStorage
    config:
      prefix: "session_"
      host: "valkey"
      port: 6379
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  SessionMiddleware:
    use: plugins.sessions.middlewares.SessionMiddleware
    config:
      storage: ValkeySessionStorage
      # ...
```

### Best Practices

- **Define Unique Prefixes**: Use a distinct prefix for each storage use case (e.g., sessions, caching) to avoid unintended key collisions.
- **Monitor the Valkey Server**: Ensure the Valkey server is healthy and scaled to handle your application’s load, especially if it is shared among multiple components.
- **Secure Access**: Protect the Valkey server with proper network configurations and authentication (if supported).
