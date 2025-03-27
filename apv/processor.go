package apv

import (
  "fmt"
  "encoding/base64"
  "net/http"
  "github.com/spf13/viper"
)

type config struct {
  username string
  password string
  ip string
  port string
}

type member struct {
  real_service string
  weight int
  priority int
  active_status bool
  active_reason string
}

type group struct {
  instance_id string
  members []member
}

type group_detail struct {
  instance_id string
  group_name string
  method string
  activation int
  failvoer int
  priority_mode bool
  enable bool
  protocol string
  proxy_protocol bool
  members []member
  health_relation string
  hc_tcp_tempalte []string
  hc_http_tempalte []string
  group_policy_scope_name []string
}

func ConfigureLogin(username string, password string) error {
  viper.ReadInConfig()
  fmt.Println("Configuring APV with username: " + username + " and password: " + password)
  viper.Set("username", username)
  viper.Set("password", password)
  viper.WriteConfigAs("./.config")
  //viper.WriteConfig()
  fmt.Println("Configuring VIPER with username: " + viper.GetString("username") + " and password: " + viper.GetString("password"))
  return nil
}

func ConfigureServer(ip string, port string) error {
  viper.ReadInConfig()
  fmt.Println("Configuring APV with ip: " + ip + " and port: " + port)
  viper.Set("ip", ip)
  viper.Set("port", port)
  viper.WriteConfigAs("./.config")
  //viper.WriteConfig()
  fmt.Println("Configuring VIPER with ip: " + viper.GetString("ip") + " and port: " + viper.GetString("port"))
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
  ip := viper.GetString("ip")
  port := viper.GetString("port")
  username := viper.GetString("username")
  password := viper.GetString("password")

  url := fmt.Sprintf("https://%s:%s/rest/apv/loadbalancing/slb/group/Group/%s/members", ip, port, groupname)
  fmt.Println("Request URL:", url)
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    panic(fmt.Errorf("fatal error create http request: %w", err))
  }

  req.Header.Add("Authorization", "Basic " + BasicAuth(username, password))
  client := &http.Client{}

  res, err := client.Do(req)

  if err != nil {
    panic(fmt.Errorf("fatal error http request: %w", err))
  }

  defer res.Body.Close()

  return nil
}

func BasicAuth(username string, password string) string {
  auth := username + ":" + password
  return base64.StdEncoding.EncodeToString([]byte(auth))
}
