package main

import (
	"fmt"
	"github.com/btr1975/adr-tool/pkg/adr_templates"
)

func main() {
	thing, err := adr_templates.NewShortTemplate("This is a great title", "I have some shit to say", []string{"mongo", "mysql"}).Render()

	if err != nil {
		panic(err)
	}

	fmt.Println(thing)
}
