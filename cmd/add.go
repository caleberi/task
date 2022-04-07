package cmd

import (
	"fmt"
	"strings"
	"task/lib"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Long:  `task [add] - add a new task to your TODO list`,
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := strings.Join(args, " ")
		task := lib.CreateTask(description, lib.OPEN)
		_, err := lib.InsertInDatabase(task)
		lib.CheckError(err)
		fmt.Printf("🌶️ Added \"%v\" to your task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
