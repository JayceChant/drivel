package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of drivel",
	Long:  `Print the version of drivel.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("drivel v0.1.0") // TODO: change to read version from file
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
