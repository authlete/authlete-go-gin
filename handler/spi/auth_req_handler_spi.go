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

package spi

import (
	"github.com/authlete/authlete-go/dto"
)

type AuthReqHandlerSpi interface {
	ClaimProvider

	// GetUserAuthenticatedAt returns the time when the user was authenticated.
	//
	// RETURNS
	//
	// The time when the current user was authenticated. The number of seconds
	// since the Unix epoch. 0 means that the time is unknown.
	GetUserAuthenticatedAt() uint64

	// GetUserSubject returns the subject (= unique identifier) of the user.
	//
	// RETURNS
	//
	// The subject of the user. It must consist of only ASCII letters and its
	// length must not exceed 100.
	GetUserSubject() string

	// GetSub returns the value of the "sub" claim that will be embedded in an ID token.
	//
	// If this method returns an empty string, the value returned from
	// GetUserSubject() will be used.
	//
	// The main purpose of this method is to hide the actual value of the subject
	// from client applications.
	//
	// RETURNS
	//
	// The value of the "sub" claim.
	GetSub() string

	// GetAcr returns the ACR that was satisfied when the user was authenticated.
	//
	// The value returned from this method has an important meaning only when
	// the "acr" claim is requested as an essential claim. See OIDC Core,
	// 5.1.1.1. for details.
	//
	// RETURNS
	//
	// The ACR (authentication context class reference) that was satisfied when
	// the user was authenticated.
	GetAcr() string

	// GetProperties returns arbitrary key-value pairs to be bound to tokens.
	//
	// Properties returned from this method will appear as top-level entries
	// (unless they are marked as hidden) in a JSON response from the
	// authorization server as shown in RFC 6749, 5.1.
	//
	// RETURNS
	//
	// List of properties.
	GetProperties() []dto.Property

	// GetScopes returns scopes to be bound to tokens.
	//
	// If nil is returned, the scopes specified in the original authorization
	// request from the client application are used. In other cases, the
	// specified scopes by this method will replace the original scopes.
	//
	// Even scopes that are not included in the original authorization request
	// can be specified. However, as an exception, the "openid" scope is ignored
	// on Authlete server side if it is not included in the original request.
	// It is because the existence of the "openid" scope considerably changes
	// the validation steps and because adding "openid" triggers generation of
	// an ID token (although the client application has not requested it) and
	// the behavior is a major violation against the specification.
	//
	// If you add "offline_access" scope although it is not included in the
	// original request, keep in mind that the specification requires explicit
	// consent from the user forthe scope (OIDC Core, 11). When "offline_access"
	// is included in the original authorization request, the current
	// implementation of Authlete's /api/auth/authorization API checks whether
	// the authorization request has come along with the "prompt" request
	// parameter and its value includes "consent". However, note that the
	// implementation of Authlete's /api/auth/authorization/issue API does not
	// perform the same validation even if the "offline_access" scope is newly
	// added via this method.
	//
	// RETURNS
	//
	// List of scope names.
	GetScopes() []string
}
