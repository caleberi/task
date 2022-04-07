/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"path/filepath"
	"task/cmd"
	"task/lib"

	"github.com/mitchellh/go-homedir"
)

var bucketName string = "TaskDB"

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	db, err := lib.InitDatabase(dbPath, bucketName)
	if err != nil {
		panic(err)
	}
	cmd.RootCmd.Execute()
	lib.CloseDatabase(db)
}
