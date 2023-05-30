/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/anandureghu/linex/helper"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list | list <command name>",
	Short: "list all commands or a specified command",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		mydir, _ := os.Getwd()
		fmt.Println("list called at", mydir)
		commands := helper.ReadJSONToken()
		if len(args) > 0 {
			fmt.Println(args[0], ":")
			for _, script := range commands[args[0]] {
				if len(commands[args[0]]) > 0 {

					fmt.Println("  ", script)
				}
			}
		} else {
			for command, scripts := range commands {
				fmt.Println(command, ":")
				for _, script := range scripts {
					fmt.Println("  ", script)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
