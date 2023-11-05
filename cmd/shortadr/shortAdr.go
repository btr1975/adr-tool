/*
Package shortadr implements the command to create a short ADR
*/
package shortadr

import (
	"fmt"
	"github.com/btr1975/adr-tool/pkg/adr_templates"
	"github.com/btr1975/adr-tool/pkg/records"

	"github.com/spf13/cobra"
)

var path string
var title string
var statement string
var options []string

// ShortCmd represents the ShortCmd command
var ShortCmd = &cobra.Command{
	Use:   "short-adr",
	Short: "Create a short ADR",
	Long:  `Create a short ADR`,
	Run: func(cmd *cobra.Command, args []string) {
		template := adr_templates.NewShortTemplate(title, statement, options)

		fileName, err := records.WriteNewADR(path, template)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("ADR created: %v\n", fileName)
		}

	},
}

func init() {
	ShortCmd.Flags().StringVarP(&path, "path", "p", "", "Path to the ADR directory")
	ShortCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the ADR")
	ShortCmd.Flags().StringVarP(&statement, "statement", "s", "", "Statement of the ADR")
	ShortCmd.Flags().StringSliceVarP(&options, "options", "o", []string{}, "Options of the ADR")

	if err := ShortCmd.MarkFlagRequired("path"); err != nil {
		fmt.Println(err)
	}

	if err := ShortCmd.MarkFlagRequired("title"); err != nil {
		fmt.Println(err)
	}

	if err := ShortCmd.MarkFlagRequired("statement"); err != nil {
		fmt.Println(err)
	}

	if err := ShortCmd.MarkFlagRequired("options"); err != nil {
		fmt.Println(err)
	}
}
