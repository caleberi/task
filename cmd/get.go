/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"task/lib"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get just list of task by provided ids",
	Long:  `Retrieve a single created task 🗼`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			iarg, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Invalid Id :{%s}", arg)
				continue
			}
			ids = append(ids, iarg)
		}
		var tasks []string
		for _, id := range ids {
			task, err := lib.RetrieveFromDatabase(id)
			if err != nil {
				fmt.Printf("Error retrieving task with Id [%d]❕", id)
				continue
			}
			tasks = append(tasks, task)
		}

		fmt.Println("You have the following tasks :")
		for _, task := range tasks {
			fmt.Printf("🏹 : %s \n", task)
		}
		fmt.Println()

	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
