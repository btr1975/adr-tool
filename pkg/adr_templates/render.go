package adr_templates

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

// ShortTemplate is a template with a Title, a Statement, and a list of Options.
type ShortTemplate struct {
	Title     string
	Date      string
	Statement string
	Options   []string
	FileName  string
}

// NewShortTemplate returns a new ShortTemplate with the given title, statement, and options.
//
// Example:
//
//	template := NewShortTemplate("My Title", "My Statement", []string{"Option 1", "Option 2"})
func NewShortTemplate(title string, statement string, options []string) *ShortTemplate {
	now := time.Now()

	return &ShortTemplate{
		Title:     title,
		Date:      fmt.Sprintf("%v", now.Format("Mon Jan 2 15:04:05 MST 2006")),
		Statement: statement,
		Options:   options,
		FileName:  fmt.Sprintf("%v.md", strings.Join(strings.Split(strings.ToLower(title), " "), "-")),
	}
}

// GetFileName returns the file name of the template
//
// Example:
//
//	name := GetFileName()
func (t *ShortTemplate) GetFileName() (name string) {
	return t.FileName
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
	FileName  string
}

// NewLongTemplate returns a new LongTemplate with the given title, deciders, statement, and options.
//
// Example:
//
//	template := NewLongTemplate("My Title", "My Deciders", "My Statement", []string{"Option 1", "Option 2"})
func NewLongTemplate(title string, deciders string, statement string, options []string) *LongTemplate {
	now := time.Now()

	return &LongTemplate{
		Title:     title,
		Deciders:  deciders,
		Date:      fmt.Sprintf("%v", now.Format("Mon Jan 2 15:04:05 MST 2006")),
		Statement: statement,
		Options:   options,
		FileName:  fmt.Sprintf("%v.md", strings.Join(strings.Split(strings.ToLower(title), " "), "-")),
	}
}

// GetFileName returns the file name of the template
//
// Example:
//
//	name := GetFileName()
func (t *LongTemplate) GetFileName() (name string) {
	return t.FileName
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

// RenderTemplate is an interface for rendering templates.
type RenderTemplate interface {
	Render() (string, error)
	GetFileName() string
}
