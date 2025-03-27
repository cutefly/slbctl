package apv

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

type Config struct {
	username string
	password string
	url      string
}

type Member struct {
	RealService  string `json:"real_service"`
	Weight       int    `json:"weight"`
	Priority     int    `json:"priority"`
	ActiveStatus bool   `json:"active_status"`
	ActiveReason string `json:"active_reason"`
}

type Group struct {
	InstanceId string   `json:"instance_id"`
	Members    []Member `json:"members"`
}

type GroupMember struct {
	Group Group `json:"Group"`
}

type GroupDetail struct {
	InstanceId           string   `json:"instance_id"`
	GroupName            string   `json:"group_name"`
	Method               string   `json:"method"`
	Activation           int      `json:"activation"`
	Failvoer             int      `json:"failover"`
	PriorityMode         bool     `json:"priority_mode"`
	Enable               bool     `json:"enable"`
	Protocol             string   `json:"protocol"`
	ProxyProtocol        bool     `json:"proxy_protocol"`
	Members              []Member `json:"members"`
	HealthRelation       string   `json:"health_relation"`
	HcTcpTempalte        []string `json:"hc_tcp_tempalte"`
	HcHttpTempalte       []string `json:"hc_http_tempalte"`
	GroupPolicyScopeName []string `json:"group_policy_scope_name"`
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

func ConfigureServer(url string) error {
	viper.ReadInConfig()
	fmt.Println("Configuring APV with url: " + url)
	viper.Set("url", url)
	viper.WriteConfigAs("./.config")
	//viper.WriteConfig()
	fmt.Println("Configuring VIPER with url: " + viper.GetString("url"))
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
	fmt.Println("Removing member: "+membername+" from group: "+groupname+" with force:", force)
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
	url := viper.GetString("url")
	username := viper.GetString("username")
	password := viper.GetString("password")

	reqUrl := fmt.Sprintf("%s/rest/apv/loadbalancing/slb/group/Group/%s/members", url, groupname)
	fmt.Println("Request URL:", reqUrl)
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		panic(fmt.Errorf("fatal error create http request: %w", err))
	}

	req.Header.Add("Authorization", "Basic "+BasicAuth(username, password))
	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		panic(fmt.Errorf("fatal error http request: %w", err))
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body) // body를 읽으면 이렇게 해야 된다.

	if err != nil {
		panic(err)
	}
	fmt.Println("body:", string(body)) // body {"code":200,"message":"hi"}

	thisRes := GroupMember{}
	parseErr := json.Unmarshal(body, &thisRes) // json parse

	if parseErr != nil {
		panic(parseErr)
	}

	fmt.Println("group:", thisRes) // body {"code":200,"message":"hi"}

	return nil
}

func BasicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

type MemberType struct {
	Name   string
	Age    int
	Active bool
}

type GroupType struct {
	Group MemberType
}

func Test() {
	fmt.Println("do test")
	member1 := Member{"kubernetes-dev-32443-1", 100, 1, true, "active"}
	member2 := Member{"kubernetes-dev-32443-2", 100, 1, true, "active"}
	member3 := Member{"kubernetes-dev-32443-3", 100, 1, true, "active"}
	group := Group{"kubernetes-dev-32443-gr", []Member{member1, member2, member3}}
	resultJSON := GroupMember{group}
	//JSON 인코딩
	jsonBytes, err := json.Marshal(resultJSON)
	if err != nil {
		panic(err)
	}
	//JSON 바이트를 문자열로 변경
	jsonString := string(jsonBytes)
	fmt.Println("group:", jsonString)

	// body := []byte(`{"group": {"instance_id": "kubernetes-dev-32443-gr"}}`)
	// fmt.Println("body:\n" + string(body)) // body {"code":200,"message":"hi"}

	groupMember := GroupMember{}
	parseErr := json.Unmarshal(jsonBytes, &groupMember) // json parse

	if parseErr != nil {
		panic(parseErr)
	}

	fmt.Println("members:", len(groupMember.Group.Members)) // body {"code":200,"message":"hi"}
	/*
		mem := MemberType{"Dusdj", 23, true}
		group := GroupType{mem}

		//JSON 인코딩
		jsonBytes, err := json.Marshal(group)
		if err != nil {
			panic(err)
		}

		//JSON 바이트를 문자열로 변경
		jsonString := string(jsonBytes)
		fmt.Println(jsonString)

		//JSON 디코딩
		newGroup := GroupType{}
		err = json.Unmarshal(jsonBytes, &newGroup)
		if err != nil {
			panic(err)
		}

		fmt.Println(newGroup.Group.Name)
	*/
}
