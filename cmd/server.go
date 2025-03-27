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

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")

		url, _ := cmd.Flags().GetString("url")

		if len(args) > 0 {
			fmt.Println("You're arguments were: " + strings.Join(args, ","))
		} else {
			// fmt.Println("Please provide the required arguments")
			fmt.Println("Value of url flag: " + url)
			apv.ConfigureServer(url)
		}
	},
}

func init() {
	configCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serverCmd.Flags().StringP("url", "u", "", "url for server")

	// Making Flags Required
	serverCmd.MarkPersistentFlagRequired("url")
}
