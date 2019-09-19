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

type TokenReqHandlerSpi interface {
	// AuthenticateUser performs user authentication with the given credentials.
	//
	// This method is called only when Resource Owner Password Credentials
	// Grant (RFC 6749, 4.3) was used. Therefore, if you have no plan to
	// supporte the flow, always return None. In mose cases, you don't
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
}
