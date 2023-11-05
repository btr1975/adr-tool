/*
Package cmd implements the command line for the adr-tool
*/
package cmd

import (
	"github.com/btr1975/adr-tool/cmd/changestatus"
	"github.com/btr1975/adr-tool/cmd/longadr"
	"github.com/btr1975/adr-tool/cmd/shortadr"
	"github.com/btr1975/adr-tool/cmd/supersede"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "adr-tool",
	Short: "A Simple tool for managing Architecture Decision Records",
	Long: `A Simple tool for managing Architecture Decision Records
Example usage:
	adr-tool short-adr -p ./dir -t "Some Title" -s "Statement of Decision" -o "opt 1" -o "opt 2" -o "opt 3"
	adr-tool long-adr -p ./dir -t "Some Title" -s "Statement of Decision" -o "opt 1" -o "opt 2" -o "opt 3" -d "John,Phil,Tom"
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// addSubCommands adds all the subcommands to the root command
func addSubCommands() {
	rootCmd.AddCommand(shortadr.Cmd)
	rootCmd.AddCommand(longadr.Cmd)
	rootCmd.AddCommand(changestatus.Cmd)
	rootCmd.AddCommand(supersede.Cmd)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.adr-tool.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	addSubCommands()
}
