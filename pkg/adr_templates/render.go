package adr_templates

import (
	"bytes"
	"fmt"
	"time"
)

// ShortTemplate is a template with a Title, a Statement, and a list of Options.
type ShortTemplate struct {
	Title     string
	Statement string
	Options   []string
}

// NewShortTemplate returns a new ShortTemplate with the given title, statement, and options.
//
// Example:
//
//	template := NewShortTemplate("My Title", "My Statement", []string{"Option 1", "Option 2"})
func NewShortTemplate(title string, statement string, options []string) *ShortTemplate {
	return &ShortTemplate{
		Title:     title,
		Statement: statement,
		Options:   options,
	}
}

// Render renders the short template
//
// Example:
//
//	rendering, err := Render()
func (t *ShortTemplate) Render() (rendering string, err error) {
	templateName := "short.template"

	tmpl, err := GetTemplate(templateName)

	if err != nil {
		return "", err
	}

	bytesBuffer := bytes.Buffer{}

	err = tmpl.ExecuteTemplate(&bytesBuffer, templateName, t)

	if err != nil {
		return "", err
	}

	return bytesBuffer.String(), nil
}

// LongTemplate is a template with a Title, Deciders, Statement and a list of Options.
type LongTemplate struct {
	Title     string
	Deciders  string
	Date      string
	Statement string
	Options   []string
}

func NewLongTemplate(title string, deciders string, statement string, options []string) *LongTemplate {
	now := time.Now()

	return &LongTemplate{
		Title:     title,
		Deciders:  deciders,
		Date:      fmt.Sprintf("%v", now.Format("Mon Jan 2 15:04:05 MST 2006")),
		Statement: statement,
		Options:   options,
	}
}

// Render renders the long template
//
// Example:
//
//	rendering, err := Render()
func (t *LongTemplate) Render() (rendering string, err error) {
	templateName := "long.template"

	tmpl, err := GetTemplate(templateName)

	if err != nil {
		return "", err
	}

	bytesBuffer := bytes.Buffer{}

	err = tmpl.ExecuteTemplate(&bytesBuffer, templateName, t)

	if err != nil {
		return "", err
	}

	return bytesBuffer.String(), nil
}
