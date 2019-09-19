Gin (Go) 用 Authlete ライブラリ
===============================

概要
----

このライブラリは、[OAuth 2.0][RFC6749] および [OpenID Connect][OIDC]
をサポートする認可サーバーと、
リソースサーバーを実装するためのユーティリティー部品群を提供します。

このライブラリは、Gin API と authlete-go ライブラリを用いて書かれています。
[Gin][Gin] は Go で書かれた Web フレームワークの一つです。
一方、[authlete-go][AuthleteGo] は Authlete
が提供するもう一つのオープンソースライブラリで、[Authlete Web API][AuthleteAPI]
とやりとりするための基本部品群を含んでいます。

[Authlete][Authlete] は OAuth 2.0 と OpenID Connect の実装を提供するクラウドサービスです
([overview][AuthleteOverview])。 認可データ (アクセストークン等)
や認可サーバー自体の設定、クライアントアプリケーション群の設定はクラウド上の Authlete
サーバーに保存されるため、Authlete を使うことで「DB レス」の認可サーバーを構築することができます。

[gin-oauth-server][GinOAuthServer] はこのライブラリを使用している認可サーバーの実装で、
認可エンドポイントやトークンエンドポイントに加え、JWK Set エンドポイント、
ディスカバリーエンドポイント、取り消しエンドポイントの実装を含んでいます。
また、[gin-resource-server][GinResourceServer]
はこのライブラリを使用しているリソースサーバーの実装です。 [OpenID Connect Core 1.0][OIDCCore]
で定義されている[ユーザー情報エンドポイント][UserInfoEndpoint]
をサポートし、また、保護リソースエンドポイントの例を含んでいます。
あなたの認可サーバーおよびリソースサーバーの実装の開始点として、
これらのサンプル実装を活用してください。

ライセンス
----------

  Apache License, Version 2.0

ソースコード
------------

  <code>https://github.com/authlete/authlete-go-gin</code>

パッケージ
----------

    import (
        "github.com/authlete/authlete-go-gin/endpoint"
        "github.com/authlete/authlete-go-gin/handler"
        "github.com/authlete/authlete-go-gin/handler/spi"
        "github.com/authlete/authlete-go-gin/middleware"
        "github.com/authlete/authlete-go-gin/web"
    )

サンプル
--------

#### ディスカバリー・エンドポイント

```go
package main

import (
    "github.com/authlete/authlete-go-gin/endpoint"
    "github.com/authlete/authlete-go-gin/middleware"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // `authlete.toml` から設定をロードし、api.AuthleteApi のインスタンスを
    // 生成し、`AuthleteApi` というキーでそのインスタンスを Gin コンテキストに
    // 設定するミドルウェアを登録する。
    r.Use(middleware.AuthleteApi_Toml(`authlete.toml`))

    // OpenID Connect Discovery 1.0 に準拠するディスカバリー・エンドポイントを
    // 定義する。
    r.GET("/.well-known/openid-configuration",
          endpoint.DiscoveryEndpoint_Handler())

    // http://localhost:8080 でこのサーバーを起動する。
    r.Run()
}
```

#### 保護リソースエンドポイント

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
    // リクエストに含まれているアクセストークンを検証する。
    valid, validator := self.ValidateAccessToken(ctx, nil)

    // アクセストークンが無効の場合
    if !valid {
        // RFC 6750 に準拠するエラー応答を生成する。
        validator.Deny(ctx)
        return
    }

    // このエンドポイントからの応答
    ctx.JSON(200, gin.H{"message":"hello"})
}

func HelloEndpoint_Handler() gin.HandlerFunc {
    // Hello エンドポイントのインスタンス
    endpoint := HelloEndpoint{}

    return func(ctx *gin.Context) {
        endpoint.Handle(ctx)
    }
}

func main() {
    r := gin.Default()

    // 環境から設定を読み込み、api.AuthleteApi のインスタンスを生成し、
    // `AuthleteApi` というキーでそのインスタンスを Gin コンテキストに
    // 設定するミドルウェアを登録する。
    r.Use(middleware.AuthleteApi_Env())

    // '/api/hello' API を定義する。
    r.GET("/api/hello", HelloEndpoint_Handler())

    // http://localhost:8080 でこのサーバーを起動する。
    r.Run()
}
```

コンタクト
----------

コンタクトフォーム : https://www.authlete.com/contact/

| 目的 | メールアドレス       |
|:-----|:---------------------|
| 一般 | info@authlete.com    |
| 営業 | sales@authlete.com   |
| 広報 | pr@authlete.com      |
| 技術 | support@authlete.com |

[Authlete]:          https://www.authlete.com/ja/
[AuthleteAPI]:       https://docs.authlete.com/
[AuthleteOverview]:  https://www.authlete.com/ja/developers/overview/
[AuthleteGo]:        https://github.com/authlete/authlete-go/
[Gin]:               https://github.com/gin-gonic/gin
[GinOAuthServer]:    https://github.com/authlete/gin-oauth-server/
[GinResourceServer]: https://github.com/authlete/gin-resource-server/
[OIDC]:              https://openid.net/connect/
[OIDCCore]:          https://openid.net/specs/openid-connect-core-1_0.html
[RFC6749]:           https://tools.ietf.org/html/rfc6749
[UserInfoEndpoint]:  https://openid.net/specs/openid-connect-core-1_0.html#UserInfo
