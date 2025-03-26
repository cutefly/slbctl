/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
  "slbctl/cmd"
  "github.com/spf13/viper"
)

func main() {
  viper.AutomaticEnv()
  viper.SetConfigName(".config") // name of config file (without extension)
  viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
  viper.AddConfigPath("./")               // optionally look for config in the working directory

	cmd.Execute()
}
