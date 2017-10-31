package linters

import (
	"k8s.io/gengo/types"
)

type optionalFieldAPIConvention struct{}

func (c optionalFieldAPIConvention) Name() string {
	return "optional_fields_linter_pointers"
}

func (c optionalFieldAPIConvention) Message() string {
	return "optional fields should be pointers"
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
			if hasOptionalTag(&m) && m.Type.Kind != types.Pointer {
				violationIDs = append(violationIDs, m.Name)
			}
		}
	}
	return violationIDs, nil
}
