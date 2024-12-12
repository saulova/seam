---
title: Session
description: A documentation page about the session plugin.
tableOfContents: false
---

The Session Plugin provides middlewares for session management and propagation in web applications. It leverages a storage backend to manage session states and offers tools to map session data to request headers for downstream services.

### Middlewares  

#### plugins.session.middlewares.LoadSessionMiddleware

- **Purpose**:  
  This middleware loads session data from a storage backend and attaches it to the request context. It also manages session lifecycle attributes, such as cookies and session expiration.

- **Usage**:  
  - **Session Handling**: Automatically retrieves the session using a `session_id` (default lookup key) and injects the session data into the request context.
  - **Cookie Management**: Configures cookie properties such as `HTTPOnly`, `Secure`, `SameSite`, expiration, and session renewal.
  - **Auto-Renewal**: Automatically extends the session's expiration and changes session id if configured.

- **Configuration Integration**:  
  The middleware requires configuration to define the session storage backend and cookie attributes.

##### Middleware Configuration  

- **Storage**
  The name of the storage backend to use for session data.  
  Default: null (store in memory, not recommended for production)

- **KeyLookup**  
  Defines how the session key is extracted from the request.  
  Default: `"cookie:session_id"`.

- **CookieHTTPOnly**  
  Indicates whether the session cookie should be HTTP-only, preventing client-side access.  
  Default: `false`.  

- **CookieSecure**  
  Specifies whether the session cookie is restricted to HTTPS.  
  Default: `false`.  

- **CookieSameSite**  
  Controls the `SameSite` attribute for the session cookie.  
  Default: `"Lax"`.  

- **CookieSessionOnly**  
  Makes the session cookie valid only for the duration of the browser session.  
  Default: `false`.  

- **CookieExpiration**  
  The duration for which the session cookie remains valid.  
  Default: `2h`.  

- **DisableAutoRenew**  
  Disables automatic renewal of the session expiration.  
  Default: `false`.  

- **AutoRenewAfter**  
  The time after which the session is automatically renewed.  
  Default: `30m`.  

- **DisableSessionForward**  
  Prevents the session data from being forwarded in the request context.  
  Default: `false`.  

##### Example Configuration 

```yaml title="middlewares.yaml"
middlewares:
  LoadSessionMiddleware:
    use: plugins.session.middlewares.LoadSessionMiddleware
    config:
      storage: "ValkeySessionStorage"
      keyLookup: "cookie:session_id"
      cookieHTTPOnly: true
      cookieSecure: false
      cookieSameSite: "Lax"
      cookieExpiration: 2h
      disableAutoRenew: false
      autoRenewAfter: 30m
      disableSessionForward: false
```

#### plugins.session.middlewares.MapSessionToHeaderMiddleware

- **Purpose**:  
  This middleware maps session data to HTTP headers for propagation to downstream services, ensuring seamless integration with other systems.
- **Usage**:  
  - **Header Mapping**: Defines how session keys should be converted to headers.  
  - **Downstream Propagation**: Simplifies sharing session information with external services.  
- **Configuration Integration**:  
  The middleware is configured with a mapping of session keys to HTTP header names.

##### Middleware Configuration  

- **Headers**
  A key-value mapping where keys are the corresponding HTTP header names, and values are session attributes.  
  Example: `{ "Authorization": "authorization", "X-User-Id": "user_id" }`.  

##### Example Configuration  

```yaml title="middlewares.yaml"
middlewares:
  MapSessionToHeaderMiddleware:
    use: plugins.session.middlewares.MapSessionToHeaderMiddleware
    config:
      headers:
        Authorization: authorization
        X-User-Id: user_id
```

### Using with Routes 

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/session.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  LoadSessionMiddleware:
    use: plugins.session.middlewares.LoadSessionMiddleware
    config:
      cookieHTTPOnly: true
      cookieSecure: false
      cookieSameSite: "Lax"
      cookieExpiration: 2h
      autoRenewAfter: 30m
      disableSessionForward: false
  MapSessionToHeaderMiddleware:
    use: plugins.session.middlewares.MapSessionToHeaderMiddleware
    config:
      headers:
        Authorization: authorization
        X-User-Id: user_id
```

```yaml title="routes.yaml"
routes:
  # ...
  - path: "/dashboard"
    method: "GET"
    middlewares:
      - LoadSessionMiddleware
      - MapSessionToHeaderMiddleware
    action: HttpForwardToDashboardService
```

## Best Practices  

- **Use Secure Cookies**: Set `CookieSecure` to `true` in production to restrict cookies to HTTPS.
- **Auto-Renewal**: Use `AutoRenewAfter` to rotate the session id.
- **Minimize Header Propagation**: Only include essential session attributes in `MapSessionToHeaderMiddleware`.