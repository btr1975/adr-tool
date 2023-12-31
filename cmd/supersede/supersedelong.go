/*
Package supersede implements the supersede long command
*/
package supersede

import (
	"fmt"
	"github.com/btr1975/adr-tool/pkg/adr_templates"
	"github.com/btr1975/adr-tool/pkg/records"
	"os"

	"github.com/spf13/cobra"
)

// longCmd represents the longCmd command
var longCmd = &cobra.Command{
	Use:   "long",
	Short: "Supersede with long ADR",
	Long: `Supersede with long ADR

Example usage:
	adr-tool supersede long -p ./dir -a 0001-some-adr.md -t "Some Title" -s "Statement of Decision" -o "opt 1" -o "opt 2" -o "opt 3" -d "John,Phil,Tom"
`,
	Run: func(cmd *cobra.Command, args []string) {
		template := adr_templates.NewLongTemplate(title, deciders, statement, options, structurizr)

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
	longCmd.Flags().StringVarP(&adr, "adr", "a", "", "ADR to supersede")
	longCmd.Flags().StringVarP(&path, "path", "p", "", "Path to the ADR directory")
	longCmd.Flags().StringVarP(&title, "title", "t", "", "Title of the ADR")
	longCmd.Flags().StringVarP(&statement, "statement", "s", "", "Statement of the ADR")
	longCmd.Flags().StringSliceVarP(&options, "options", "o", []string{}, "Options of the ADR")
	longCmd.Flags().StringVarP(&deciders, "deciders", "d", "", "Deciders of the ADR")
	longCmd.Flags().BoolVarP(&structurizr, "structurizr-compat", "c", false, "Structurizr Compatible ADR")

	if err := longCmd.MarkFlagRequired("adr"); err != nil {
		fmt.Println(err)
	}

	if err := longCmd.MarkFlagRequired("path"); err != nil {
		fmt.Println(err)
	}

	if err := longCmd.MarkFlagRequired("title"); err != nil {
		fmt.Println(err)
	}

	if err := longCmd.MarkFlagRequired("statement"); err != nil {
		fmt.Println(err)
	}

	if err := longCmd.MarkFlagRequired("options"); err != nil {
		fmt.Println(err)
	}

	if err := longCmd.MarkFlagRequired("deciders"); err != nil {
		fmt.Println(err)
	}

}
