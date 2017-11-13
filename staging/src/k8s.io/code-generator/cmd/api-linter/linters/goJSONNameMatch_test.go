package linters

import "testing"

func TestGoJSONNameMatch(t *testing.T) {
	tables := []struct {
		goName   string
		jsonName string
		match    bool
	}{
		{"PodSpec", "podSpec", true},
		{"PodSpec", "spec", false},
		{"Spec", "podSpec", false},
		{"JSONSpec", "jsonSpec", true},
		{"JSONSpec", "jsonspec", true},
	}

	for _, table := range tables {
		match := goJSONNameMatch(table.goName, table.jsonName)
		if match != table.match {
			t.Errorf("unexpected match result: goName %v, jsonName %v, want: %v, got: %v", table.goName, table.jsonName, table.match, match)
		}
	}
}
