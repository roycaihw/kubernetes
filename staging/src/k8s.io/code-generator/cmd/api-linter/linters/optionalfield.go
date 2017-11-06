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

package linters

import (
	"k8s.io/gengo/types"
)

type optionalFieldAPIConvention struct{}

func (c optionalFieldAPIConvention) Name() string {
	return "optional_fields_linter_pointers"
}

func (c optionalFieldAPIConvention) Validate(t *types.Type) ([]string, error) {
	violationIDs := make([]string, 0)

	// Only validate struct type and ignore the rest
	switch t.Kind {
	case types.Struct:
		for _, m := range t.Members {
			if hasAPITagValue(m.CommentLines, tagValueFalse) {
				continue
			}
			// Get underlying member type for alias type
			mt := m.Type
			for mt.Kind == types.Alias {
				mt = mt.Underlying
			}
			// Optional fields should be nillable (pointer, map or slice)
			if hasOptionalTag(&m) && !(mt.Kind == types.Pointer || mt.Kind == types.Map || mt.Kind == types.Slice) {
				violationIDs = append(violationIDs, m.Name)
			}
		}
	}
	return violationIDs, nil
}
