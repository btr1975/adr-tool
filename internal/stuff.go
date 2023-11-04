package main

import (
	"github.com/btr1975/adr-tool/pkg/adr_templates"
	"github.com/btr1975/adr-tool/pkg/records"
)

func main() {
	template := adr_templates.NewShortTemplate("My Title", "My Statement", []string{"Option 1", "Option 2"})

	err := records.SupersedeADR("./temp", template, "0001-my-title.md")

	if err != nil {
		panic(err)
	}
}
