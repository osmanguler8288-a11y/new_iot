package test

import (
	"encoding/json"
	"fmt"
	"new_iot/define"
	"new_iot/helper"
	"testing"
)

var userServiceAddr = "http://127.0.0.1:9001"

func TestUserLogin(t *testing.T) {
	m := define.M{
		"username": "get",
		"password": "mqh12345",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/user/login", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
