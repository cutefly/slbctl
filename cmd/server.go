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

    ip, _ := cmd.Flags().GetString("ip")
    port, _ := cmd.Flags().GetString("port")

    if len(args) > 0 {
      fmt.Println("You're arguments were: " + strings.Join(args, ","))
    } else {
      // fmt.Println("Please provide the required arguments")
      fmt.Println("Value of ip flag: " + ip)
      fmt.Println("Value of port flag: " + port)
      apv.ConfigureServer(ip, port)
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
	serverCmd.Flags().StringP("ip", "i", "", "ip for server")
	serverCmd.Flags().StringP("port", "p", "", "port for server")

	// Making Flags Required
	serverCmd.MarkPersistentFlagRequired("ip")
	serverCmd.MarkPersistentFlagRequired("port")
}
