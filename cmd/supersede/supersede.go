/*
Package supersede implements the supersede command
*/
package supersede

import (
	"github.com/spf13/cobra"
	"os"
)

var deciders string
var path string
var title string
var statement string
var options []string
var adr string
var structurizr bool

// Cmd represents the Cmd command
var Cmd = &cobra.Command{
	Use:   "supersede",
	Short: "Supersede an ADR",
	Long: `Supersede an ADR

Example usage:
	adr-tool supersede long -p ./dir -a 0001-some-adr.md -t "Some Title" -s "Statement of Decision" -o "opt 1" -o "opt 2" -o "opt 3" -d "John,Phil,Tom"
	adr-tool supersede short -p ./dir -a 0001-some-adr.md -t "Some Title" -s "Statement of Decision" -o "opt 1" -o "opt 2" -o "opt 3"
	adr-tool supersede short -p ./dir -a 0001-some-adr.md -t "Some Title" -s "Statement of Decision"
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

// addSubCommands adds all the subcommands to the supersede command
func addSubCommands() {
	Cmd.AddCommand(longCmd)
	Cmd.AddCommand(shortCmd)
}

func init() {
	addSubCommands()
}
