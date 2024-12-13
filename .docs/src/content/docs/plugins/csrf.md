---
title: CSRF
description: A documentation page about the CSRF plugin.
tableOfContents: false
---

The CSRF plugin provides middleware to protect your application from Cross-Site Request Forgery (CSRF) attacks. It ensures that only requests with a valid CSRF token are processed, helping to secure sensitive actions.

### Middlewares

#### plugins.csrf.middlewares.CSRFMiddleware

- **Purpose**:  
  This middleware implements CSRF protection by verifying the presence of a valid CSRF token in incoming requests. It manages token generation, validation, and storage, integrating seamlessly with your routes.

- **Usage**:  
  When configured in a route's middleware list, the CSRFMiddleware ensures that requests include a valid CSRF token in the specified location (e.g., headers or cookies). Invalid requests are rejected with a `403 Forbidden` status.

- **Configuration Integration**:  
  The middleware is configured through the plugin and referenced by name in the route's middleware list.

##### Middleware Configuration

The CSRFMiddleware accepts the following configuration options:

- **KeyLookup**  
  Defines where to look for the CSRF token in the request. This can be a header, cookie, or query parameter.  
  Example: `"header:X-CSRF-Token"` or `"cookie:csrf_token"`.

- **CookieName**  
  The name of the cookie to store the CSRF token. Default: `"csrf_token"`.

- **CookieDomain**  
  Specifies the domain for the CSRF cookie. If not set, defaults to the request’s domain.

- **CookiePath**  
  Specifies the path for the CSRF cookie. Default: `"/"`.

- **CookieSecure**  
  Indicates whether the CSRF cookie should only be sent over HTTPS. Default: `false`.

- **CookieHTTPOnly**  
  Ensures the CSRF cookie is inaccessible to JavaScript (`HTTPOnly` flag). Default: `false`.

- **CookieSameSite**  
  Controls the SameSite attribute of the cookie. Accepted values: `"Strict"`, `"Lax"`, `"None"`. Default: `"Lax"`.

- **CookieSessionOnly**  
  If `true`, the CSRF cookie is a session cookie and does not have an expiration date. Default: `false`.

- **CookieExpiration**  
  Sets the expiration duration for the CSRF cookie. Default: `0`.

- **SingleUseToken**  
  If `true`, the CSRF token is valid for only one request and is refreshed afterward. Default: `false`.

- **Storage**  
  Specifies the storage backend for managing CSRF tokens.

##### Example Configuration

Below is an example of how to configure the CSRF plugin:

```yaml title="middlewares.yaml"
middlewares:
  # ...
  CSRFMiddleware:
    use: plugins.csrf.middlewares.CSRFMiddleware
    config:
      keyLookup: "header:X-CSRF-Token"
      cookieName: "csrf_token"
      cookieDomain: "example.com"
      cookiePath: "/"
      cookieSecure: true
      cookieHTTPOnly: true
      cookieSameSite: "Strict"
      cookieSessionOnly: false
      cookieExpiration: 86400 # 24 hours in seconds
      singleUseToken: true
      storage: "RedisStorage"
```

##### How It Works

1. **Token Generation**:  
   The middleware generates a CSRF token and stores it in the configured location (e.g., a cookie or a storage backend).

2. **Token Validation**:  
   For each incoming request, the middleware checks the presence and validity of the CSRF token in the specified location (e.g., headers or cookies).

3. **Single-Use Token Option**:  
   If `singleUseToken` is enabled, the token is invalidated after a single request, and a new one is generated for subsequent requests.

4. **Response**:
   - If the token is valid, the request proceeds to the next middleware or action.
   - If the token is invalid or missing, the middleware responds with:
     - HTTP status: `403 Forbidden`.

##### Using as Global Middleware

To use CSRFMiddleware as global middleware, configure the middleware file to include the middleware as global middleware:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/csrf.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  CSRFMiddleware:
    use: plugins.csrf.middlewares.CSRFMiddleware
    config:
      keyLookup: header:X-CSRF-Token
      cookieName: csrf_token
      cookieSecure: true
      cookieHTTPOnly: true
      singleUseToken: true
      storage: ValkeyCSRFStorage
  # ...
globalMiddlewares:
  - CSRFMiddleware
```

##### Using with Routes

To use CSRFMiddleware with a specific route, configure the route to include the middleware:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/csrf.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  CSRFMiddleware:
    use: plugins.csrf.middlewares.CSRFMiddleware
    config:
      keyLookup: header:X-CSRF-Token
      cookieName: csrf_token
      cookieSecure: true
      cookieHTTPOnly: true
      singleUseToken: true
      storage: ValkeyCSRFStorage
```

```yaml title="routes.yaml"
routes:
  # ...
  - path: "/form/submit"
    method: "POST"
    middlewares:
      - CSRFMiddleware
    action: HttpForwardToFormService
```

### Best Practices

- **Use Secure Cookies**: Set `cookieSecure: true` to ensure tokens are only sent over HTTPS.
- **Single-Use Tokens for Extra Security**: Enable `singleUseToken` to mitigate token reuse attacks.
- **Pair with Authentication**: Combine CSRF protection with user authentication mechanisms.
