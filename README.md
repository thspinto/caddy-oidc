# caddy-oidc
[OpenID Connect](https://openid.net/connect/) plugin for Caddy Server.

This middleware providades authentication against an OIDC provider.

Features:

- Custom Scopes
- JWT token in cookies 
- Not before (nbf): forces reissuence of JWT before date
- Validate claims

## Caddyfile Syntax

## Configuration Options

| Parameter                   | Type        | Default  | Description                               |
|-----------------------------|-------------|----------|-------------------------------------------|
| cookie-domain               | string      |          | Optional cookie domain parameter          |
| cookie-http-only	          | boolean     | true     | Set cookie HTTP only flag                 |
| cookie-name	              | string      |          | Optional cookie domain parameter          |
| cookie-domain               | string      |          | Optional cookie domain parameter          |
| issuer                      | string      |          | OIDC issuer URL                           |
| client-id                   | string      |          | OIDC client id                            |
| client-secret               | string      |          | OIDC client id                            |
| redirect-url                | string      |          | Authentication callback URL               |
| scopes                      | string      |          | Optional Requested scopes                 |

## Setup
