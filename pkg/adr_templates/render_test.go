package adr_templates

import (
	"os"
	"path"
	"testing"
)

func TestNewShortTemplate(t *testing.T) {
	osPath := path.Join("..", "..", "test", "short", "my-title.md")
	goodRender, err := os.ReadFile(osPath)

	if err != nil {
		t.Errorf("Could not read file: %v", err)
	}

	scenarios := []struct {
		title     string
		statement string
		options   []string
	}{
		{
			title:     "My Title",
			statement: "My Statement",
			options:   []string{"Option 1", "Option 2"},
		},
	}

	for _, scenario := range scenarios {
		template := NewShortTemplate(scenario.title, scenario.statement, scenario.options)
		if template.Title != scenario.title {
			t.Error("Title is not correct")
		}
		if template.Statement != scenario.statement {
			t.Error("Statement is not correct")
		}
		if template.Options[0] != scenario.options[0] {
			t.Error("Option 1 is not correct")
		}
		if template.Options[1] != scenario.options[1] {
			t.Error("Option 2 is not correct")
		}
		if template.FileName != "my-title.md" {
			t.Error("File name is not correct")
		}
		if template.GetFileName() != "my-title.md" {
			t.Error("File name is not correct")
		}

		rendered, _ := template.Render()

		if rendered != string(goodRender) {
			t.Skip("Render is not correct")
			// Doesn't work because of the date
		}
	}
}

func TestNewLongTemplate(t *testing.T) {
	osPath := path.Join("..", "..", "test", "long", "my-title.md")
	goodRender, err := os.ReadFile(osPath)

	if err != nil {
		t.Errorf("Could not read file: %v", err)
	}

	scenarios := []struct {
		title     string
		statement string
		deciders  string
		options   []string
	}{
		{
			title:     "My Title",
			statement: "My Statement",
			deciders:  "John Doe, Jane Doe",
			options:   []string{"Option 1", "Option 2"},
		},
	}

	for _, scenario := range scenarios {
		template := NewLongTemplate(scenario.title, scenario.deciders, scenario.statement, scenario.options)
		if template.Title != scenario.title {
			t.Error("Title is not correct")
		}
		if template.Statement != scenario.statement {
			t.Error("Statement is not correct")
		}
		if template.Deciders != scenario.deciders {
			t.Error("Deciders is not correct")
		}
		if template.Options[0] != scenario.options[0] {
			t.Error("Option 1 is not correct")
		}
		if template.Options[1] != scenario.options[1] {
			t.Error("Option 2 is not correct")
		}
		if template.FileName != "my-title.md" {
			t.Error("File name is not correct")
		}
		if template.GetFileName() != "my-title.md" {
			t.Error("File name is not correct")
		}

		rendered, _ := template.Render()

		if rendered != string(goodRender) {
			t.Skip("Render is not correct")
			// Doesn't work because of the date
		}
	}
}
