//
// Copyright (C) 2019 Authlete, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific
// language governing permissions and limitations under the
// License.

package handler

import (
	"github.com/authlete/authlete-go/api"
	"github.com/gin-gonic/gin"
)

type JwksReqHandler struct {
	BaseReqHandler
}

func JwksReqHandler_New(api api.AuthleteApi) *JwksReqHandler {
	handler := JwksReqHandler{}
	handler.Init(api)

	return &handler
}

func (self *JwksReqHandler) Handle(ctx *gin.Context, pretty bool) {
	// Call Authlete's /api/service/jwks/get API. The API returns the JWK Set
	// (RFC 7517) of the service. The second argument given to GetServiceJwks()
	// is false not to include private keys.
	jwks, err := self.Api.GetServiceJwks(pretty, false)

	if err != nil {
		if err.Response == nil || err.Response.StatusCode != 302 {
			// Something wrong happened.
			self.ResUtil.WithAuthleteError(ctx, err)
		} else {
			// 302 Found
			location := err.Response.Header.Get(`Location`)
			self.ResUtil.Location(ctx, location)
		}
		return
	}

	// If no JWK Set for the service is registered
	if jwks == `` {
		// 204 No Content
		self.ResUtil.NoContent(ctx)
	} else {
		// 200 OK, application/json;charset=UTF-8
		self.ResUtil.OkJson(ctx, jwks)
	}
}
