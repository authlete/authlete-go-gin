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

type DiscoveryReqHandler struct {
	BaseReqHandler
}

func DiscoveryReqHandler_New(api api.AuthleteApi) *DiscoveryReqHandler {
	handler := DiscoveryReqHandler{}
	handler.Init(api)

	return &handler
}

func (self *DiscoveryReqHandler) Handle(ctx *gin.Context, pretty bool) {
	// Call Authlete's /api/service/configuration API. The API returns
	// JSON that conforms to OpenID Connect Discovery 1.0.
	jsn, err := self.Api.GetServiceConfiguration(pretty)

	// If the API call failed.
	if err != nil {
		self.ResUtil.WithAuthleteError(ctx, err)
		return
	}

	// 200 OK, application/json;charset=UTF-8
	self.ResUtil.OkJson(ctx, jsn)
}
