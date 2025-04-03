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
	Use:     "member <groupname> <membername>",
	Aliases: []string{"mem"},
	Short:   "remove a member from a group",
	Long: `This command allows you to remove a specified service from a designated loadbalancer group within the system.  

For example:

To remove a service named "https-service-1" from the "https-service-gr" group, use the following command:

  slbctl remove member https-service-1 https-service-gr

This ensures that the specified service is successfully removed from the loadbalancer group, revoking its access and association.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("removeMember called")
		force, _ := cmd.Flags().GetBool("force")
		if len(args) > 2 {
			fmt.Println("You're arguments were: " + strings.Join(args, ","))
		} else if len(args) == 2 {
			//   fmt.Println("Value of group arg: " + args[0])
			//   fmt.Println("Value of member arg: " + args[1])
			//   fmt.Println("Value of force flag:", force)
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
