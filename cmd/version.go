package cmd

import (
	"fmt"

	"github.com/JayceChant/drivel/common/global"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of drivel",
	Long:  `Print the version of drivel.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s v%s\n", cmd.Parent().Name(), global.CurrentVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
