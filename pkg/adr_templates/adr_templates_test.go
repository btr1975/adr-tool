package adr_templates

import (
	"fmt"
	"testing"
)

func TestGetTemplate(t *testing.T) {
	scenarios := []struct {
		name string
		err  error
	}{
		{
			name: "short.template",
			err:  nil,
		},
		{
			name: "long.template",
			err:  nil,
		},
		{
			name: "template.md",
			err:  fmt.Errorf("template: \"template.md\" is an incomplete or empty template"),
		},
	}

	for _, scenario := range scenarios {
		_, err := GetTemplate(scenario.name)
		if err != nil && scenario.err == nil {
			t.Error("ADR template is empty")
		} else if err == nil && scenario.err != nil {
			t.Error("ADR template is not empty")
		} else if err == nil && scenario.err == nil {
			t.Log("ADR template is not empty")
		} else if err != nil && scenario.err != nil {
			t.Log("ADR template is empty")
		} else {
			t.Error("ADR template is in an unknown state")
		}
	}
}
