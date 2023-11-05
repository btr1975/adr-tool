/*
Package supersede implements the supersede short command
*/
package supersede

import (
	"fmt"
	"github.com/btr1975/adr-tool/pkg/adr_templates"
	"github.com/btr1975/adr-tool/pkg/records"
	"os"

	"github.com/spf13/cobra"
)

// shortCmd represents the shortCmd command
var shortCmd = &cobra.Command{
	Use:   "short",
	Short: "Supersede with short ADR",
	Long:  `Supersede with short ADR`,
	Run: func(cmd *cobra.Command, args []string) {
		template := adr_templates.NewShortTemplate(title, statement, options)

		fileName, err := records.SupersedeADR(path, template, adr)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Printf("ADR created: %v Supersedes %v \n", fileName, adr)
		}
	},
}

func init() {
	shortCmd.Flags().StringVarP(&adr, "adr", "a", "", "ADR to supersede")
	shortCmd.Flags().StringVarP(&path, "path", "p", "", "Path to the ADR directory")
	shortCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the ADR")
	shortCmd.Flags().StringVarP(&statement, "statement", "s", "", "Statement of the ADR")
	shortCmd.Flags().StringSliceVarP(&options, "options", "o", []string{}, "Options of the ADR")

	if err := shortCmd.MarkFlagRequired("adr"); err != nil {
		fmt.Println(err)
	}

	if err := shortCmd.MarkFlagRequired("path"); err != nil {
		fmt.Println(err)
	}

	if err := shortCmd.MarkFlagRequired("title"); err != nil {
		fmt.Println(err)
	}

	if err := shortCmd.MarkFlagRequired("statement"); err != nil {
		fmt.Println(err)
	}

	if err := shortCmd.MarkFlagRequired("options"); err != nil {
		fmt.Println(err)
	}
}
