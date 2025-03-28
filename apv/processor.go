package apv

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"crypto/tls"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

type Config struct {
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	URL        string `yaml:"url"`
	SkipVerify bool   `yaml:"skip-verify"`
}

type Member struct {
	RealService  string `json:"real_service"`
	Weight       int    `json:"weight"`
	Priority     int    `json:"priority"`
	ActiveStatus bool   `json:"active_status"`
	ActiveReason string `json:"active_reason"`
}

type Group struct {
	InstanceId           string   `json:"instance_id"`
	Group                string   `json:"group"`
	RealService          string   `json:"real_service"`
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

type SimpleGroup struct {
	InstanceId string   `json:"instance_id"`
	Members    []Member `json:"members"`
}

type GroupRequest struct {
	RealService string `json:"real_service"`
}

type GroupResponse struct {
	Group Group `json:"Group"`
}

type CliRequest struct {
	Cmd string `json:"cmd"`
}

type CliResponse struct {
	Contents string `json:"contents"`
}

var config Config

func ConfigureLogin(username string, password string) error {
	_ = viper.Unmarshal(&config)
	// fmt.Println("Configuring APV with username: " + username + " and password: " + password)
	viper.Set("username", username)
	viper.Set("password", password)
	config = Config{Username: username, Password: password}
	viper.WriteConfig()
	//fmt.Println("Configuring VIPER with username: " + viper.GetString("username") + " and password: " + viper.GetString("password"))
	fmt.Println("Configuring viper with username and password")
	return nil
}

func ConfigureServer(url string, skipVerify bool) error {
	_ = viper.Unmarshal(&config)
	// fmt.Println("Configuring APV with username: " + username + " and password: " + password)
	viper.Set("url", url)
	viper.Set("skip-verify", skipVerify)
	config = Config{URL: url}
	viper.WriteConfig()
	// fmt.Println("Configuring VIPER with url: " + viper.GetString("url"))
	fmt.Println("Configuring viper with URL")
	return nil
}

func AddGroupMember(groupname string, membername string) error {
	fmt.Println("Adding member: " + membername + " to group: " + groupname)

	_ = viper.Unmarshal(&config)
	// show slb group member를 통해 그룹에 소속되어 있는지 확인
	isMember, err := isGroupMember(groupname, membername)
	// 소속이 되어 있는 경우 skip, no error
	if err != nil {
		return fmt.Errorf("error checking group membership: %w", err)
	}
	if isMember {
		fmt.Println("Member: " + membername + " is already a member of group: " + groupname)
		return nil
	}

	// 소속이 되어 있지 않은 경우 그룹에 추가
	reqUrl := fmt.Sprintf("%s/rest/apv/loadbalancing/slb/group/Group/%s/members", config.URL, groupname)
	fmt.Println("Request URL:", reqUrl)
	groupRequest := GroupRequest{membername}
	//JSON 인코딩
	jsonBytes, err := json.Marshal(groupRequest)
	if err != nil {
		panic(err)
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: config.SkipVerify}
	req, err := http.NewRequest(http.MethodPost, reqUrl, bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(fmt.Errorf("fatal error create http request: %w", err))
	}

	req.Header.Add("Authorization", "Basic "+basicAuth(config.Username, config.Password))
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		panic(fmt.Errorf("fatal error http request: %w", err))
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body) // body를 읽으면 이렇게 해야 된다.

	if err != nil {
		panic(err)
	}
	// fmt.Println("body:", string(body))

	thisRes := GroupResponse{}
	parseErr := json.Unmarshal(body, &thisRes) // json parse

	if parseErr != nil {
		panic(parseErr)
	}

	if thisRes.Group.RealService != membername {
		panic(fmt.Errorf("error adding member: %s to group: %s", membername, groupname))
	}
	fmt.Println("Member: " + membername + " is a member of group: " + groupname)

	return nil
}

