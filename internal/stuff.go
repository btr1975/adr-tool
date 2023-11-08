package main

import (
	"fmt"
	"github.com/btr1975/adr-tool/pkg/adr_templates"
)

func main() {
	template := adr_templates.NewShortTemplate("Some Title", "Statement of Decision", []string{"opt 1", "opt 2", "opt 3"})

	rendering, err := template.Render()

	if err != nil {
		panic(err)
	}

	fmt.Println(rendering)
}
