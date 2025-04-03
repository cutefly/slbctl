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

// configServerCmd represents the server command
var configServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Configure the server and debug options",
	Long: `This command allows you to configure server settings and debug options within the system.  

For example:

To set the server address to "192.168.1.100" and enable debug mode, use the following command:

  cli-tool configure server --url https://192.168.1.100:9997 skip-verify true --debug true

This ensures that the specified server configuration is applied and debug mode is enabled for troubleshooting and monitoring.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configServer called")

		url, _ := cmd.Flags().GetString("url")
		skipVerify, _ := cmd.Flags().GetBool("skip-verify")
		debug, _ := cmd.Flags().GetBool("debug")

		if len(args) > 0 {
			fmt.Println("You're arguments were: " + strings.Join(args, ","))
		} else {
			// fmt.Println("Please provide the required arguments")
			// fmt.Println("Value of url flag: " + url)
			apv.ConfigureServer(url, skipVerify, debug)
		}
	},
}

func init() {
	configCmd.AddCommand(configServerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configServerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configServerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	configServerCmd.Flags().StringP("url", "u", "", "url for server")
	configServerCmd.Flags().BoolP("skip-verify", "s", false, "skip verify SSL cert")
	configServerCmd.Flags().BoolP("debug", "d", false, "debug flag")

	// Making Flags Required
	configServerCmd.MarkPersistentFlagRequired("url")
}
