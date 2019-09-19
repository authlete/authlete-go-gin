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

type UserInfoReqHandler struct {
	UserInfoReqBaseHandler
	Spi spi.UserInfoReqHandlerSpi
}

func (self *UserInfoReqHandler) InitWithSpi(api api.AuthleteApi, spi spi.UserInfoReqHandlerSpi) {
	self.Init(api)
	self.Spi = spi
}

func UserInfoReqHandler_New(api api.AuthleteApi, spi spi.UserInfoReqHandlerSpi) *UserInfoReqHandler {
	handler := UserInfoReqHandler{}
	handler.InitWithSpi(api, spi)

	return &handler
}

func (self *UserInfoReqHandler) Handle(ctx *gin.Context) {
	// Extract the access token from 'Authorization: Bearer {accessToken}'.
	accessToken := self.ReqUtil.ExtractBearer(ctx)

	if accessToken == `` {
		// 400 Bad Request with a WWW-Authenticate header.
		challenge := `Bearer error="invalid_token",error_description="An access token is required."`
		self.ResUtil.WwwAuthenticate(ctx, 400, challenge)
		return
	}

	// Call Authlete's /api/auth/userinfo API.
	res, err := self.callUserInfoApi(ctx, accessToken)
	if err != nil {
		return
	}

	// 'action' in the response denotes the next action which the
	// implementation of the userinfo endpoint should take.
	action := res.Action

	// The content of the response to the client application.
	content := res.ResponseContent

	switch action {
	case dto.UserInfoAction_INTERNAL_SERVER_ERROR:
		// 500 Internal Server Error
		self.ResUtil.WwwAuthenticate(ctx, 500, content)
	case dto.UserInfoAction_BAD_REQUEST:
		// 400 Bad Request
		self.ResUtil.WwwAuthenticate(ctx, 400, content)
	case dto.UserInfoAction_UNAUTHORIZED:
		// 401 Unauthorized
		self.ResUtil.WwwAuthenticate(ctx, 401, content)
	case dto.UserInfoAction_FORBIDDEN:
		// 403 Forbidden
		self.ResUtil.WwwAuthenticate(ctx, 403, content)
	case dto.UserInfoAction_OK:
		// Return information about the user.
		self.generateUserInfo(ctx, res)
	default:
		// 500 Internal Server Error
		// The /api/auth/userinfo API returned an unknown action.
		self.UnknownAction(ctx, `/api/auth/userinfo`)
	}
}

func (self *UserInfoReqHandler) callUserInfoApi(
	ctx *gin.Context, accessToken string) (
	res *dto.UserInfoResponse, err *api.AuthleteError) {
	// Prepare a request for /api/auth/userinfo API.
	req := dto.UserInfoRequest{}
	req.Token = accessToken

	// Call /api/auth/userinfo API.
	res, err = self.Api.UserInfo(&req)

	if err != nil {
		self.ResUtil.WithAuthleteError(ctx, err)
	}

	return
}

func (self *UserInfoReqHandler) generateUserInfo(ctx *gin.Context, res *dto.UserInfoResponse) {
	// Collect information about the user.
	collector := ClaimCollector_New(res.Subject, res.Claims, nil, self.Spi)
	claims := collector.Collect()

	// The value of the "sub" claim (optional)
	sub := self.Spi.GetSub()

	// Generate a response which is returned from the userinfo endpoint.
	self.UserInfoIssue(ctx, res.Token, claims, sub)
}
