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
	"fmt"

	"github.com/authlete/authlete-go-gin/web"
	"github.com/authlete/authlete-go/api"
	"github.com/gin-gonic/gin"
)

type BaseReqHandler struct {
	Api     api.AuthleteApi
	ReqUtil web.RequestUtility
	ResUtil web.ResponseUtility
}

func (self *BaseReqHandler) Init(api api.AuthleteApi) {
	self.Api = api
	self.ReqUtil = web.RequestUtility{}
	self.ResUtil = web.ResponseUtility{}
}

func (self *BaseReqHandler) UnknownAction(ctx *gin.Context, apiPath string) {
	description := fmt.Sprintf(`Authlete's %s API returned an unknown action.`, apiPath)
	self.ResUtil.WithErrorDescription(ctx, description)
}
