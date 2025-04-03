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

// showMemberCmd represents the showMember command
var showMemberCmd = &cobra.Command{
	Use:     "member <groupname>",
	Aliases: []string{"mem"},
	Short:   "shows the members of a group",
	Long: `This command allows you to view the services that are currently assigned to a designated loadbalancer group.  

For example:

To list all services in the "https-service-gr" group, use the following command:

  slbctl show members https-service-gr

This displays a list of all services that belong to the specified loadbalancer group, helping with access management and monitoring.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("showMember called")
		if len(args) > 1 {
			fmt.Println("You're arguments were: " + strings.Join(args, ","))
		} else if len(args) == 1 {
			// fmt.Println("Value of group arg: " + args[0])
			apv.ShowGroupMember(args[0])
		} else {
			fmt.Println("Please provide the required arguments")
		}
	},
}

func init() {
	showCmd.AddCommand(showMemberCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showMemberCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showMemberCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
