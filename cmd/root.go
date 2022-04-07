/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a CLI for managing your TODOs.",
	Long:  `A simple yet powerful task managing CLI application that uses BOLT DB for persistent storage of related task and their status`,
}
