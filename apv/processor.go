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

var config Config

func init() {
	// viper.SetConfigName(".config") // name of config file (without extension)
	// viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	// viper.AddConfigPath("./")      // optionally look for config in the working directory
	viper.SetConfigFile("config.yaml")
	viper.SetEnvPrefix("SLBCTL")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	// 설정 파일 읽어오기
	err := viper.ReadInConfig()
	if err != nil {
		// fmt.Println("Error on Reading Viper Config")
		config = Config{}
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("Error on Unmarshal Viper Config")
		panic(err)
	}

	fmt.Println("Viper Config: ", config)
	fmt.Println("Configuring APV with username: " + config.Username + " and password: " + config.Password)
}

/**
 * Configure username and password
 */
func ConfigureLogin(username string, password string) error {

	// fmt.Println("Configuring APV with username: " + username + " and password: " + password)
	viper.Set("username", username)
	viper.Set("password", password)
	//config = Config{Username: username, Password: password}
	viper.WriteConfig()
	//fmt.Println("Configuring VIPER with username: " + viper.GetString("username") + " and password: " + viper.GetString("password"))
	fmt.Println("Configuring viper with username and password")
	return nil
}

/**
 * ConfigureServer configures the server URL, skip-verify and debug flag
 */
func ConfigureServer(url string, skipVerify bool, debug bool) error {
	viper.Set("url", url)
	viper.Set("skip-verify", skipVerify)
	viper.Set("debug", debug)
	//config = Config{URL: url, SkipVerify: skipVerify}
	viper.WriteConfig()
	// fmt.Println("Configuring VIPER with url: " + viper.GetString("url"))
	fmt.Println("Configuring viper with url, skip-verify and debug flag")
	return nil
}

/**
 * AddGroupMember adds a member to a group.
 */
func AddGroupMember(groupname string, membername string) error {
	if config.Debug {
		fmt.Println("Adding member: " + membername + " to group: " + groupname)
	}

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
	if config.Debug {
		fmt.Println("Request URL:", http.MethodPost, reqUrl)
	}
	groupRequest := GroupRequest{membername}
	jsonBytes, err := json.Marshal(groupRequest)
	if err != nil {
		panic(err)
	}
	if config.Debug {
		fmt.Println("req body:", string(jsonBytes))
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

	if config.Debug {
		fmt.Println("res body:", string(body))
	}

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

/**
 * RemoveGroupMember removes a member from a group.
 */
func RemoveGroupMember(groupname string, membername string, force bool) error {
	if config.Debug {
		fmt.Println("Removing member: "+membername+" from group: "+groupname+" with force:", force)
	}

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
	/*
		reqUrl := fmt.Sprintf("%s/rest/apv/cli_extend", config.URL)
		fmt.Println("Request URL:", reqUrl)
		deleteCommand := fmt.Sprintf("no slb group member %s %s", groupname, membername)
		thisReq := CliRequest{deleteCommand}
		//JSON 인코딩
		jsonBytes, err := json.Marshal(thisReq)
		if err != nil {
			panic(err)
		}
	*/

	reqUrl := fmt.Sprintf("%s/rest/apv/batch_cli", config.URL)
	if config.Debug {
		fmt.Println("Request URL:", http.MethodPost, reqUrl)
	}
	deleteCommand := fmt.Sprintf("no slb group member %s %s", groupname, membername)
	if config.Debug {
		fmt.Println("req body:", deleteCommand)
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: config.SkipVerify}
	req, err := http.NewRequest(http.MethodPost, reqUrl, bytes.NewBufferString(deleteCommand))
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

	if config.Debug {
		fmt.Println("res body:", string(body))
	}

	thisRes := BatchCliResponse{}
	parseErr := json.Unmarshal(body, &thisRes) // json parse

	if parseErr != nil {
		panic(parseErr)
	}

	if thisRes.Output != "" {
		panic(fmt.Errorf("error removing member: %s from group: %s", membername, groupname))
	}

	isMember, err = isGroupMember(groupname, membername)
	if isMember || err != nil {
		panic(fmt.Errorf("error checking group member: %w", err))
	}
	fmt.Println("Member: " + membername + " is removed from group: " + groupname)

	return nil
}

/**
 * ShowGroupMember shows members in a group.
 */
func ShowGroupMember(groupname string) error {
	if config.Debug {
		fmt.Println("Showing members of group: " + groupname)
	}
	// fmt.Println("config: ", config)

	members, err := getMembers(groupname)
	if err != nil {
		panic(fmt.Errorf("fatal error create http request: %w", err))
	}

	for _, s := range members {
		fmt.Println(groupname + "\t" + s.RealService)
	}

	return nil
}

/**
 * ExecuteCommand executes a command on the APV.
 */
func ExecuteCommand(cmd string) error {
	_ = viper.Unmarshal(&config)
	if config.Debug {
		fmt.Println("Executing command: " + cmd)
	}

	// fmt.Println("config: ", config)

	reqUrl := fmt.Sprintf("%s/rest/apv/cli_extend", config.URL)
	if config.Debug {
		fmt.Println("Request URL:", http.MethodPost, reqUrl)
	}
	thisReq := CliRequest{cmd}
	if config.Debug {
		fmt.Println("req body:", thisReq)
	}
	//JSON 인코딩
	jsonBytes, err := json.Marshal(thisReq)
	if err != nil {
		panic(err)
	}
	if config.Debug {
		fmt.Println("req body:", string(jsonBytes))
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

	if config.Debug {
		fmt.Println("res body:", string(body))
	}

	thisRes := CliResponse{}
	parseErr := json.Unmarshal(body, &thisRes) // json parse

	if parseErr != nil {
		panic(parseErr)
	}

	fmt.Println(thisRes.Contents)

	return nil
}

/**
 * [private] isGroupMember checks if a member is in a group.
 */
func isGroupMember(groupname string, membername string) (bool, error) {
	members, err := getMembers(groupname)
	if err != nil {
		panic(fmt.Errorf("fatal error get members: %w", err))
	}

	for _, s := range members {
		if s.RealService == membername {
			return true, nil
		}
	}

	return false, nil
}

/**
 * [private] GetMembers gets members of a group.
 */
func getMembers(groupname string) ([]Member, error) {
	reqUrl := fmt.Sprintf("%s/rest/apv/loadbalancing/slb/group/Group/%s/members", config.URL, groupname)
	if config.Debug {
		fmt.Println("Request URL:", http.MethodGet, reqUrl)
	}
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

	if config.Debug {
		fmt.Println("res body:", string(body))
	}
	thisRes := GroupResponse{}
	parseErr := json.Unmarshal(body, &thisRes) // json parse

	if parseErr != nil {
		panic(parseErr)
	}

	// fmt.Println("members:", thisRes.Group.Members)
	return thisRes.Group.Members, nil
}

/**
 * [private] basicAuth encodes the username and password in base64.
 */
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
	// fmt.Println("res body:\n" + string(body))

	thisRes := GroupResponse{}
	parseErr := json.Unmarshal(jsonBytes, &thisRes) // json parse

	if parseErr != nil {
		panic(parseErr)
	}

	fmt.Println("members:", len(thisRes.Group.Members))
}
