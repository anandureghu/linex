/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/anandureghu/linex/helper"
	"github.com/spf13/cobra"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply <command name>",
	Short: "execute all scripts of mentioned command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		mydir, _ := os.Getwd()
		fmt.Println("apply called at", mydir)

		commands := helper.ReadJSONToken()

		if commands[args[0]] != nil && len(commands[args[0]]) > 0 {
			for _, script := range commands[args[0]] {
				if len(commands[args[0]]) > 0 {
					fmt.Printf("applying '%s'\n", script)

					scriptArgs := helper.SplitString(script)
					cmd := exec.Command(scriptArgs[0], scriptArgs[1:]...)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr

					err := cmd.Run()
					if err != nil {
						log.Fatal(err)
					}

					log.Println("Command executed successfully.")
				}
			}
		} else {
			fmt.Printf("no scripts set for '%s', set scripts for this command to conitinue\n", args[0])
		}

	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
