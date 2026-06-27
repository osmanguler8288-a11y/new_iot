package config

import "github.com/zeromicro/go-zero/zrpc"

type MqttConf struct {
	Broker string
}

type Config struct {
	zrpc.RpcServerConf
	Mqtt struct {
		Broker   string
		ClientID string
		Password string
	}
}
