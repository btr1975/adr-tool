/*
Package supersede implements the supersede long command
*/
package supersede

import (
	"fmt"

	"github.com/spf13/cobra"
)

// longCmd represents the longCmd command
var longCmd = &cobra.Command{
	Use:   "long",
	Short: "Supersede with long ADR",
	Long:  `Supersede with long ADR`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("long called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// supersedelongCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// supersedelongCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
