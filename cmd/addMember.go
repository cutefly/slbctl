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
	Use:     "member <groupname> <membername>",
	Aliases: []string{"mem"},
	Short:   "add member to group",
	Long: `This command allows you to add a specified service to a designated loadbalancer group within the system.  

For example:

To add a service named "https-service-1" to the "https-service-gr" group, use the following command:

  slbctl add member https-service-1 https-service-gr

This ensures that the specified service is successfully assigned to the loadbalancer group, enabling appropriate access and management.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addMember called")
		if len(args) > 2 {
			fmt.Println("You're arguments were: " + strings.Join(args, ","))
		} else if len(args) == 2 {
			// fmt.Println("Value of group arg: " + args[0])
			// fmt.Println("Value of member arg: " + args[1])
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
