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
	"github.com/authlete/authlete-go/dto"
	"github.com/gin-gonic/gin"
)

type IntrospectionReqHandler struct {
	BaseReqHandler
}

func IntrospectionReqHandler_New(api api.AuthleteApi) *IntrospectionReqHandler {
	handler := IntrospectionReqHandler{}
	handler.Init(api)

	return &handler
}

func (self *IntrospectionReqHandler) Handle(ctx *gin.Context) {
	// Request parameters in the request body.
	params := self.ReqUtil.ExtractForm(ctx)

	// Call Authlete's /api/auth/introspection/standard API.
	res, err := self.callStandardIntrospectionApi(ctx, params)
	if err != nil {
		return
	}

	// 'action' in the response denotes the next action which the
	// implementation of the introspection endpoint should take.
	action := res.Action

	// The content of the response to the API caller.
	content := res.ResponseContent

	switch action {
	case dto.StandardIntrospectionAction_INTERNAL_SERVER_ERROR:
		// 500 Internal Server Error
		self.ResUtil.InternalServerError(ctx, content)
	case dto.StandardIntrospectionAction_BAD_REQUEST:
		// 400 Bad Request
		self.ResUtil.BadRequest(ctx, content)
	case dto.StandardIntrospectionAction_OK:
		// 200 OK
		self.ResUtil.OkJson(ctx, content)
	default:
		// 500 Internal Server Error
		// The /api/auth/introspection/standard API returned an unknown action.
		self.UnknownAction(ctx, `/api/auth/introspection/standard`)
	}
}

func (self *IntrospectionReqHandler) callStandardIntrospectionApi(
	ctx *gin.Context, params string) (
	res *dto.StandardIntrospectionResponse, err *api.AuthleteError) {
	// Create a request for /api/auth/introspection/standard API.
	req := dto.StandardIntrospectionRequest{}
	req.Parameters = params

	// Call /api/auth/introspection/standard API.
	res, err = self.Api.StandardIntrospection(&req)

	if err != nil {
		self.ResUtil.WithAuthleteError(ctx, err)
	}

	return
}
