---
title: Services
description: A documentation page about services.
tableOfContents: false
---

Services represent logical groupings of functionality within the application. They act as abstractions for specific business domains or system components and define which routes belong to them. This organization helps manage and scale the system by clearly separating responsibilities.

Example Configuration:

```yaml title="routes.yaml"
services:
  # ...
  ServiceName:
    routes:
      - RouteName
```

### Best Practices

- **Meaningful Service Names**: Use clear and descriptive names for services that reflect their purpose.
- **Limit Service Scope**: Avoid assigning unrelated routes to the same service. Keep services focused and cohesive.
- **Reuse Routes**: If multiple services use the same route, ensure proper abstraction and avoid duplicating route definitions.
