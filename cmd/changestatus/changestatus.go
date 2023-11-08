/*
Package changestatus implements the change-status command
*/
package changestatus

import (
	"fmt"
	"github.com/btr1975/adr-tool/pkg/records"
	"os"

	"github.com/spf13/cobra"
)

var path string
var adr string
var status string

// Cmd represents the Cmd command
var Cmd = &cobra.Command{
	Use:   "change-status",
	Short: "Change ADR status",
	Long: `Change ADR status

Example usage:
	adr-tool change-status -p ./dir -a 0001-some-adr.md -s accepted
`,
	Run: func(cmd *cobra.Command, args []string) {
		statusType, err := records.StringToStatus(status)

		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		err = records.ChangeADRStatus(path, adr, statusType, false)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Printf("ADR status changed to %s: %v\n", string(statusType), adr)
		}
	},
}

func init() {
	Cmd.Flags().StringVarP(&path, "path", "p", "", "Path to the ADR directory")
	Cmd.Flags().StringVarP(&adr, "adr", "a", "", "ADR to change status of")
	Cmd.Flags().StringVarP(&status, "status", "s", "", "Status to change ADR to")

	if err := Cmd.MarkFlagRequired("path"); err != nil {
		fmt.Println(err)
	}

	if err := Cmd.MarkFlagRequired("adr"); err != nil {
		fmt.Println(err)
	}

	if err := Cmd.MarkFlagRequired("status"); err != nil {
		fmt.Println(err)
	}
}
