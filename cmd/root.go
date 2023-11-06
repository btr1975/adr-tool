/*
Package cmd implements the command line for the adr-tool
*/
package cmd

import (
	"github.com/btr1975/adr-tool/cmd/changestatus"
	"github.com/btr1975/adr-tool/cmd/longadr"
	"github.com/btr1975/adr-tool/cmd/shortadr"
	"github.com/btr1975/adr-tool/cmd/supersede"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

var versionFlag bool
var Author = "Benjamin P. Trachtenberg"
var Version = "0.0.1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "adr-tool",
	Short: "A Simple tool for managing Architecture Decision Records",
	Long: `A Simple tool for managing Architecture Decision Records

Example usage:
	adr-tool change-status
	adr-tool long-adr
	adr-tool short-adr
	adr-tool supersede
`,
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			cmd.Printf("Author: %v\n", Author)
			cmd.Printf("Version: %v\n", Version)
			cmd.Printf("OS: %v\n", runtime.GOOS)
			cmd.Printf("Arch: %v\n", runtime.GOARCH)
			os.Exit(0)
		} else {
			cmd.Help()
			os.Exit(0)
		}
	},
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
	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "Version")

	addSubCommands()
}
