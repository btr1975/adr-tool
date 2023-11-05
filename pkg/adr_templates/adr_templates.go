/*
Package adr_templates provides templates for creating ADRs.
*/
package adr_templates

import (
	"embed"
	"fmt"
	"html/template"
)

// templateFileSystem is the filesystem that contains template files.
//
//go:embed templates/*.template
var templateFileSystem embed.FS

// GetTemplate returns the template with the given name.
//
//	:param name: The name of the template to return.
//
// Example:
//
//	template, err := GetTemplate("template.md")
func GetTemplate(name string) (template *template.Template, err error) {
	return template.ParseFS(templateFileSystem, fmt.Sprintf("templates/%s", name))
}
