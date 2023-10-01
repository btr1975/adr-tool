package adr_templates

import "bytes"

// ShortTemplate is a template with a name, a statement, and a list of options.
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

func (t *ShortTemplate) Render() (rendering string, err error) {
	tmpl, err := GetTemplate("short.template")

	if err != nil {
		return "", err
	}

	bytesBuffer := bytes.Buffer{}

	err = tmpl.ExecuteTemplate(&bytesBuffer, "short.template", t)

	if err != nil {
		return "", err
	}

	return bytesBuffer.String(), nil
}

// LongTemplate is a template with a name, a statement, and a list of options.
type LongTemplate struct {
	Title     string
	Deciders  string
	Date      string
	Statement string
	Options   []string
}
