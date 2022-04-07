package cmd

import (
	"fmt"
	"strconv"
	"task/lib"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Long:  `task [do] - mark a task on your TODO list as complete`,
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			lib.CheckError(err)
			ids = append(ids, id)
		}
		tasks, err := lib.ListDataFromDatabase()
		lib.CheckError(err)
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := tasks[id-1]
			err := lib.DeleteDataFromDatabase(task.Id)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("✅ Marked \"%d\" as completed.\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
