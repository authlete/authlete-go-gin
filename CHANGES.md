CHANGES
=======

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
