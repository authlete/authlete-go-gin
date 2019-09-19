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

package middleware

import (
	"log"

	"github.com/authlete/authlete-go/api"
	"github.com/authlete/authlete-go/conf"
	"github.com/gin-gonic/gin"
)

func AuthleteApi_Conf(cnf conf.AuthleteConfiguration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(`AuthleteApi`, api.New(cnf))
		ctx.Next()
	}
}

func AuthleteApi_Env() gin.HandlerFunc {
	cnf := conf.AuthleteEnvConfiguration{}
	return AuthleteApi_Conf(&cnf)
}

func AuthleteApi_Toml(file string) gin.HandlerFunc {
	cnf := conf.AuthleteTomlConfiguration{}

	err := cnf.Load(file)
	if err != nil {
		log.Printf("Failed to load '%s'\n", file)
	}

	return AuthleteApi_Conf(&cnf)
}
