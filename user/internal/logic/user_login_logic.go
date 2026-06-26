// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"

	"new_iot/helper"
	"new_iot/models"
	"new_iot/user/internal/svc"
	"new_iot/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginRequest) (resp *types.UserLoginReply, err error) {
	ub := new(models.UserBasic)
	resp = new(types.UserLoginReply)
	err = l.svcCtx.DB.Where("name= ? AND password= ?", req.UserName, helper.Md5(req.PassWord)).First(ub).Error
	if err != nil {
		logx.Error("[error is]", err)
		err = errors.New("用户名或者密码错误")
		return
	}
	token, err := helper.GenerateToken(ub.ID, ub.Identity, ub.Name, 1*1000)
	if err != nil {
		logx.Error("[error is]", err)
		err = errors.New("用户名或者密码错误")
		return
	}
	resp.Token = token
	return
}
