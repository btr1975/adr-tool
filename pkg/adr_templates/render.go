package adr_templates

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

// GetDateString returns a date string based on the dateType
//
//	:param dateType: The type of date to return.
//
// Example:
//
//	date := GetDateString("long")
func GetDateString(dateType string) (date string) {
	now := time.Now()

	switch dateType {
	case "long":
		return fmt.Sprintf("%v", now.Format("Mon Jan 2 15:04:05 MST 2006"))
	case "short":
		return fmt.Sprintf("%v", now.Format("2006-01-02"))
	default:
		return fmt.Sprintf("%v", now.Format("01/02/2006"))
	}
}

// ShortTemplate is a template with a Title, a Statement, and a list of Options.
type ShortTemplate struct {
	Title       string
	Date        string
	Statement   string
	Options     []string
	FileName    string
	useTemplate string
}

// NewShortTemplate returns a new ShortTemplate with the given title, statement, and options.
//
//	:param title: The title of the ADR.
//	:param statement: The statement of the ADR.
//	:param options: The options of the ADR.
//	:param structurizr: Structurizr compatible ADR.
//
// Example:
//
//	template := NewShortTemplate("My Title", "My Statement", []string{"Option 1", "Option 2"})
func NewShortTemplate(title string, statement string, options []string, structurizr bool) *ShortTemplate {
	if structurizr {
		return &ShortTemplate{
			Title:       title,
			Date:        GetDateString("short"),
			Statement:   statement,
			Options:     options,
			FileName:    fmt.Sprintf("%v.md", strings.Join(strings.Split(strings.ToLower(title), " "), "-")),
			useTemplate: "short_structurizr.template",
		}
	}

	return &ShortTemplate{
		Title:       title,
		Date:        GetDateString("long"),
		Statement:   statement,
		Options:     options,
		FileName:    fmt.Sprintf("%v.md", strings.Join(strings.Split(strings.ToLower(title), " "), "-")),
		useTemplate: "short.template",
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
	tmpl, err := GetTemplate(t.useTemplate)

	if err != nil {
		return "", err
	}

	bytesBuffer := bytes.Buffer{}

	err = tmpl.ExecuteTemplate(&bytesBuffer, t.useTemplate, t)

	if err != nil {
		return "", err
	}

	return bytesBuffer.String(), nil
}

// LongTemplate is a template with a Title, Deciders, Statement and a list of Options.
type LongTemplate struct {
	Title       string
	Deciders    string
	Date        string
	Statement   string
	Options     []string
	FileName    string
	useTemplate string
}

// NewLongTemplate returns a new LongTemplate with the given title, deciders, statement, and options.
//
//	:param title: The title of the ADR.
//	:param deciders: The deciders of the ADR.
//	:param statement: The statement of the ADR.
//	:param options: The options of the ADR.
//	:param structurizr: Structurizr compatible ADR.
//
// Example:
//
//	template := NewLongTemplate("My Title", "My Deciders", "My Statement", []string{"Option 1", "Option 2"})
func NewLongTemplate(title string, deciders string, statement string, options []string, structurizr bool) *LongTemplate {
	if structurizr {
		return &LongTemplate{
			Title:       title,
			Deciders:    deciders,
			Date:        GetDateString("short"),
			Statement:   statement,
			Options:     options,
			FileName:    fmt.Sprintf("%v.md", strings.Join(strings.Split(strings.ToLower(title), " "), "-")),
			useTemplate: "long_structurizr.template",
		}
	}

	return &LongTemplate{
		Title:       title,
		Deciders:    deciders,
		Date:        GetDateString("long"),
		Statement:   statement,
		Options:     options,
		FileName:    fmt.Sprintf("%v.md", strings.Join(strings.Split(strings.ToLower(title), " "), "-")),
		useTemplate: "long.template",
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
	tmpl, err := GetTemplate(t.useTemplate)

	if err != nil {
		return "", err
	}

	bytesBuffer := bytes.Buffer{}

	err = tmpl.ExecuteTemplate(&bytesBuffer, t.useTemplate, t)

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
