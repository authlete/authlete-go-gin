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

type ClaimProvider interface {
	// GetUserClaimValue returns the value of a specified claim.
	//
	// ARGS
	//
	//     subject: The subject (unique identifier) of a user.
	//     claimName: A claim name such as "name" and "family_name".
	//     languageTag: A language tag such as "en" and "ja".
	//
	// RETURNS
	//
	// The value of the claim. nil if the value is not available.
	// In most cases, an instance of string should be returned. When
	// "claimName" is "address", an instance of *dto.Address should
	// be returned.
	//
	GetUserClaimValue(subject string, claimName string, languageTag string) interface{}
}
