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
	"time"

	"github.com/authlete/authlete-go-gin/handler/spi"
	"github.com/authlete/authlete-go/api"
	"github.com/authlete/authlete-go/dto"
	"github.com/gin-gonic/gin"
)

type NoInteractionHandler struct {
	AuthReqBaseHandler
	Spi spi.NoInteractionHandlerSpi
}

func (self *NoInteractionHandler) InitWithSpi(
	api api.AuthleteApi, spi spi.NoInteractionHandlerSpi) {
	self.Init(api)
	self.Spi = spi
}

func NoInteractionHandler_New(
	api api.AuthleteApi, spi spi.NoInteractionHandlerSpi) *NoInteractionHandler {
	handler := NoInteractionHandler{}
	handler.InitWithSpi(api, spi)

	return &handler
}

func (self *NoInteractionHandler) Handle(ctx *gin.Context, res *dto.AuthorizationResponse) {
	// If the value of the "action" parameter in the response from Authlete's
	// /api/auth/authorization API is not "NO_INTERACTION".
	if res.Action != dto.AuthorizationAction_NO_INTERACTION {
		// This handler does not handle other cases than NO_INTERACTION.
		return
	}

	// Check 1: User Authentication
	if self.checkUserAuthentication() == false {
		// A user must have logged in.
		reason := dto.AuthorizationFailReason_NOT_LOGGED_IN
		self.AuthorizationFail(ctx, res.Ticket, reason)
		return
	}

	// Get the last time when the user was authenticated.
	authTime := self.Spi.GetUserAuthenticatedAt()

	// Check 2: Max Age
	if self.checkMaxAge(res, authTime) == false {
		// The maximum authentication age has elapsed since the last time
		// when the user was authenticated.
		reason := dto.AuthorizationFailReason_EXCEEDS_MAX_AGE
		self.AuthorizationFail(ctx, res.Ticket, reason)
		return
	}

	// The subject (unique ID) of the current user.
	subject := self.Spi.GetUserSubject()

	// Check 3: Subject
	if self.checkSubject(res, subject) == false {
		// The requested subject and that of the current user don't match.
		reason := dto.AuthorizationFailReason_DIFFERENT_SUBJECT
		self.AuthorizationFail(ctx, res.Ticket, reason)
		return
	}

	// Get the value of the "sub" claim. This is optional. When "sub" is empty,
	// the value of "subject" will be used as the value of the "sub" claim.
	sub := self.Spi.GetSub()

	// Get the ACR that was satisfied when the current user was authenticated.
	acr := self.Spi.GetAcr()

	// Check 4: ACR
	if self.checkAcr(res, acr) == false {
		// None of the requested ACRs is satisfied.
		reason := dto.AuthorizationFailReason_ACR_NOT_SATISFIED
		self.AuthorizationFail(ctx, res.Ticket, reason)
		return
	}

	// Collect claim values.
	collector := ClaimCollector_New(subject, res.Claims, res.ClaimsLocales, self.Spi)
	claims := collector.Collect()

	// Properties to be bound to an access token and/or an authorization code.
	properties := self.Spi.GetProperties()

	// Scopes associated with an access token and/or an authorization code.
	// If the value returned from spi.GetScopes() is not nil, the scope set
	// replaces the scopes that were given by the original authorization request.
	scopes := self.Spi.GetScopes()

	// Issue tokens without user interaction.
	self.AuthorizationIssue(
		ctx, res.Ticket, subject, authTime, acr, claims, properties, scopes, sub)
}

func (self *NoInteractionHandler) checkUserAuthentication() bool {
	return self.Spi.IsUserAuthenticated()
}

func (self *NoInteractionHandler) checkMaxAge(
	res *dto.AuthorizationResponse, authTime uint64) bool {
	// Get the requested maximum authentication age.
	maxAge := res.MaxAge

	// If no maximum authentication age is requested.
	if maxAge == uint32(0) {
		// No need to care about the maximum authentication age.
		return true
	}

	// The time in seconds with the authentication expires.
	expiresAt := authTime + uint64(maxAge)

	// If the authentication has not expired yet.
	current := uint64(time.Now().Unix())
	if current < expiresAt {
		// The authentication has not expired.
		return true
	}

	// The authentication has expired.
	return false
}

func (self *NoInteractionHandler) checkSubject(
	res *dto.AuthorizationResponse, subject string) bool {
	// Get the requested subject.
	requestedSubject := res.Subject

	// If no subject is requested.
	if requestedSubject == `` {
		// No need to care about the subject.
		return true
	}

	// If the requested subject matches that of the current user.
	if requestedSubject == subject {
		// The subjects match.
		return true
	}

	// The subjects don't match.
	return false
}

func (self *NoInteractionHandler) checkAcr(
	res *dto.AuthorizationResponse, acr string) bool {
	// Get the list of requested ACRs.
	requestedAcrs := res.Acrs

	// If no ACR is requested
	if len(requestedAcrs) == 0 {
		// No need to care about ACRs.
		return true
	}

	// For each requested ACR
	for _, requestedAcr := range requestedAcrs {
		if requestedAcr == acr {
			// OK. The ACR satisfied when the current user was authenticated
			// matches one of the requested ACRs.
			return true
		}
	}

	// If one of the requested ACRs must be satisfied.
	if res.AcrEssential {
		// None of the requested ACRs is satisfied.
		return false
	}

	// The ACR satisfied when the current user was authenticated does not
	// match any one of the requested ACRs, but the authorization request
	//  from the client application did not request ACR as essential.
	// Therefore, it is not necessary to raise an error here.
	return true
}
