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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
