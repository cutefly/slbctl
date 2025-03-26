/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
  "strings"

	"github.com/spf13/cobra"

  "slbctl/apv"
)

// addMemberCmd represents the addMember command
var addMemberCmd = &cobra.Command{
	Use:   "member",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
  Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addMember called")
    if len(args) > 2 {
      fmt.Println("You're arguments were: " + strings.Join(args, ","))
    } else if len(args) == 2 {
      fmt.Println("Value of group arg: " + args[0])
      fmt.Println("Value of member arg: " + args[1])
      apv.AddGroupMember(args[0], args[1])
    } else {
      fmt.Println("Please provide the required arguments")
    }
	},
}

func init() {
	addCmd.AddCommand(addMemberCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addMemberCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addMemberCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