func RemoveGroupMember(groupname string, membername string, force bool) error {
	fmt.Println("Removing member: "+membername+" from group: "+groupname+" with force:", force)

	_ = viper.Unmarshal(&config)
	// show slb group member를 통해 그룹에 소속되어 있는지 확인
	isMember, err := isGroupMember(groupname, membername)
	// 소속이 되어 있지 않은 경우 skip, no error
	if err != nil {
		return fmt.Errorf("error checking group member: %w", err)
	}
	if !isMember {
		fmt.Println("Member: " + membername + " is already not a member of group: " + groupname)
		return nil
	}

	// 소속이 되어 있는 경우 force=false 인 경우 다른 멤버가 있는지 확인하여 다른 멤버가 있는 경우 요청한 멤버를 그룹에서 제거
	if !force {
		// 소속이 되어 있는 경우 force=false 인 경우 다른 멤버가 있는지 확인하여 다른 멤버가 없는 경우 skip, err 발생
		members, err := getMembers(groupname)
		if err != nil {
			panic(fmt.Errorf("fatal error get members: %w", err))
		}
		if len(members) <= 1 && members[0].RealService == membername {
			fmt.Println("Member: " + membername + " is the only member of group: " + groupname)
			return fmt.Errorf("error removing member: %s from group: %s, no other members", membername, groupname)
		}
	}

	// 소속이 되어 있는 경우 force=true 인 경우 요청한 멤버를 그룹에서 제거
	// Delete API가 정상동작하지 않아 cli_extend로 대체
	// reqUrl := fmt.Sprintf("%s/rest/apv/loadbalancing/slb/group/Group/%s/members", config.URL, groupname)
	reqUrl := fmt.Sprintf("%s/rest/apv/cli_extend", config.URL)
	fmt.Println("Request URL:", reqUrl)
	deleteCommand := fmt.Sprintf("no slb group member %s %s", groupname, membername)
	thisReq := CliRequest{deleteCommand}
	//JSON 인코딩
	jsonBytes, err := json.Marshal(thisReq)
	if err != nil {
		panic(err)
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: config.SkipVerify}
	req, err := http.NewRequest(http.MethodPost, reqUrl, bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(fmt.Errorf("fatal error create http request: %w", err))
	}

	req.Header.Add("Authorization", "Basic "+basicAuth(config.Username, config.Password))
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		panic(fmt.Errorf("fatal error http request: %w", err))
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body) // body를 읽으면 이렇게 해야 된다.

	if err != nil {
		panic(err)
	}
	//fmt.Println("body:", string(body))

	thisRes := CliResponse{}
	parseErr := json.Unmarshal(body, &thisRes) // json parse

	if parseErr != nil {
		panic(parseErr)
	}

	if thisRes.Contents != "" {
		panic(fmt.Errorf("error removing member: %s from group: %s", membername, groupname))
	}

	isMember, err = isGroupMember(groupname, membername)
	if isMember || err != nil {
		panic(fmt.Errorf("error checking group member: %w", err))
	}
	fmt.Println("Member: " + membername + " is removed from group: " + groupname)

	return nil
}

func ShowGroupMember(groupname string) error {
	fmt.Println("Showing members of group: " + groupname)

	_ = viper.Unmarshal(&config)

	members, err := getMembers(groupname)
	if err != nil {
		panic(fmt.Errorf("fatal error create http request: %w", err))
	}

	for _, s := range members {
		fmt.Println(groupname + "\t" + s.RealService)
	}

	return nil
}

func isGroupMember(groupname string, membername string) (bool, error) {
	members, err := getMembers(groupname)
	if err != nil {
		panic(fmt.Errorf("fatal error get members: %w", err))
	}

	for _, s := range members {
		if s.RealService == membername {
			//fmt.Println("member:", s.RealService)
			return true, nil
		}
	}

	return false, nil
}

func getMembers(groupname string) ([]Member, error) {
	reqUrl := fmt.Sprintf("%s/rest/apv/loadbalancing/slb/group/Group/%s/members", config.URL, groupname)
	// fmt.Println("Request URL:", reqUrl)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: config.SkipVerify}
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		panic(fmt.Errorf("fatal error create http request: %w", err))
	}

	req.Header.Add("Authorization", "Basic "+basicAuth(config.Username, config.Password))
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		panic(fmt.Errorf("fatal error http request: %w", err))
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body) // body를 읽으면 이렇게 해야 된다.

	if err != nil {
		panic(err)
	}
	//fmt.Println("body:", string(body))

	thisRes := GroupResponse{}
	parseErr := json.Unmarshal(body, &thisRes) // json parse

	if parseErr != nil {
		panic(parseErr)
	}

	// fmt.Println("members:", thisRes.Group.Members)
	return thisRes.Group.Members, nil
}

func basicAuth(username string, password string) string {
	auth := username + ":" + password

	//fmt.Println("auth:", auth)
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func Test() {
	fmt.Println("do test")
	member1 := Member{"kubernetes-dev-32443-1", 100, 1, true, "active"}
	member2 := Member{"kubernetes-dev-32443-2", 100, 1, true, "active"}
	member3 := Member{"kubernetes-dev-32443-3", 100, 1, true, "active"}
	group := Group{GroupName: "kubernetes-dev-32443-gr", Members: []Member{member1, member2, member3}}
	resultJSON := GroupResponse{group}
	//JSON 인코딩
	jsonBytes, err := json.Marshal(resultJSON)
	if err != nil {
		panic(err)
	}
	//JSON 바이트를 문자열로 변경
	jsonString := string(jsonBytes)
	fmt.Println("group:", jsonString)

	// body := []byte(`{"group": {"instance_id": "kubernetes-dev-32443-gr"}}`)
	// fmt.Println("body:\n" + string(body))

	thisRes := GroupResponse{}
	parseErr := json.Unmarshal(jsonBytes, &thisRes) // json parse

	if parseErr != nil {
		panic(parseErr)
	}

	fmt.Println("members:", len(thisRes.Group.Members))
}

type MemberType struct {
	Name   string
	Age    int
	Active bool
}

type GroupType struct {
	Group MemberType
}

func TestJson() {
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
}
