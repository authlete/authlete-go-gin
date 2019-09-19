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

package web

import (
	"fmt"

	"github.com/authlete/authlete-go/api"
	"github.com/gin-gonic/gin"
)

type ResponseUtility struct {
}

func (self *ResponseUtility) OkJson(ctx *gin.Context, content string) {
	// 200 OK, application/json;charset=UTF-8
	self.json(ctx, 200, content)
}

func (self *ResponseUtility) OkJavaScript(ctx *gin.Context, content string) {
	// 200 OK, application/javascript;charset=UTF-8
	self.javascript(ctx, 200, content)
}

func (self *ResponseUtility) OkJwt(ctx *gin.Context, content string) {
	// 200 OK, application/jwt
	self.jwt(ctx, 200, content)
}

func (self *ResponseUtility) OkHtml(ctx *gin.Context, content string) {
	// 200 OK, text/html;charset=UTF-8
	self.html(ctx, 200, content)
}

func (self *ResponseUtility) NoContent(ctx *gin.Context) {
	// 204 No Content
	self.response(ctx, 204, ``, ``)
}

func (self *ResponseUtility) Location(ctx *gin.Context, location string) {
	// 302 Found with a Location header
	ctx.Header(`Location`, location)
	self.response(ctx, 302, ``, ``)
}

func (self *ResponseUtility) BadRequest(ctx *gin.Context, content string) {
	// 400 Bad Request, application/json;charset=UTF-8
	self.json(ctx, 400, content)
}

func (self *ResponseUtility) Unauthorized(ctx *gin.Context, challenge string) {
	// 401 Unauthorized with a WWW-Authenticate header
	self.UnauthorizedWithContent(ctx, challenge, ``)
}

func (self *ResponseUtility) UnauthorizedWithContent(ctx *gin.Context, challenge string, content string) {
	// 401 Unauthorized with a WWW-Authenticate header
	self.WwwAuthenticateWithContent(ctx, 401, challenge, content)
}

func (self *ResponseUtility) Forbidden(ctx *gin.Context, content string) {
	// 403 Forbidden, application/json;charset=UTF-8
	self.json(ctx, 403, content)
}

func (self *ResponseUtility) NotFound(ctx *gin.Context, content string) {
	// 404 Not Found, application/json;charset=UTF-8
	self.json(ctx, 404, content)
}

func (self *ResponseUtility) InternalServerError(ctx *gin.Context, content string) {
	// 500 Internal Server Error, application/json;charset=UTF-8
	self.json(ctx, 500, content)
}

func (self *ResponseUtility) WithErrorDescription(ctx *gin.Context, description string) {
	content := fmt.Sprintf(`{
  "error": "server_error",
  "error_description": "%s"
}
`, description)
	self.InternalServerError(ctx, content)
}

func (self *ResponseUtility) WithAuthleteError(ctx *gin.Context, err *api.AuthleteError) {
	description := err.Error()

	if err.Cause != nil {
		description = err.Cause.Error()
	}

	self.WithErrorDescription(ctx, description)
}

func (self *ResponseUtility) WwwAuthenticate(ctx *gin.Context, status int, challenge string) {
	self.WwwAuthenticateWithContent(ctx, status, challenge, ``)
}

func (self *ResponseUtility) WwwAuthenticateWithContent(ctx *gin.Context, status int, challenge string, content string) {
	ctx.Header(`WWW-Authenticate`, challenge)
	self.json(ctx, status, content)
}

func (self *ResponseUtility) response(
	ctx *gin.Context, status int, content string, contentType string) {
	ctx.Header(`Cache-Control`, `no-store`)
	ctx.Header(`Pragma`, `no-cache`)

	if content == `` {
		ctx.Status(status)
	} else {
		ctx.Data(status, contentType, []byte(content))
	}
}

func (self *ResponseUtility) json(ctx *gin.Context, status int, content string) {
	self.response(ctx, status, content, `application/json;charset=UTF-8`)
}

func (self *ResponseUtility) javascript(ctx *gin.Context, status int, content string) {
	self.response(ctx, status, content, `application/javascript;charset=UTF-8`)
}

func (self *ResponseUtility) jwt(ctx *gin.Context, status int, content string) {
	self.response(ctx, status, content, `application/jwt`)
}

func (self *ResponseUtility) html(ctx *gin.Context, status int, content string) {
	self.response(ctx, status, content, `text/html;charset=UTF-8`)
}
