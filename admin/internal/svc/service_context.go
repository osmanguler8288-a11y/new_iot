package svc

import (
	"new_iot/admin/internal/config"
	"new_iot/models"
	"new_iot/user/rpc/user_client"

	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	RpcUser user_client.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB()
	return &ServiceContext{
		Config:  c,
		DB:      models.DB,
		RpcUser: user_client.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
