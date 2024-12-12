---
title: Helmet
description: A documentation page about the Helmet plugin.
tableOfContents: false
---

The Helmet plugin provides middleware for setting HTTP security headers, enhancing your application’s protection against common web vulnerabilities. It allows fine-grained control over headers to secure your application effectively.

### Middlewares

#### plugins.helmet.middlewares.HelmetMiddleware

- **Purpose**:  
  This middleware configures various HTTP headers to improve security, such as mitigating XSS attacks, preventing content sniffing, and enforcing secure policies for content, permissions, and cross-origin interactions.

- **Usage**:  
  When configured in a route's middleware list, the HelmetMiddleware ensures the specified security headers are added to responses, improving security without modifying application logic.

- **Configuration Integration**:  
  The middleware is configured through the plugin and referenced by name in the route's middleware list.

##### Middleware Configuration

The HelmetMiddleware accepts the following configuration options:

- **XSSProtection**  
  Configures the `X-XSS-Protection` header to prevent reflected XSS attacks.  
  Default: `"0"` (disabled).

- **ContentTypeNosniff**  
  Sets the `X-Content-Type-Options` header to `nosniff`, preventing browsers from interpreting files as a different MIME type than declared. Default: `"nosniff"`.

- **XFrameOptions**  
  Controls the `X-Frame-Options` header to prevent clickjacking by disallowing the page from being embedded in an iframe. Default: `"SAMEORIGIN"`.

- **HSTSMaxAge**  
  Configures the `Strict-Transport-Security` (HSTS) header's `max-age` directive in seconds, enforcing HTTPS connections.

- **HSTSExcludeSubdomains**  
  Determines whether HSTS applies to subdomains. If `true`, excludes subdomains. Default: `false`.

- **ContentSecurityPolicy**  
  Sets the `Content-Security-Policy` (CSP) header to control sources of content.

- **CSPReportOnly**  
  Enables `Content-Security-Policy-Report-Only`, logging violations without blocking requests. Default: `false`.

- **HSTSPreloadEnabled**  
  Adds the `preload` directive to HSTS, allowing the domain to be included in browser preload lists. Default: `false`.

- **ReferrerPolicy**  
  Configures the `Referrer-Policy` header to control the information sent in the `Referer` header. Default: `"ReferrerPolicy"`.

- **PermissionPolicy**  
  Sets the `Permissions-Policy` header to control access to browser features.

- **CrossOriginEmbedderPolicy**  
  Sets the `Cross-Origin-Embedder-Policy` header. Default: `"require-corp"`.

- **CrossOriginOpenerPolicy**  
  Sets the `Cross-Origin-Opener-Policy` header. Default: `"same-origin"`.

- **CrossOriginResourcePolicy**  
  Sets the `Cross-Origin-Resource-Policy` header. Default: `"same-origin"`.

- **OriginAgentCluster**  
  Adds the `Origin-Agent-Cluster` header. Default: `"?1"`.

- **XDNSPrefetchControl**  
  Configures the `X-DNS-Prefetch-Control` header to enable or disable DNS prefetching. Default: `"off"`.

- **XDownloadOptions**  
  Sets the `X-Download-Options` header to prevent Internet Explorer from executing downloads in the site's context. Default: `"noopen"`.

- **XPermittedCrossDomain**  
  Sets the `X-Permitted-Cross-Domain-Policies` header to control permissible policies for cross-domain content. Default: `"none"`.

##### Example Configuration

Below is an example of how to configure the Helmet plugin:

```yaml title="middlewares.yaml"
middlewares:
  HelmetMiddleware:
    use: plugins.helmet.middlewares.HelmetMiddleware
    config:
      xssProtection: "1; mode=block"
      contentTypeNosniff: "nosniff"
      xFrameOptions: "SAMEORIGIN"
      hstsMaxAge: 31536000
      referrerPolicy: "no-referrer"
      crossOriginEmbedderPolicy: "require-corp"
      crossOriginOpenerPolicy: "same-origin"
      crossOriginResourcePolicy: "same-origin"
      xdnsPrefetchControl: "off"
      xDownloadOptions: "noopen"
      xPermittedCrossDomain: "none"
```

##### How It Works

1. **Header Injection**:  
   HelmetMiddleware adds security-related HTTP headers to outgoing responses based on the configuration.

2. **Policy Enforcement**:  
   These headers instruct browsers to apply security restrictions, such as preventing MIME type sniffing, blocking insecure content, and enforcing strict origin policies.

3. **Response**:
   - For correctly configured routes, headers are injected into every response.
   - If a policy is misconfigured, browser behavior may vary.

##### Using as Global Middleware

To use HelmetMiddleware as global middleware, configure the middleware file to include the middleware as global middleware:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/helmet.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  HelmetMiddleware:
    use: plugins.helmet.middlewares.HelmetMiddleware
    config:
      xssProtection: "1; mode=block"
      contentTypeNosniff: "nosniff"
      hstsMaxAge: 31536000
      xFrameOptions: "SAMEORIGIN"
  # ...
globalMiddlewares:
  - HelmetMiddleware
```

##### Using with Routes

To use HelmetMiddleware with a specific route, configure the route to include the middleware:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/helmet.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  HelmetMiddleware:
    use: plugins.helmet.middlewares.HelmetMiddleware
    config:
      xssProtection: "1; mode=block"
      contentTypeNosniff: "nosniff"
      hstsMaxAge: 31536000
      xFrameOptions: "SAMEORIGIN"
```

```yaml title="routes.yaml"
routes:
  # ...
  - path: "/form/submit"
    method: "POST"
    middlewares:
      - HelmetMiddleware
    action: HttpForwardToFormService
```

### Best Practices

- **Enable HSTS**: Use `hstsMaxAge` with a reasonable duration (e.g., 1 year) to enforce HTTPS connections.
- **Restrict CSP Sources**: Specify strict policies for `contentSecurityPolicy` to control content loading.
- **Regularly Review Policies**: Periodically audit your security headers to align with evolving best practices.
- **Test in Staging**: Validate the middleware in a staging environment to avoid unexpected behavior in production.
