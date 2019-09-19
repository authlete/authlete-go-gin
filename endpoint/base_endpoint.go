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

package endpoint

import (
	"log"

	"github.com/authlete/authlete-go-gin/web"
	"github.com/authlete/authlete-go/api"
	"github.com/gin-gonic/gin"
)

type BaseEndpoint struct {
	Api     api.AuthleteApi
	ReqUtil web.RequestUtility
	ResUtil web.ResponseUtility
}

// CallerAuthenticateFunc is a function that performs authentication of endpoint callers.
type CallerAuthenticateFunc func(*gin.Context) bool

// CallerRejectFunc is a function that generates a response to reject endpoint callers.
type CallerRejectFunc func(*gin.Context)

func (self *BaseEndpoint) GetAuthleteApi(ctx *gin.Context) api.AuthleteApi {
	if self.Api != nil {
		return self.Api
	}

	// Extract the value of `AuthleteApi` which has been added by
	// the function returned from middleware.AuthleteApi_Conf().
	value, exists := ctx.Get(`AuthleteApi`)
	if !exists {
		self.generateError(ctx, "'AuthleteApi' does not exist in the gin context.")
		return nil
	}

	// Convert the type of the value to `*api.AuthleteApi`.
	api, ok := value.(api.AuthleteApi)
	if !ok {
		self.generateError(ctx, "'AuthleteApi' in the gin context is invalid.")
		return nil
	}

	self.Api = api

	return api
}

func (self *BaseEndpoint) generateError(ctx *gin.Context, message string) {
	log.Println(message)

	resutil := web.ResponseUtility{}
	resutil.WithErrorDescription(ctx, message)
}

func (self *BaseEndpoint) GetAccessTokenValidator(ctx *gin.Context) *web.AccessTokenValidator {
	api := self.GetAuthleteApi(ctx)
	if api == nil {
		return nil
	}

	return web.AccessTokenValidator_New(api)
}

func (self *BaseEndpoint) ExtractAccessToken(ctx *gin.Context) string {
	requtil := web.RequestUtility{}
	return requtil.ExtractBearer(ctx)
}

func (self *BaseEndpoint) ValidateAccessToken(ctx *gin.Context, scopes []string) (
	valid bool, validator *web.AccessTokenValidator) {
	// Access token validator.
	validator = self.GetAccessTokenValidator(ctx)

	// Extract an access token from the request.
	accessToken := self.ExtractAccessToken(ctx)

	// Validate the access token.
	valid = validator.ValidateWithScopes(accessToken, scopes)

	return
}
