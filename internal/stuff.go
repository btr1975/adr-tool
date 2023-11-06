package main

import (
	"fmt"
	"github.com/btr1975/adr-tool/pkg/adr_templates"
	"os"
)

func main() {
	template := adr_templates.NewLongTemplate("My Title", "John Doe, Jane Doe", "My Statement", []string{"Option 1", "Option 2"})

	rendered, err := template.Render()

	if err != nil {
		panic(err)
	}

	fullPath := fmt.Sprintf("test/long/%v", template.GetFileName())

	err = os.WriteFile(fullPath, []byte(rendered), 0644)

	if err != nil {
		panic(err)
	}
}
