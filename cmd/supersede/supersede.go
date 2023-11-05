/*
Package supersede implements the supersede command
*/
package supersede

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Cmd represents the Cmd command
var Cmd = &cobra.Command{
	Use:   "supersede",
	Short: "Supersede an ADR",
	Long:  `Supersede an ADR`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("supersede called")
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
