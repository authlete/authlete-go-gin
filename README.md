Authlete Library for Gin (Go)
=============================

Overview
--------

This library provides utility components to make it easy for developers to
implement an authorization server which supports [OAuth 2.0][RFC6749] and
[OpenID Connect][OIDC] and a resource server.

This library is written using Gin API and authlete-go library. [Gin][Gin] is
a web framework written in Go. On the other hand, [authlete-go][AuthleteGo]
is another Authlete's open source library which provides basic components to
communicate with [Authlete Web APIs][AuthleteAPI].

[Authlete][Authlete] is a cloud service that provides an implementation of
OAuth 2.0 & OpenID Connect ([overview][AuthleteOverview]). You can build a
_DB-less_ authorization server by using Authlete because authorization data
(e.g. access tokens), settings of authorization servers and settings of client
applications are stored in the Authlete server on cloud.

[gin-oauth-server][GinOAuthServer] is an authorization server implementation
which uses this library. It implements not only an authorization endpoint and
a token endpoint but also a JWK Set endpoint, a discovery endpoint, an
introspection endpoint and a revocation endpoint.
[gin-resource-server][GinResourceServer] is a resource server implementation
which also uses this library. It supports a [userinfo endpoint][UserInfoEndpoint]
defined in [OpenID Connect Core 1.0][OIDCCore] and includes an example of a
protected resource endpoint, too. Use these sample implementations as a
starting point of your own implementations of an authorization server and a
resource server.

License
-------

  Apache License, Version 2.0

Source Code
-----------

  <code>https://github.com/authlete/authlete-go-gin</code>

Packages
--------

    import (
        "github.com/authlete/authlete-go-gin/endpoint"
        "github.com/authlete/authlete-go-gin/handler"
        "github.com/authlete/authlete-go-gin/handler/spi"
        "github.com/authlete/authlete-go-gin/middleware"
        "github.com/authlete/authlete-go-gin/web"
    )

Samples
-------

#### Discovery Endpoint

```go
package main

import (
    "github.com/authlete/authlete-go-gin/endpoint"
    "github.com/authlete/authlete-go-gin/middleware"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Register middleware that loads settings from `authlete.toml`,
    // creates an instance of api.AuthleteApi and sets the instance
    // to the given gin context with the key `AuthleteApi`.
    r.Use(middleware.AuthleteApi_Toml(`authlete.toml`))

    // Define a discovery endpoint that conforms to OpenID Connect
    // Discovery 1.0.
    r.GET("/.well-known/openid-configuration",
          endpoint.DiscoveryEndpoint_Handler())

    // Start this server at http://localhost:8080.
    r.Run()
}
```

#### Protected Resource Endpoint

```go
package main

import (
    "github.com/authlete/authlete-go-gin/endpoint"
    "github.com/authlete/authlete-go-gin/middleware"
    "github.com/gin-gonic/gin"
)

type HelloEndpoint struct {
    endpoint.BaseEndpoint
}

func (self *HelloEndpoint) Handle(ctx *gin.Context) {
    // Validate the access token included in the request.
    valid, validator := self.ValidateAccessToken(ctx, nil)

    // If the access token is not valid.
    if !valid {
        // Generate an error response that conforms to RFC 6750.
        validator.Deny(ctx)
        return
    }

    // Response from this endpoint.
    ctx.JSON(200, gin.H{"message":"hello"})
}

func HelloEndpoint_Handler() gin.HandlerFunc {
    // Instance of hello endpoint
    endpoint := HelloEndpoint{}

    return func(ctx *gin.Context) {
        endpoint.Handle(ctx)
    }
}

func main() {
    r := gin.Default()

    // Register middleware that reads settings from the environment,
    // creates an instance of api.AuthleteApi and sets the instance
    // to the given gin context with the key `AuthleteApi`.
    r.Use(middleware.AuthleteApi_Env())

    // Define '/api/hello' API.
    r.GET("/api/hello", HelloEndpoint_Handler())

    // Start this server at http://localhost:8080.
    r.Run()
}
```

Contact
-------

Contact Form : https://www.authlete.com/contact/

| Purpose   | Email Address        |
|:----------|:---------------------|
| General   | info@authlete.com    |
| Sales     | sales@authlete.com   |
| PR        | pr@authlete.com      |
| Technical | support@authlete.com |

[Authlete]:          https://www.authlete.com/
[AuthleteAPI]:       https://docs.authlete.com/
[AuthleteOverview]:  https://www.authlete.com/developers/overview/
[AuthleteGo]:        https://github.com/authlete/authlete-go/
[Gin]:               https://github.com/gin-gonic/gin
[GinOAuthServer]:    https://github.com/authlete/gin-oauth-server/
[GinResourceServer]: https://github.com/authlete/gin-resource-server/
[OIDC]:              https://openid.net/connect/
[OIDCCore]:          https://openid.net/specs/openid-connect-core-1_0.html
[RFC6749]:           https://tools.ietf.org/html/rfc6749
[UserInfoEndpoint]:  https://openid.net/specs/openid-connect-core-1_0.html#UserInfo
