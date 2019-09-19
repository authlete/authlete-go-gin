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

type AuthReqDecisionHandler struct {
	AuthReqBaseHandler
	Spi spi.AuthReqDecisionHandlerSpi
}

func (self *AuthReqDecisionHandler) InitWithSpi(
	api api.AuthleteApi, spi spi.AuthReqDecisionHandlerSpi) {
	self.Init(api)
	self.Spi = spi
}

func AuthReqDecisionHandler_New(
	api api.AuthleteApi, spi spi.AuthReqDecisionHandlerSpi) *AuthReqDecisionHandler {
	handler := AuthReqDecisionHandler{}
	handler.InitWithSpi(api, spi)

	return &handler
}

func (self *AuthReqDecisionHandler) Handle(
	ctx *gin.Context, ticket string, claimNames []string, claimLocales []string) {
	// If the user did not grant authorization to the client application.
	if !self.Spi.IsClientAuthorized() {
		// The user denied the authorization request.
		self.AuthorizationFail(ctx, ticket, dto.AuthorizationFailReason_DENIED)
		return
	}

	// The subject (= unique identifier) of the user.
	subject := self.Spi.GetUserSubject()

	// If the subject of the user is not available.
	if subject == `` {
		// The user is not authenticated.
		self.AuthorizationFail(ctx, ticket, dto.AuthorizationFailReason_NOT_AUTHENTICATED)
		return
	}

	// Get the value of the "sub" claim. This is optional. When "sub" is
	// an empty string, the value of "subject" will be used as the value
	// of the "sub" claim.
	sub := self.Spi.GetSub()

	// The time when the user was authenticated.
	authTime := self.Spi.GetUserAuthenticatedAt()

	// The ACR (Authentication Context Class Reference) of the user authentication.
	acr := self.Spi.GetAcr()

	// Collect claim values.
	collector := ClaimCollector_New(subject, claimNames, claimLocales, self.Spi)
	claims := collector.Collect()

	// Properties to be bound to an access token and/or an authorization code.
	properties := self.Spi.GetProperties()

	// Scopes bound to an access token and/or an authorization code. If the
	// value returned from spi.GetScopes() is not nil, the scope set replaces
	// the scopes that were given by the original authorization request.
	scopes := self.Spi.GetScopes()

	// Issue required tokens by calling /api/auth/authorization/issue API.
	self.AuthorizationIssue(
		ctx, ticket, subject, authTime, acr, claims, properties, scopes, sub)
}
