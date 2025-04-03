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

// configLoginCmd represents the login command
var configLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Configure username and password",
	Long: `This command allows you to configure a username and password for authentication within the system.  

For example:

To set the username to "username" and the password to "securepass", use the following command:

  cli-tool configure login username securepass

This ensures that the specified username and password are properly set, enabling secure access and authentication.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configLogin called")

		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		if len(args) > 0 {
			fmt.Println("You're arguments were: " + strings.Join(args, ","))
		} else {
			// fmt.Println("Please provide the required arguments")
			// fmt.Println("Value of username flag: " + username)
			// fmt.Println("Value of password flag: " + password)
			apv.ConfigureLogin(username, password)
		}
	},
}

func init() {
	configCmd.AddCommand(configLoginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configLoginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configLoginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	configLoginCmd.Flags().StringP("username", "u", "", "username for login")
	configLoginCmd.Flags().StringP("password", "p", "", "password for login")

	// Making Flags Required
	configLoginCmd.MarkPersistentFlagRequired("username")
	configLoginCmd.MarkPersistentFlagRequired("password")
}
