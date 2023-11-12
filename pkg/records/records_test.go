package records

import (
	"fmt"
	"github.com/btr1975/adr-tool/pkg/adr_templates"
	"os"
	"path"
	"testing"
)

func TestStringToStatus(t *testing.T) {
	scenarios := []struct {
		name     string
		input    string
		expected Status
		err      bool
	}{
		{
			name:     "proposed",
			input:    "proposed",
			expected: Proposed,
			err:      false,
		},
		{
			name:     "accepted",
			input:    "accepted",
			expected: Accepted,
			err:      false,
		},
		{
			name:     "rejected",
			input:    "rejected",
			expected: Rejected,
			err:      false,
		},
		{
			name:     "deprecated",
			input:    "deprecated",
			expected: Deprecated,
			err:      false,
		},
		{
			name:     "superseded",
			input:    "superseded",
			expected: Superseded,
			err:      false,
		},
		{
			name:     "invalid",
			input:    "invalid",
			expected: "",
			err:      true,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			actual, err := StringToStatus(scenario.input)
			if err != nil {
				if !scenario.err {
					t.Errorf("expected no error, got %v", err)
				} else if scenario.err {
					t.Logf("expected error, got %v", err)
				}
			}
			if actual != scenario.expected {
				t.Errorf("expected %s, got %s", scenario.expected, actual)
			}
		})
	}
}

func TestFileExists(t *testing.T) {
	goodPath := path.Join("..", "..", "test", "long", "my-title.md")
	badPath := path.Join("..", "..", "test", "long", "my-title.txt")

	scenarios := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "file exists",
			input:    goodPath,
			expected: true,
		},
		{
			name:     "file does not exist",
			input:    badPath,
			expected: false,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			actual := FileExists(scenario.input)
			if actual != scenario.expected {
				t.Errorf("expected %t, got %t", scenario.expected, actual)
			}
		})
	}
}

func TestDirectoryExists(t *testing.T) {
	goodPath := path.Join("..", "..", "test", "long")
	badPath := path.Join("..", "..", "test", "unknown")

	scenarios := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "directory exists",
			input:    goodPath,
			expected: true,
		},
		{
			name:     "directory does not exist",
			input:    badPath,
			expected: false,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			actual := DirectoryExists(scenario.input)
			if actual != scenario.expected {
				t.Errorf("expected %t, got %t", scenario.expected, actual)
			}
		})
	}
}

func TestWriteNewADR(t *testing.T) {
	scenarios := []struct {
		name      string
		path      string
		title     string
		statement string
		options   []string
	}{
		{
			name:      "write new ADR",
			path:      path.Join("..", "..", "test", "long"),
			title:     "My Title",
			statement: "My Statement",
			options:   []string{"Option 1", "Option 2"},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			template := adr_templates.NewShortTemplate(scenario.title, scenario.statement, scenario.options, false)
			fileName, err := WriteNewADR(scenario.path, template)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			createdFileName := fmt.Sprintf("0001-%s", template.GetFileName())
			if fileName != createdFileName {
				t.Errorf("expected %s, got %s", createdFileName, fileName)
			}
			if !FileExists(path.Join(scenario.path, fileName)) {
				t.Errorf("expected file to exist, got %s", fileName)
			}
			_ = os.Remove(path.Join(scenario.path, fileName))
		})

	}
}

func TestSupersedeADR(t *testing.T) {
	scenarios := []struct {
		name      string
		path      string
		title     string
		statement string
		options   []string
		adr       string
	}{
		{
			name:      "supersede ADR",
			path:      path.Join("..", "..", "test", "supersede"),
			title:     "My Title",
			statement: "My Statement",
			options:   []string{"Option 1", "Option 2"},
			adr:       "0001-my-title.md",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			template := adr_templates.NewShortTemplate(scenario.title, scenario.statement, scenario.options, false)
			fileName, err := SupersedeADR(scenario.path, template, scenario.adr)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			createdFileName := fmt.Sprintf("0003-%s", template.GetFileName())
			if fileName != createdFileName {
				t.Errorf("expected %s, got %s", createdFileName, fileName)
			}
			if !FileExists(path.Join(scenario.path, fileName)) {
				t.Errorf("expected file to exist, got %s", fileName)
			}
			_ = os.Remove(path.Join(scenario.path, fileName))
		})
	}
}
