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

type AuthReqErrorHandler struct {
	AuthReqBaseHandler
}

func AuthReqErrorHandler_New(api api.AuthleteApi) *AuthReqErrorHandler {
	handler := AuthReqErrorHandler{}
	handler.Init(api)

	return &handler
}

func (self *AuthReqErrorHandler) Handle(ctx *gin.Context, res *dto.AuthorizationResponse) {
	// 'action' in the response denotes the next action which the
	// implementation of the authorization endpoint should take.
	action := res.Action

	// The content of the response to the user agent. The format
	// of the content varies depending on the action.
	content := res.ResponseContent

	switch action {
	case dto.AuthorizationAction_INTERNAL_SERVER_ERROR:
		// 500 Internal Server Error
		self.ResUtil.InternalServerError(ctx, content)
	case dto.AuthorizationAction_BAD_REQUEST:
		// 400 Bad Request
		self.ResUtil.BadRequest(ctx, content)
	case dto.AuthorizationAction_LOCATION:
		// 302 Found
		self.ResUtil.Location(ctx, content)
	case dto.AuthorizationAction_FORM:
		// 200 OK
		self.ResUtil.OkHtml(ctx, content)
	case dto.AuthorizationAction_INTERACTION:
		// This is not an error case. The implementation of the
		// authorization endpoint should show an authorization
		// page to the user.
	case dto.AuthorizationAction_NO_INTERACTION:
		// This is not an error case. The implementation of the
		// authorization endpoint should handle the authorization
		// request without user interaction.
	default:
		// 500 Internal Server Error
		// Authlete's /api/auth/authorization API returned unknown action.
		self.UnknownAction(ctx, `/api/auth/authorization`)
	}
}
