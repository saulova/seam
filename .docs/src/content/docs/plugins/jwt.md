---
title: JWT
description: A documentation page about the JWT plugin.
tableOfContents: false
---

The JWT Plugin provides middleware for managing JSON Web Tokens (JWT) in your application. It validates and decodes tokens from incoming requests, ensuring secure access to protected routes.

### Middlewares

#### plugins.jwt.middlewares.JWTValidationMiddleware

- **Purpose**: The `JWTValidationMiddleware` validates the presence and authenticity of a JWT in the request. It ensures that only requests with valid tokens can proceed to protected routes or actions.

##### Middleware Configuration

The JWT Plugin uses the following configuration structure:

- **SigningKey**: Specifies the key used to validate the JWT signature.

  - **Algorithm**  
    The signing algorithm used for token verification. Supported algorithms include `"HS256"`, `"RS256"`, etc.  
    Example: `"HS256"`.

  - **Key**  
    The signing key or secret used for verifying JWTs. This key should match the one used to sign the tokens.  
    Example: `"my-secret-key"`.

- **JwkUrls**: A list of URLs pointing to JSON Web Key (JWK) sets. These URLs provide public keys for validating JWTs signed using asymmetric algorithms (e.g., `RS256`). The middleware will fetch the keys dynamically and use them for verification.

### Example Configuration

Below is an example configuration for the `JWTValidationMiddleware`:

```yaml title="middlewares.yaml"
middlewares:
  JWTValidationMiddleware:
    use: plugins.jwt.middlewares.JWTValidationMiddleware
    config:
      signingKey:
        algorithm: "HS256"
        key: "my-secret-key"
```

or

```yaml title="middlewares.yaml"
middlewares:
  JWTValidationMiddleware:
    use: plugins.jwt.middlewares.JWTValidationMiddleware
    config:
      jwkUrls:
        - "https://example.com/.well-known/jwks.json"
```

### How It Works

1. **Token Extraction**:  
   The middleware extracts the JWT from the `Authorization` header in the incoming request. The expected format is:

   ```http
   Authorization: Bearer <token>
   ```

2. **Validation**:

   - Verifies the token signature using the provided `SigningKey` or public keys from `JwkUrls`.
   - Ensures the token is not expired and validates its claims (e.g., audience, issuer) if applicable.

3. **Context Injection**:  
   On successful validation, the decoded token payload is injected into the request context, making it accessible to subsequent middlewares or actions.

4. **Error Handling**:  
   If validation fails, the middleware responds with:
   - HTTP status: `401 Unauthorized`
   - Body: An error message describing the issue.

### Using with Routes

To use the `JWTValidationMiddleware` with a route, configure the middleware and reference it in the route definition:

```yaml title="plugins.yaml"
plugins:
  # ...
  - path: /plugins/jwt.so
```

```yaml title="middlewares.yaml"
middlewares:
  # ...
  JWTValidationMiddleware:
    use: plugins.jwt.middlewares.JWTValidationMiddleware
    config:
      signingKey:
        algorithm: "HS256"
        key: "my-secret-key"
```

```yaml title="routes.yaml"
routes:
  # ...
  - path: "/protected"
    method: "GET"
    middlewares:
      - JWTValidationMiddleware
    action: HttpForwardToSecureService
```

### Best Practices

- **Secure Key Management**: Ensure the signing key is stored securely and not exposed in the codebase.
- **Use HTTPS**: Always use HTTPS to protect the `Authorization` header from being intercepted.
- **Rotate Keys**: Regularly rotate signing keys or update JWK URLs to maintain a strong security posture.
- **Validate Claims**: Ensure the token contains and validates required claims (e.g., audience, issuer, etc.) based on your application's needs.
