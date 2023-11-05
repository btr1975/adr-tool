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

// LongCmd represents the LongCmd command
var LongCmd = &cobra.Command{
	Use:   "long-adr",
	Short: "Create a long ADR",
	Long:  `Create a long ADR`,
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
	LongCmd.Flags().StringVarP(&path, "path", "p", "", "Path to the ADR directory")
	LongCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the ADR")
	LongCmd.Flags().StringVarP(&statement, "statement", "s", "", "Statement of the ADR")
	LongCmd.Flags().StringSliceVarP(&options, "options", "o", []string{}, "Options of the ADR")
	LongCmd.Flags().StringVarP(&deciders, "deciders", "d", "", "Deciders of the ADR")

	if err := LongCmd.MarkFlagRequired("path"); err != nil {
		fmt.Println(err)
	}

	if err := LongCmd.MarkFlagRequired("title"); err != nil {
		fmt.Println(err)
	}

	if err := LongCmd.MarkFlagRequired("statement"); err != nil {
		fmt.Println(err)
	}

	if err := LongCmd.MarkFlagRequired("options"); err != nil {
		fmt.Println(err)
	}

	if err := LongCmd.MarkFlagRequired("deciders"); err != nil {
		fmt.Println(err)
	}
}
