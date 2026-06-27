package test

import (
	"fmt"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func TestMqtt(t *testing.T) {
	opt := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883").SetClientID("go-test").
		SetUsername("get").SetPassword("123456")

	// 回调
	opt.SetDefaultPublishHandler(func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("MESSAGE : %s\n", message.Payload())
		fmt.Printf("TOPIC : %s\n", message.Topic())
	})

	c := mqtt.NewClient(opt)

	// 连接
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}
	time.Sleep(time.Second * 10)
	// 订阅主题
	if token := c.Subscribe("/sys/1/device_key/ping", 0, nil); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}

	// 发布
	if token := c.Publish("/sys/1/device_key/ping", 0, false, "你好，我是侯慧琳"); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}

	time.Sleep(time.Second * 10)

	// 取消订阅
	if token := c.Unsubscribe("/sys/1/device_key/ping"); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}
	fmt.Println("取消成功")
	// 关闭连接
	c.Disconnect(250)
}
