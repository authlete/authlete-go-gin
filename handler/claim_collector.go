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
	"strings"

	"github.com/authlete/authlete-go-gin/handler/spi"
)

type ClaimCollector struct {
	subject       string
	claimNames    []string
	claimLocales  []string
	claimProvider spi.ClaimProvider
}

func (self *ClaimCollector) Init(
	subject string, claimNames []string, claimLocales []string, claimProvider spi.ClaimProvider) {
	self.subject = subject
	self.claimNames = claimNames
	self.claimLocales = normalizeClaimLocales(claimLocales)
	self.claimProvider = claimProvider
}

func ClaimCollector_New(
	subject string, claimNames []string, claimLocales []string,
	claimProvider spi.ClaimProvider) *ClaimCollector {
	collector := ClaimCollector{}
	collector.Init(subject, claimNames, claimLocales, claimProvider)

	return &collector
}

func normalizeClaimLocales(claimLocales []string) []string {
	if len(claimLocales) == 0 {
		return nil
	}

	// From "5.2. Claims Languages and Scripts" in OpenID Connect Core 1.0
	//
	//   However, since BCP47 language tag values are case insensitive,
	//   implementations SHOULD interpret the language tag values supplied
	//   in a case insensitive manner.
	//

	localeList := make([]string, 0)
	localeSet := make(map[string]bool)

	for _, claimLocale := range claimLocales {
		if claimLocale == `` {
			continue
		}

		locale := strings.ToLower(claimLocale)
		if localeSet[locale] {
			continue
		}

		localeSet[locale] = true
		localeList = append(localeList, locale)
	}

	if len(localeList) == 0 {
		return nil
	}

	return localeList
}

func (self *ClaimCollector) Collect() map[string]interface{} {
	if self.claimNames == nil || len(self.claimNames) == 0 {
		return nil
	}

	// Pairs of claim name and its value.
	collectedClaims := make(map[string]interface{})

	// For each required claim.
	for _, claimName := range self.claimNames {
		if claimName == `` {
			continue
		}

		// Split the claim name into the name part and the language tag part.
		elements := strings.SplitN(claimName, `#`, 2)
		name := elements[0]
		tag := ``
		if len(elements) == 2 {
			tag = elements[1]
		}

		// If the name part is empty.
		if name == `` {
			continue
		}

		// Get the value of the claim.
		value := self.getClaimValue(name, tag)

		// If the value of the claim was not obtained.
		if value == nil {
			continue
		}

		// If the type is string but its value is empty.
		str, ok := value.(string)
		if ok && str == `` {
			continue
		}

		// Just for an edge case where claimName ends with '#'. e.g. 'family_name#'
		if tag == `` {
			claimName = name
		}

		// Add the pair of the claim name and its value.
		collectedClaims[claimName] = value
	}

	// If no claim value has been obtained.
	if len(collectedClaims) == 0 {
		return nil
	}

	return collectedClaims
}

func (self *ClaimCollector) getClaimValue(name string, tag string) interface{} {
	provider := self.claimProvider
	subject := self.subject

	// If a language tag is explicitly appended.
	if tag != `` {
		// Get the claim value of the claim with the specific language tag.
		return provider.GetUserClaimValue(subject, name, tag)
	}

	// If claim locales are not specified by the "claims_locales" parameter.
	if self.claimLocales == nil {
		// Get the claim value of the claim without any language tag.
		return provider.GetUserClaimValue(subject, name, ``)
	}

	// For each claim locale. They are ordered by preference.
	for _, locale := range self.claimLocales {
		// Try to get the claim value with the claim locale.
		value := provider.GetUserClaimValue(subject, name, locale)

		// If the claim value was obtained.
		if value != `` {
			return value
		}
	}

	// The last resort. Try to get the claim value without any language tag.
	return provider.GetUserClaimValue(subject, name, ``)
}
