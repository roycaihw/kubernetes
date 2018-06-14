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

package rules

import "testing"

func TestNamesMatch(t *testing.T) {
	tcs := []struct {
		goName   string
		jsonName string
		match    bool
	}{
		{"PodSpec", "podSpec", true},
		{"podSpec", "podSpec", false},
		{"PodSpec", "spec", false},
		{"Spec", "podSpec", false},
		{"JSONSpec", "jsonSpec", true},
		{"JSONSpec", "jsonspec", false},
		{"HTTPJSONSpec", "httpJSONSpec", true},
		// NOTE: this validator cannot tell two sequential all-capital words from one word,
		// therefore the case below is also considered matched.
		{"HTTPJSONSpec", "httpjsonSpec", true},
	}

	for _, tc := range tcs {
		match := namesMatch(tc.goName, tc.jsonName)
		if match != tc.match {
			t.Errorf("unexpected match result: goName %v, jsonName %v, want: %v, got: %v",
				tc.goName, tc.jsonName, tc.match, match)
		}
	}
}
