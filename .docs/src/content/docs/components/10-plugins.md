---
title: Plugins
description: A documentation page about plugins.
tableOfContents: false
---

Plugins allow you to extend the functionality of the system by dynamically adding new middlewares, actions, and storages. Plugins are modular components that can be loaded at runtime, making it easy to customize or enhance the system without modifying its core codebase.

### Plugin Structure

A plugin consists of:

1. **Path**: The file path to the plugin's shared object (`.so` file).
2. **Configuration**: Specific settings required for the plugin's functionality. This is passed as a key-value structure to the plugin.

### Example Configuration

Below is an example of how to configure and load a plugin for session management:

```yaml title="plugins.yaml"
plugins:
  - path: /plugins/valkey.so
  - path: /plugins/session.so
    config:
      storage: ValkeySessionStorage
```

:::caution
Plugins are imported in the order they are defined, from top to bottom. For example, if session depends on valkey, the valkey plugin must be imported first.
:::

### How Plugins Work

1. **Loading**: The system dynamically loads the plugin from the specified `.so` file path.
2. **Initialization**: The plugin initializes its functionality based on the provided configuration.
3. **Integration**:
   - Adds new **middlewares**, **actions**, or **storages** to the system.
   - Configures routes or other components if needed.

### Plugin Development

To create a plugin:

- **Implement the Plugin Interface**: Define the logic for new middlewares, actions, or storages (check the plugin.go inside a default plugin folder).
- **Compile the Plugin**: Use the Go `plugin` package to compile the code into a `.so` file.
- **Deploy the Plugin**: Place the `.so` file in the designated directory and reference it in the configuration.

### Best Practices

- **Isolate Plugin Logic**: Keep plugins modular and focused on a specific functionality to ensure maintainability.
- **Manage Dependencies**: Ensure plugins have access to required dependencies (e.g., storages) as part of their configuration.
- **Monitor Performance**: Test plugins for performance impact, especially if they add middlewares or actions in critical request paths.
