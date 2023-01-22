CHANGES
=======

v1.0.7 (2022 年 01 月 22 日)
----------------------------

- `go.mod` 
    * Updated `authlete-go => v1.1.10`.
    * Updated `toml => v1.2.1`.
    * Updated `gin => v1.8.2`.
    * Updated `locales => v0.14.1`.
    * Updated `go-json => v0.10.0`.
    * Updated `go-isatty => v0.0.17`.
    * Updated `codec => v1.2.8`.
    * Updated `crypto => v0.5.0`.

v1.0.6 (2022-08-11)
-------------------

- `TokenReqHandler` struct
    * Supported `TokenAction_JWT_BEARER`.

- `TokenReqHandlerSpi` interface
    * Added `JwtBearer(ctx *gin.Context, res *dto.TokenResponse)` method.

- `TokenReqHandlerSpiAdapter` struct
    * Implemented `JwtBearer(ctx *gin.Context, res *dto.TokenResponse)` method.

- `go.mod`
    * Updated the version of authlete-go from v1.1.6 to v1.1.7.

v1.0.5 (2022-08-09)
-------------------

- `TokenReqHandler` struct
    * Supported `TokenAction_TOKEN_EXCHANGE`.

- `TokenReqHandlerSpi` interface
    * Added `TokenExchange(ctx *gin.Context, res *dto.TokenResponse)` method.

- `TokenReqHandlerSpiAdapter` struct
    * Implemented `TokenExchange(ctx *gin.Context, res *dto.TokenResponse)` method.

- `go.mod`
    * Updated the version of authlete-go from v1.0.5 to v1.1.6.

v1.0.4 (2019-09-25)
-------------------

- Updated the version of authlete-go this library refers to from v1.0.4 to v1.0.5.
- Modified `ClaimCollector` to regard empty strings as "unavailable".
- Added `no_interaction_handler.go`.

v1.0.3 (2019-09-20)
-------------------

- Added `ReqUtil` and `ResUtil` to `BaseEndpoint`.

v1.0.2 (2019-09-19)
-------------------

- Renamed `api` in `BaseEndpoint` to `Api`.

v1.0.1 (2019-09-19)
-------------------

- Updated the version of authlete-go this library refers to from v1.0.3 to v1.0.4.

v1.0.0 (2019-09-19)
-------------------

- First release
