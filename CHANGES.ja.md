変更点
======

v1.0.7 (2022 年 01 月 22 日)
----------------------------

- `go.mod` 
    * `authlete-go => v1.1.10` 更新しました。
    * `toml => v1.2.1` 更新しました。
    * `gin => v1.8.2` 更新しました。
    * `locales => v0.14.1` 更新しました。
    * `go-json => v0.10.0` 更新しました。
    * `go-isatty => v0.0.17` 更新しました。
    * `codec => v1.2.8` 更新しました。
    * `crypto => v0.5.0` 更新しました。

v1.0.6 (2022 年 08 月 11 日)
----------------------------

- `TokenReqHandler` 構造体
    * `TokenAction_JWT_BEARER` をサポート。

- `TokenReqHandlerSpi` インターフェース
    * `JwtBearer(ctx *gin.Context, res *dto.TokenResponse)` メソッドを追加。

- `TokenReqHandlerSpiAdapter` 構造体
    * `JwtBearer(ctx *gin.Context, res *dto.TokenResponse)` メソッドを実装。

- `go.mod`
    * authlete-go のバージョンを v1.1.6 から v1.1.7 へ更新。

v1.0.5 (2022 年 08 月 09 日)
----------------------------

- `TokenReqHandler` 構造体
    * `TokenAction_TOKEN_EXCHANGE` をサポート。

- `TokenReqHandlerSpi` インターフェース
    * `TokenExchange(ctx *gin.Context, res *dto.TokenResponse)` メソッドを追加。

- `TokenReqHandlerSpiAdapter` 構造体
    * `TokenExchange(ctx *gin.Context, res *dto.TokenResponse)` メソッドを実装。

- `go.mod`
    * authlete-go のバージョンを v1.0.5 から v1.1.6 へ更新。

v1.0.4 (2019 年 09 月 25 日)
----------------------------

- このライブラリが参照する authlete-go のバージョンを v1.0.4 から v1.0.5 に変更。
- 空文字列を「値無し」とみなすよう `ClaimCollector` を変更。
- `no_interaction_handler.go` を追加。

v1.0.3 (2019 年 09 月 20 日)
----------------------------

- `ReqUtil` と `ResUtil` を `BaseEndpoint` に追加。

v1.0.2 (2019 年 09 月 19 日)
----------------------------

- `BaseEndpoint` の `api` を `Api` へと名称変更。

v1.0.1 (2019 年 09 月 19 日)
----------------------------

- このライブラリが参照する authlete-go のバージョンを v1.0.3 から v1.0.4 に変更。

v1.0.0 (2019 年 09 月 19 日)
----------------------------

- 最初のリリース
