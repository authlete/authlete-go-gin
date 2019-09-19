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

type UserInfoReqBaseHandler struct {
	BaseReqHandler
}

func (self *UserInfoReqBaseHandler) UserInfoIssue(
	ctx *gin.Context, token string, claims map[string]interface{}, sub string) {
	// Call Authlete's /api/auth/userinfo/issue API.
	res, err := self.callUserInfoIssueApi(ctx, token, claims, sub)
	if err != nil {
		return
	}

	// 'action' in the response denotes the next action which the
	// implementation of the userinfo endpoint should take.
	action := res.Action

	// The content of the response to the client application.
	content := res.ResponseContent

	switch action {
	case dto.UserInfoIssueAction_INTERNAL_SERVER_ERROR:
		// 500 Internal Server Error
		self.ResUtil.WwwAuthenticate(ctx, 500, content)
	case dto.UserInfoIssueAction_BAD_REQUEST:
		// 400 Bad Request
		self.ResUtil.WwwAuthenticate(ctx, 400, content)
	case dto.UserInfoIssueAction_UNAUTHORIZED:
		// 401 Unauthorized
		self.ResUtil.WwwAuthenticate(ctx, 401, content)
	case dto.UserInfoIssueAction_FORBIDDEN:
		// 403 Forbidden
		self.ResUtil.WwwAuthenticate(ctx, 403, content)
	case dto.UserInfoIssueAction_JSON:
		// 200 OK, application/json;charset=UTF-8
		self.ResUtil.OkJson(ctx, content)
	case dto.UserInfoIssueAction_JWT:
		// 200 OK, application/jwt
		self.ResUtil.OkJwt(ctx, content)
	default:
		// 500 Internal Server Error
		// The /api/auth/userinfo/issue API returned an unknown action.
		self.UnknownAction(ctx, `/api/auth/userinfo/issue`)
	}
}

func (self *UserInfoReqBaseHandler) callUserInfoIssueApi(
	ctx *gin.Context, token string, claims map[string]interface{}, sub string) (
	res *dto.UserInfoIssueResponse, err *api.AuthleteError) {
	// Prepare a request for /api/auth/userinfo/issue API.
	req := dto.UserInfoIssueRequest{}
	req.Token = token
	req.Sub = sub

	if claims != nil {
		byt, _ := json.Marshal(claims)
		req.Claims = string(byt)
	}

	// Call /api/auth/userinfo/issue API.
	res, err = self.Api.UserInfoIssue(&req)

	if err != nil {
		self.ResUtil.WithAuthleteError(ctx, err)
	}

	return
}
