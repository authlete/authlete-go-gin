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
	"encoding/json"

	"github.com/authlete/authlete-go/api"
	"github.com/authlete/authlete-go/dto"
	"github.com/gin-gonic/gin"
)

type AuthReqBaseHandler struct {
	BaseReqHandler
}

func (self *AuthReqBaseHandler) AuthorizationIssue(
	ctx *gin.Context, ticket string, subject string, authTime uint64,
	acr string, claims map[string]interface{}, properties []dto.Property,
	scopes []string, sub string) {
	// Call Authlete's /api/auth/authorization/issue API.
	res, err := self.callAuthorizationIssueApi(
		ctx, ticket, subject, authTime, acr, claims, properties, scopes, sub)
	if err != nil {
		return
	}

	// 'action' in the response denotes the next action which the
	// implementation of the authorization endpoint should take.
	action := res.Action

	// The content of the response to the user agent. The format
	// of the content varies depending on the action.
	content := res.ResponseContent

	switch action {
	case dto.AuthorizationIssueAction_INTERNAL_SERVER_ERROR:
		// 500 Internal Server Error
		self.ResUtil.InternalServerError(ctx, content)
	case dto.AuthorizationIssueAction_BAD_REQUEST:
		// 400 Bad Request
		self.ResUtil.BadRequest(ctx, content)
	case dto.AuthorizationIssueAction_LOCATION:
		// 302 Found
		self.ResUtil.Location(ctx, content)
	case dto.AuthorizationIssueAction_FORM:
		// 200 OK
		self.ResUtil.OkHtml(ctx, content)
	default:
		// 500 Internal Server Error
		// The /api/auth/authorization/issue API returned an unknown action.
		self.UnknownAction(ctx, `/api/auth/authorization/issue`)
	}
}

func (self *AuthReqBaseHandler) callAuthorizationIssueApi(
	ctx *gin.Context, ticket string, subject string, authTime uint64,
	acr string, claims map[string]interface{}, properties []dto.Property,
	scopes []string, sub string) (res *dto.AuthorizationIssueResponse, err *api.AuthleteError) {
	// Prepare a request for /api/auth/authorization/issue API.
	req := dto.AuthorizationIssueRequest{}

	req.Ticket = ticket
	req.Subject = subject
	req.AuthTime = authTime
	req.Acr = acr
	req.Properties = properties
	req.Scopes = scopes
	req.Sub = sub

	if claims != nil {
		byt, _ := json.Marshal(claims)
		req.Claims = string(byt)
	}

	// Call /api/auth/authorization/issue API.
	res, err = self.Api.AuthorizationIssue(&req)

	if err != nil {
		self.ResUtil.WithAuthleteError(ctx, err)
	}

	return
}

func (self *AuthReqBaseHandler) AuthorizationFail(
	ctx *gin.Context, ticket string, reason dto.AuthorizationFailReason) {
	// Call /api/auth/authorization/fail API.
	res, err := self.callAuthorizationFailApi(ctx, ticket, reason)
	if err != nil {
		return
	}

	// 'action' in the response denotes the next action which the
	// implementation of the authorization endpoint should take.
	action := res.Action

	// The content of the response to the user agent. The format
	// of the content varies depending on the action.
	content := res.ResponseContent

	switch action {
	case dto.AuthorizationFailAction_INTERNAL_SERVER_ERROR:
		// 500 Internal Server Error
		self.ResUtil.InternalServerError(ctx, content)
	case dto.AuthorizationFailAction_BAD_REQUEST:
		// 400 Bad Request
		self.ResUtil.BadRequest(ctx, content)
	case dto.AuthorizationFailAction_LOCATION:
		// 302 Found
		self.ResUtil.Location(ctx, content)
	case dto.AuthorizationFailAction_FORM:
		// 200 OK
		self.ResUtil.OkHtml(ctx, content)
	default:
		// 500 Internal Server Error
		// The /api/auth/authorization/fail API returned an unknown action.
		self.UnknownAction(ctx, `/api/auth/authorization/fail`)
	}
}

func (self *AuthReqBaseHandler) callAuthorizationFailApi(
	ctx *gin.Context, ticket string, reason dto.AuthorizationFailReason) (
	res *dto.AuthorizationFailResponse, err *api.AuthleteError) {
	// Prepare a request for /api/auth/authorization/fail API.
	req := dto.AuthorizationFailRequest{}
	req.Ticket = ticket
	req.Reason = reason

	// Call /api/auth/authorization/fail API.
	res, err = self.Api.AuthorizationFail(&req)

	if err != nil {
		self.ResUtil.WithAuthleteError(ctx, err)
	}

	return
}
