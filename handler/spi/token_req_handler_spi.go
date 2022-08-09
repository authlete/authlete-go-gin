//
// Copyright (C) 2019-2022 Authlete, Inc.
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
	"github.com/gin-gonic/gin"
)

type TokenReqHandlerSpi interface {
	// AuthenticateUser performs user authentication with the given credentials.
	//
	// This method is called only when Resource Owner Password Credentials
	// Grant (RFC 6749, 4.3) was used. Therefore, if you have no plan to
	// support the flow, always return None. In most cases, you don't
	// have to support the flow. RFC 6749 says "The authorization server
	// should take special care when enabling this grant type and only
	// allow it when other flows are not viable."
	//
	// ARGS
	//
	//     username: The value of the "username" request parameter of the token request.
	//     password: The value of the "password" request parameter of the token request.
	//
	// RETURNS
	//
	// The subject (= unique identifier) of the authenticated user.
	// An empty string if not authenticated.
	AuthenticateUser(username string, password string) string

	// GetProperties returns arbitrary key-value pairs to be bound to an access token.
	//
	// Properties returned from this method will appear as top-level entries
	// (unless they are marked as hidden) in a JSON response from the
	// authorization server as shown in RFC 6749, 5.1.
	//
	// RETURNS
	//
	// List of properties.
	GetProperties() []dto.Property

	// Handle a token exchange request.
	//
	// This method is called when the grant type of the token request is
	// "urn:ietf:params:oauth:grant-type:token-exchange". The grant type is
	// defined in RFC 8693 OAuth 2.0 Token Exchange.
	//
	// RFC 8693 is very flexible. In other words, the specification does not
	// define details that are necessary for secure token exchange. Therefore,
	// implementations have to complement the specification with their own rules.
	//
	// The second argument passed to this method is an instance of TokenResponse
	// that represents a response from Authlete's /auth/token API. The instance
	// contains information about the token exchange request such as the value
	// of the "subject_token" request parameter. Implementations of this method
	// are supposed to (1) validate the information based on their own rules,
	// (2) generate a token (e.g. an access token) using the information, and
	// (3) prepare a token response in the JSON format that conforms to Section
	// 2.2 of RFC 8693.
	//
	// Authlete's /auth/token API performs validation of token exchange requests
	// to some extent. Therefore, authorization server implementations don't
	// have to repeat the same validation steps. See the online document on
	// Authlete website for details.
	//
	// NOTE: Token Exchange is supported by Authlete 2.3 and newer versions.
	// If the Authlete server of your system is older than version 2.3, the
	// grant type ("urn:ietf:params:oauth:grant-type:token-exchange") is not
	// supported and so this method is never called.
	//
	// Since v1.0.5.
	//
	// ARGS
	//
	//     ctx: A context which can be used to prepare a token response
	//     res: A response from Authlete's /auth/token API
	//
	// RETURNS
	//
	// true to indicate that the implementation of this method has prepared
	// a token response. false to indicate that the implementation of this
	// method has done nothing. When false is returned, TokenReqHandler
	// will generate 400 Bad Request with "error":"unsupported_grant_type".
	TokenExchange(ctx *gin.Context, res *dto.TokenResponse) bool
}
