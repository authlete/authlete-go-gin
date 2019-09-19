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

package endpoint

import (
	"github.com/authlete/authlete-go-gin/handler"
	"github.com/gin-gonic/gin"
)

type JwksEndpoint struct {
	BaseEndpoint
}

func (self *JwksEndpoint) Handle(ctx *gin.Context) {
	api := self.GetAuthleteApi(ctx)
	if api == nil {
		return
	}

	handler := handler.JwksReqHandler_New(api)
	handler.Handle(ctx, true)
}

func JwksEndpoint_Handler() gin.HandlerFunc {
	// Instance of JWK Set endpoint
	endpoint := JwksEndpoint{}

	return func(ctx *gin.Context) {
		endpoint.Handle(ctx)
	}
}
