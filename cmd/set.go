/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/anandureghu/linex/helper"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set <command name> script_1 script_2 .... script_n",
	Short: "add or update command and scripts",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		commands := helper.ReadJSONToken()
		scripts := []string{}
		for i, script := range args {
			if i != 0 {
				scripts = append(scripts, script)
			}
		}
		commands[args[0]] = scripts
		newCommands, err := json.MarshalIndent(commands, "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		helper.WriteJSON(newCommands)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
