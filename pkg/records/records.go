package records

import (
	"github.com/btr1975/adr-tool/pkg/adr_templates"
	"os"
	"regexp"
)

func WriteADR(path string, template adr_templates.RenderTemplate) (err error) {
	render, err := template.Render()

	if err != nil {
		return err
	}

	err = os.WriteFile(path, []byte(render), 0644)

	return err
}

func ReadADRDirectory(path string) {
	thing, err := os.ReadDir(path)

	if err != nil {
		panic(err)
	}

	regexp.Compile(`^\d{5}-`)

	for _, entry := range thing {
		if entry.IsDir() != true {
			println(entry.Name())
		}
	}

}
