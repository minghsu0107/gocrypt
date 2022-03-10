package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version",
	Long:  `Print the current version of Gocrypt`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}
