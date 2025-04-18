/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"slbctl/apv"

	"github.com/spf13/cobra"
)

// commandCmd represents the command command
var commandCmd = &cobra.Command{
	Use:   "command <command string>",
	Short: "Execute loadbalancer's command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("command called")
		if len(args) > 1 {
			fmt.Println("You're arguments were: " + strings.Join(args, ","))
		} else if len(args) == 1 {
			// fmt.Println("Value of command arg: " + args[0])
			apv.ExecuteCommand(args[0])
		} else {
			fmt.Println("Please provide the required arguments")
		}
	},
}

func init() {
	execCmd.AddCommand(commandCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commandCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commandCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
