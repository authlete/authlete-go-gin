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

type RevocationReqHandler struct {
	BaseReqHandler
}

func RevocationReqHandler_New(api api.AuthleteApi) *RevocationReqHandler {
	handler := RevocationReqHandler{}
	handler.Init(api)

	return &handler
}

func (self *RevocationReqHandler) Handle(ctx *gin.Context) {
	// Message body of the request.
	params := self.ReqUtil.ExtractForm(ctx)

	// Basic authentication of the request.
	user, pass, _ := ctx.Request.BasicAuth()

	// Call Authlete's /api/auth/revocation API.
	res, err := self.callRevocationApi(ctx, params, user, pass)
	if err != nil {
		return
	}

	// 'action' in the response denotes the next action which the
	// implementation of the revocation endpoint should take.
	action := res.Action

	// The content of the response to the API caller.
	content := res.ResponseContent

	switch action {
	case dto.RevocationAction_INVALID_CLIENT:
		// 401 Unauthorized
		self.ResUtil.UnauthorizedWithContent(ctx, `Basic realm="revocation"`, content)
	case dto.RevocationAction_INTERNAL_SERVER_ERROR:
		// 500 Internal Server Error
		self.ResUtil.InternalServerError(ctx, content)
	case dto.RevocationAction_BAD_REQUEST:
		// 400 Bad Request
		self.ResUtil.BadRequest(ctx, content)
	case dto.RevocationAction_OK:
		// 200 OK
		self.ResUtil.OkJavaScript(ctx, content)
	default:
		// 500 Internal Server Error
		// /api/auth/revocation API returned an unknown action.
		self.UnknownAction(ctx, `/api/auth/revocation`)
	}
}

func (self *RevocationReqHandler) callRevocationApi(
	ctx *gin.Context, parameters string, user string, pass string) (
	res *dto.RevocationResponse, err *api.AuthleteError) {
	// Create a request for /api/auth/revocation API.
	req := dto.RevocationRequest{}
	req.Parameters = parameters
	req.ClientId = user
	req.ClientSecret = pass

	// Call /api/auth/revocation API.
	res, err = self.Api.Revocation(&req)

	if err != nil {
		self.ResUtil.WithAuthleteError(ctx, err)
	}

	return
}
