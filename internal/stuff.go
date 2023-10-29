package main

import (
	"github.com/btr1975/adr-tool/pkg/adr_templates"
	"github.com/btr1975/adr-tool/pkg/records"
)

func main() {
	thing := adr_templates.NewLongTemplate("My Title", "My Deciders", "My Statement", []string{"Option 1", "Option 2"})

	err := records.WriteNewADR("./temp", thing)

	if err != nil {
		panic(err)
	}

	records.ReadADRDirectory("./")
}
