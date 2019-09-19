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
	"github.com/authlete/authlete-go-gin/handler/spi"
	"github.com/authlete/authlete-go/api"
	"github.com/authlete/authlete-go/dto"
	"github.com/gin-gonic/gin"
)

type TokenReqHandler struct {
	TokenReqBaseHandler
	Spi spi.TokenReqHandlerSpi
}

func (self *TokenReqHandler) InitWithSpi(api api.AuthleteApi, spi spi.TokenReqHandlerSpi) {
	self.Init(api)
	self.Spi = spi
}

func TokenReqHandler_New(api api.AuthleteApi, spi spi.TokenReqHandlerSpi) *TokenReqHandler {
	handler := TokenReqHandler{}
	handler.InitWithSpi(api, spi)

	return &handler
}

func (self *TokenReqHandler) Handle(ctx *gin.Context) {
	// Message body of the request.
	params := self.ReqUtil.ExtractForm(ctx)

	// Basic authentication of the request.
	user, pass, _ := ctx.Request.BasicAuth()

	// Call Authlete's /api/auth/token API.
	res, err := self.callTokenApi(ctx, params, user, pass)
	if err != nil {
		return
	}

	// 'action' in the response denotes the next action which the
	// implementation of the token endpoint should take.
	action := res.Action

	// The content of the response to the client application.
	content := res.ResponseContent

	switch action {
	case dto.TokenAction_INVALID_CLIENT:
		// 401 Unauthorized
		self.ResUtil.UnauthorizedWithContent(ctx, `Basic realm="token"`, content)
	case dto.TokenAction_INTERNAL_SERVER_ERROR:
		// 500 Internal Server Error
		self.ResUtil.InternalServerError(ctx, content)
	case dto.TokenAction_BAD_REQUEST:
		// 400 Bad Request
		self.ResUtil.BadRequest(ctx, content)
	case dto.TokenAction_PASSWORD:
		// Process the token request whose flow is
		// "Resource Owner Password Credentials".
		self.handlePassword(ctx, res)
	case dto.TokenAction_OK:
		// 200 OK
		self.ResUtil.OkJson(ctx, content)
	default:
		// 500 Internal Server Error
		// /api/auth/token API returned an unknown action.
		self.UnknownAction(ctx, `/api/auth/token`)
	}
}

func (self *TokenReqHandler) callTokenApi(
	ctx *gin.Context, parameters string, user string, pass string) (
	res *dto.TokenResponse, err *api.AuthleteError) {
	// Prepare a request for /api/auth/token API.
	req := dto.TokenRequest{}
	req.Parameters = parameters
	req.ClientId = user
	req.ClientSecret = pass
	req.Properties = self.Spi.GetProperties()

	// Call /api/auth/token API.
	res, err = self.Api.Token(&req)

	if err != nil {
		self.ResUtil.WithAuthleteError(ctx, err)
	}

	return
}

func (self *TokenReqHandler) handlePassword(ctx *gin.Context, res *dto.TokenResponse) {
	// The ticket to call Authlete's /api/auth/token/* API with.
	ticket := res.Ticket

	// The credentials of the resource owner.
	username := res.Username
	password := res.Password

	// Validate the credentials (user authentication)
	subject := self.Spi.AuthenticateUser(username, password)

	// If the credentials of the resource owner are invalid.
	if subject == `` {
		// The credentials are invalid. Nothing is issued.
		reason := dto.TokenFailReason_INVALID_RESOURCE_OWNER_CREDENTIALS
		self.TokenFail(ctx, ticket, reason)
		return
	}

	// Issue tokens.
	self.TokenIssue(ctx, ticket, subject, self.Spi.GetProperties())
}
