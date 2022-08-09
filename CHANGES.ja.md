変更点
======

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
