package cmd

import (
	// "fmt"
	// "task/lib"

	"fmt"
	"task/lib"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Long:  `task [add] - mark a task on your TODO list as complete`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := lib.ListDataFromDatabase()
		lib.CheckError(err)
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete! Why not take a vacation? 🏖")
			return
		}
		fmt.Println("You have the following tasks :")
		for _, task := range tasks {
			fmt.Printf(
				"[%d]. %s [%s ❌ ]\n",
				task.Id,
				task.Description,
				task.GetStatusString())
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
