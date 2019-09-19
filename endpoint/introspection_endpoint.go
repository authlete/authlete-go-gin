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

type IntrospectionEndpoint struct {
	BaseEndpoint
	Authenticate CallerAuthenticateFunc
	Reject       CallerRejectFunc
}

func (self *IntrospectionEndpoint) Handle(ctx *gin.Context) {
	// "1.1. Introspection Request" in RFC 7662 says as follows:
	//
	//   To prevent token scanning attacks, the endpoint MUST also require
	//   some form of authorization to access this endpoint, such as client
	//   authentication as described in OAuth 2.0 [RFC6749] or a separate
	//   OAuth 2.0 access token such as the bearer token described in OAuth
	//   2.0 Bearer Token Usage [RFC6750]. The methods of managing and
	//   validating these authentication credentials are out of scope of
	//   this specification.
	//
	// Therefore, this API must be protected in some way or other. Let's
	// perform authentication of the API caller.
	authenticated := self.Authenticate(ctx)

	// If the API caller does not have necessary privilages to call this API.
	if !authenticated {
		// Reject the API call.
		self.Reject(ctx)
		return
	}

	api := self.GetAuthleteApi(ctx)
	if api == nil {
		return
	}

	handler := handler.IntrospectionReqHandler_New(api)
	handler.Handle(ctx)
}

func IntrospectionEndpoint_Handler(
	authenticateFunc CallerAuthenticateFunc, rejectFunc CallerRejectFunc) gin.HandlerFunc {
	// Instance of introspection endpoint
	endpoint := IntrospectionEndpoint{}
	endpoint.Authenticate = authenticateFunc
	endpoint.Reject = rejectFunc

	return func(ctx *gin.Context) {
		endpoint.Handle(ctx)
	}
}
