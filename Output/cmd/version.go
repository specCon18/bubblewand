package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of test-test",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test-test ~ 0.0.6")
	},
}
