package apv

import (
	"encoding/json"
	"fmt"
)

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
