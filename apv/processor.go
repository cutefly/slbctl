package apv

import (
  "fmt"
  "github.com/spf13/viper"
)

func Configure(username string, password string) error {
  fmt.Println("Configuring APV with username: " + username + " and password: " + password)
  viper.Set("username", username)
  viper.Set("password", password)
  viper.WriteConfigAs("./.config")
  fmt.Println("Configuring VIPER with username: " + viper.GetString("username") + " and password: " + viper.GetString("password"))
  return nil
}

func AddGroupMember(groupname string, membername string) error {
  fmt.Println("Adding member: " + membername + " to group: " + groupname)
  err := viper.ReadInConfig()
  if err != nil { // Handle errors reading the config file
    panic(fmt.Errorf("fatal error config file: %w", err))
  }
  fmt.Println("Configuring VIPER with username: " + viper.GetString("username") + " and password: " + viper.GetString("password"))
  // show slb group member를 통해 그룹에 소속되어 있는지 확인
  // 소속이 되어 있는 경우 skip, no error
  // 소속이 되어 있지 않은 경우 그룹에 추가
  return nil
}


func RemoveGroupMember(groupname string, membername string, force bool) error {
  fmt.Println("Removing member: " + membername + " from group: " + groupname + " with force:", force)
  err := viper.ReadInConfig()
  if err != nil { // Handle errors reading the config file
    panic(fmt.Errorf("fatal error config file: %w", err))
  }
  fmt.Println("Configuring VIPER with username: " + viper.GetString("username") + " and password: " + viper.GetString("password"))
  // show slb group member를 통해 그룹에 소속되어 있는지 확인
  // 소속이 되어 있지 않은 경우 skip, no error
  // 소속이 되어 있는 경우 force=false 인 경우 다른 멤버가 있는지 확인하여 다른 멤버가 있는 경우 요청한 멤버를 그룹에서 제거
  // 소속이 되어 있는 경우 force=false 인 경우 다른 멤버가 있는지 확인하여 다른 멤버가 없는 경우 skip, err 발생
  // 소속이 되어 있는 경우 force=true 인 경우 요청한 멤버를 그룹에서 제거
  return nil
}

func ShowGroupMember(groupname string) error {
  fmt.Println("Showing members of group: " + groupname)

  err := viper.ReadInConfig()
  if err != nil { // Handle errors reading the config file
    panic(fmt.Errorf("fatal error config file: %w", err))
  }
  fmt.Println("Configuring VIPER with username: " + viper.GetString("username") + " and password: " + viper.GetString("password"))
  return nil
}

