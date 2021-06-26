package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download and activate a binary release",
	Long: `EXAMPLE:
	  > brm splicectl v1.2.4`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("download called")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
