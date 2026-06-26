package logic

import (
	"context"
	"errors"

	"new_iot/helper"
	"new_iot/user/rpc/internal/svc"
	"new_iot/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthLogic) Auth(in *user.UserAuthRequest) (*user.UserAuthReply, error) {
	// todo: add your logic here and delete this line
	if in.Token == "" {
		return nil, errors.New("参数不能为空")
	}
	Userclaim, err := helper.AnalyzeToken(in.Token)
	if err != nil {
		return nil, err
	}
	resp := new(user.UserAuthReply)
	resp.Id = uint64(Userclaim.Id)
	resp.Identity = Userclaim.Identity
	resp.Extend = map[string]string{
		"name": Userclaim.Name,
	}
	return resp, nil
}
