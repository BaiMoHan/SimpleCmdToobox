package cmd

import "github.com/spf13/cobra"

// 放置根命令

var rootCmd =&cobra.Command{}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// 完成相应子命令的注册
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(sqlCmd)
}