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

var force bool

// removeMemberCmd represents the removeMember command
var removeMemberCmd = &cobra.Command{
	Use:   "member",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
  Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("removeMember called")
    force, _ := cmd.Flags().GetBool("force")
    if len(args) > 2 {
      fmt.Println("You're arguments were: " + strings.Join(args, ","))
    } else if len(args) == 2 {
      fmt.Println("Value of group arg: " + args[0])
      fmt.Println("Value of member arg: " + args[1])
      fmt.Println("Value of force flag:", force)
      apv.RemoveGroupMember(args[0], args[1], force)
    } else {
      fmt.Println("Please provide the required arguments")
    }
	},
}

func init() {
	removeCmd.AddCommand(removeMemberCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeMemberCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeMemberCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// removeMemberCmd.Flags().BoolP("force", "f", false, "force execute command")
	removeMemberCmd.Flags().BoolVarP(&force, "force", "f", false, "force execute command")
}
