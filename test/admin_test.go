package test

import (
	"encoding/json"
	"fmt"
	"new_iot/helper"
	"testing"
)

var adminServiceAddr = "http://127.0.0.1:9002"

func TestDeviceList(t *testing.T) {
	// 第一步：登录拿新 token
	loginBody, _ := json.Marshal(map[string]string{
		"username": "get",
		"password": "mqh12345",
	})
	loginResp, err := helper.HttpPost(userServiceAddr+"/user/login", loginBody)
	if err != nil {
		t.Fatal("登录失败:", err)
	}
	var result map[string]string
	json.Unmarshal(loginResp, &result)
	token := result["token"]
	fmt.Println("获取到 token:", token)

	// 第二步：拿 token 调 admin
	header, _ := json.Marshal(map[string]string{"token": token})
	rep, err := helper.HttpGet(adminServiceAddr+"/device/list?page=16&size=20&name=", header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

func TestProDuctList(t *testing.T) {
	// 不填 token，验证被拦截
	header, _ := json.Marshal(map[string]string{"token": ""})
	rep, err := helper.HttpGet(adminServiceAddr+"/device/list?page=16&size=20&name=", header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}
