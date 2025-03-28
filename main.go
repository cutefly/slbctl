/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"slbctl/apv"
	"slbctl/cmd"

	"github.com/spf13/viper"
)

var VERSION string = "0.0.1"

func main() {
	// viper.AutomaticEnv()
	// viper.SetConfigName(".config") // name of config file (without extension)
	// viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	// viper.AddConfigPath("./")      // optionally look for config in the working directory

	var config apv.Config

	// *viper.Viper 초기화
	// viperConfig := viper.New()
	// 설정 파일의 디렉토리 세팅
	viper.AddConfigPath(".")
	// 설정 파일명 세팅
	viper.SetConfigFile("config.yaml")

	// 설정 파일 읽어오기
	err := viper.ReadInConfig()
	if err != nil {
		// fmt.Println("Error on Reading Viper Config")
		config = apv.Config{}
	}

	//	config.Username = "user"
	//	config.Password = "pass"
	//	viperConfig.Set("username", config.Username)
	//	viperConfig.Set("password", config.Password)
	//	viperConfig.WriteConfig()
	// 읽어온 설정값을 config 로 언마샬
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("Error on Unmarshal Viper Config")
		panic(err)
	}
	// fmt.Println("config:", config)

	cmd.Execute(VERSION)
}
