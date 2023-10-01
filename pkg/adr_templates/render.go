package adr_templates

import (
	"fmt"
)

// ShortTemplate is a template with a name, a statement, and a list of options.
type ShortTemplate struct {
	Title     string
	Statement string
	Options   []string
}

func NewShortTemplate(title string, statement string, options []string) *ShortTemplate {
	return &ShortTemplate{
		Title:     title,
		Statement: statement,
		Options:   options,
	}
}

func (t *ShortTemplate) Render() string {
	tmpl, err := GetTemplate("short.template")

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf(string(tmpl), t.Title, t.Statement, t.Options)
}
