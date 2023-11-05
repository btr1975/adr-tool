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

// StatusCmd represents the StatusCmd command
var StatusCmd = &cobra.Command{
	Use:   "change-status",
	Short: "Change ADR status",
	Long:  `Change ADR status`,
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
			fmt.Printf("ADR status changed: %v\n", adr)
		}
	},
}

func init() {
	StatusCmd.Flags().StringVarP(&path, "path", "p", "", "Path to the ADR directory")
	StatusCmd.Flags().StringVarP(&adr, "adr", "a", "", "ADR to change status of")
	StatusCmd.Flags().StringVarP(&status, "status", "s", "", "Status to change ADR to")

	if err := StatusCmd.MarkFlagRequired("path"); err != nil {
		fmt.Println(err)
	}

	if err := StatusCmd.MarkFlagRequired("adr"); err != nil {
		fmt.Println(err)
	}

	if err := StatusCmd.MarkFlagRequired("status"); err != nil {
		fmt.Println(err)
	}
}
