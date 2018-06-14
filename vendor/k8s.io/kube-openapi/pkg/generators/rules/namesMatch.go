/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Go field names must be CamelCase. JSON field names must be camelCase. Other than capitalization of the initial letter, the two should almost always match. No underscores nor dashes in either.
// This rule verifies the convention "Other than capitalization of the initial letter, the two should almost always match."
// Examples (also in unit test):
//     Go name      | JSON name    | match
//     PodSpec        podSpec        true
//     podSpec        podSpec        false
//     PodSpec        spec           false
//     Spec           podSpec        false
//     JSONSpec       jsonSpec       true
//     JSONSpec       jsonspec       false
//     HTTPJSONSpec   httpJSONSpec   true
// NOTE: this validator cannot tell two sequential all-capital words from one word, therefore the case below
// is also considered matched.
//     HTTPJSONSpec   httpjsonSpec   true

package rules

import (
	"reflect"
	"strings"

	"k8s.io/gengo/types"
)

// NamesMatch implements APIRule interface. It verifies the convention "Other than
// capitalization of the initial letter, the two should almost always match."
type NamesMatch struct{}

// Name returns the name of APIRule. The name should not contain ','.
func (n *NamesMatch) Name() string {
	return "names_match"
}

// Validate evaluates API rule on type t and returns a list of field names in
// the type that violate the rule. Empty field name [""] implies the entire
// type violates the rule.
func (n *NamesMatch) Validate(t *types.Type) ([]string, error) {
	fields := make([]string, 0)

	// Only validate struct type and ignore the rest
	switch t.Kind {
	case types.Struct:
		for _, m := range t.Members {
			goName := m.Name
			jsonTag := reflect.StructTag(m.Tags).Get("json")
			jsonName := strings.Split(jsonTag, ",")[0]
			// Skip empty JSON name "", omitted JSON name "-" and special case "metadata"
			// for object and list meta.
			if jsonName == "" || jsonName == "-" || jsonName == "metadata" {
				continue
			}

			if !namesMatch(goName, jsonName) {
				fields = append(fields, goName)
			}
		}
	}
	return fields, nil
}

func namesMatch(goName, jsonName string) bool {
	if strings.ToLower(goName) != strings.ToLower(jsonName) {
		return false
	}
	if len(jsonName) == 0 {
		return true
	}
	// Go field names must be CamelCase
	if !isCapital(goName[0]) {
		return false
	}
	for i := 0; i < len(goName); i++ {
		if goName[i] == jsonName[i] {
			return (i == 1 || isCapital(goName[i])) && goName[i:] == jsonName[i:]
		}
	}
	return true
}

func isCapital(b byte) bool {
	return b >= 'A' && b <= 'Z'
}
