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

type AuthReqHandlerSpiAdapter struct {
	ClaimProviderAdapter
}

func (self *AuthReqHandlerSpiAdapter) GetUserAuthenticatedAt() uint64 {
	return uint64(0)
}

func (self *AuthReqHandlerSpiAdapter) GetUserSubject() string {
	return ``
}

func (self *AuthReqHandlerSpiAdapter) GetSub() string {
	return ``
}

func (self *AuthReqHandlerSpiAdapter) GetAcr() string {
	return ``
}

func (self *AuthReqHandlerSpiAdapter) GetProperties() []dto.Property {
	return nil
}

func (self *AuthReqHandlerSpiAdapter) GetScopes() []string {
	return nil
}
