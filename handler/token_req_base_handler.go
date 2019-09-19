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

type TokenReqBaseHandler struct {
	BaseReqHandler
}

func (self *TokenReqBaseHandler) TokenIssue(
	ctx *gin.Context, ticket string, subject string, properties []dto.Property) {
	// Call /api/auth/token/issue API.
	res, err := self.callTokenIssueApi(ctx, ticket, subject, properties)
	if err != nil {
		return
	}

	// 'action' in the response denotes the next action which the
	// implementation of the token endpoint should take.
	action := res.Action

	// The content of the response to the client application.
	content := res.ResponseContent

	switch action {
	case dto.TokenIssueAction_INTERNAL_SERVER_ERROR:
		// 500 Internal Server Error
		self.ResUtil.InternalServerError(ctx, content)
	case dto.TokenIssueAction_OK:
		// 200 OK
		self.ResUtil.OkJson(ctx, content)
	default:
		// 500 Internal Server Error
		// The /api/auth/token/issue API returned an unknown action.
		self.UnknownAction(ctx, `/api/auth/token/issue`)
	}
}

func (self *TokenReqBaseHandler) callTokenIssueApi(
	ctx *gin.Context, ticket string, subject string, properties []dto.Property) (
	res *dto.TokenIssueResponse, err *api.AuthleteError) {
	// Prepare a request for /api/auth/token/issue API.
	req := dto.TokenIssueRequest{}
	req.Ticket = ticket
	req.Subject = subject
	req.Properties = properties

	// Call /api/auth/token/issue API.
	res, err = self.Api.TokenIssue(&req)

	if err != nil {
		self.ResUtil.WithAuthleteError(ctx, err)
	}

	return
}

func (self *TokenReqBaseHandler) TokenFail(
	ctx *gin.Context, ticket string, reason dto.TokenFailReason) {
	// Call /api/auth/token/fail API.
	res, err := self.callTokenFailApi(ctx, ticket, reason)
	if err != nil {
		return
	}

	// 'action' in the response denotes the next action which the
	// implementation of the token endpoint should take.
	action := res.Action

	// The content of the response to the client application.
	content := res.ResponseContent

	switch action {
	case dto.TokenFailAction_INTERNAL_SERVER_ERROR:
		// 500 Internal Server Error
		self.ResUtil.InternalServerError(ctx, content)
	case dto.TokenFailAction_BAD_REQUEST:
		// 400 Bad Request
		self.ResUtil.BadRequest(ctx, content)
	default:
		// 500 Internal Server Error
		// The /api/auth/token/fail API returned an unknown action.
		self.UnknownAction(ctx, `/api/auth/token/fail`)
	}
}

func (self *TokenReqBaseHandler) callTokenFailApi(
	ctx *gin.Context, ticket string, reason dto.TokenFailReason) (
	res *dto.TokenFailResponse, err *api.AuthleteError) {
	// Prepare a request for /api/auth/token/fail API.
	req := dto.TokenFailRequest{}
	req.Ticket = ticket
	req.Reason = reason

	// Call /api/auth/token/fail API.
	res, err = self.Api.TokenFail(&req)

	if err != nil {
		self.ResUtil.WithAuthleteError(ctx, err)
	}

	return
}
