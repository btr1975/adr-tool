/*
Package longadr implements the long-adr command
*/
package longadr

import (
	"fmt"
	"github.com/btr1975/adr-tool/pkg/adr_templates"
	"github.com/btr1975/adr-tool/pkg/records"
	"os"

	"github.com/spf13/cobra"
)

var path string
var title string
var statement string
var options []string
var deciders string

// Cmd represents the Cmd command
var Cmd = &cobra.Command{
	Use:   "long-adr",
	Short: "Create a long ADR",
	Long: `Create a long ADR

Example usage:
	adr-tool long-adr -p ./dir -t "Some Title" -s "Statement of Decision" -o "opt 1" -o "opt 2" -o "opt 3" -d "John,Phil,Tom"
`,
	Run: func(cmd *cobra.Command, args []string) {
		template := adr_templates.NewLongTemplate(title, deciders, statement, options)

		fileName, err := records.WriteNewADR(path, template)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Printf("ADR created: %v\n", fileName)
		}
	},
}

func init() {
	Cmd.Flags().StringVarP(&path, "path", "p", "", "Path to the ADR directory")
	Cmd.Flags().StringVarP(&title, "title", "t", "", "Title of the ADR")
	Cmd.Flags().StringVarP(&statement, "statement", "s", "", "Statement of the ADR")
	Cmd.Flags().StringSliceVarP(&options, "options", "o", []string{}, "Options of the ADR")
	Cmd.Flags().StringVarP(&deciders, "deciders", "d", "", "Deciders of the ADR")

	if err := Cmd.MarkFlagRequired("path"); err != nil {
		fmt.Println(err)
	}

	if err := Cmd.MarkFlagRequired("title"); err != nil {
		fmt.Println(err)
	}

	if err := Cmd.MarkFlagRequired("statement"); err != nil {
		fmt.Println(err)
	}

	if err := Cmd.MarkFlagRequired("options"); err != nil {
		fmt.Println(err)
	}

	if err := Cmd.MarkFlagRequired("deciders"); err != nil {
		fmt.Println(err)
	}
}
