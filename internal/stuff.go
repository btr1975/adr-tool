package main

import (
	"fmt"
	"github.com/btr1975/adr-tool/pkg/adr_templates"
)

func main() {
	thing, err := adr_templates.NewLongTemplate("My Title", "My Deciders", "My Statement", []string{"Option 1", "Option 2"}).Render()

	if err != nil {
		panic(err)
	}

	fmt.Println(thing)
}
