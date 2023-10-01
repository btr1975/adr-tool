/*
Package adr_templates provides templates for creating ADRs.
*/
package adr_templates

import (
	"embed"
	"fmt"
)

// templateFileSystem is the filesystem that contains template files.
//
//go:embed templates/*.template
var templateFileSystem embed.FS

// GetTemplate returns the template with the given name.
//
// Example:
//
//	template, err := GetTemplate("template.md")
func GetTemplate(name string) ([]byte, error) {
	return templateFileSystem.ReadFile(fmt.Sprintf("templates/%s", name))
}
