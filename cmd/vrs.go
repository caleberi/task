package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of TaskCLI",
	Long:  `All software has versions. This is TaskCLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TaskCLI v0.1 -- HEAD")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
