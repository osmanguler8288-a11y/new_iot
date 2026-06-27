// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"new_iot/admin/internal/svc"
	"new_iot/admin/internal/types"
	"new_iot/helper"
	"new_iot/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductListLogic) ProductList(req *types.ProductListRequest) (resp *types.ProductListReply, err error) {
	list := make([]*types.ProductListBasic, 0)
	req.Page = helper.If(req.Page == 0, 0, (req.Page-1)*req.Size).(int)
	req.Size = helper.If(req.Size == 0, 0, req.Size).(int)
	resp = new(types.ProductListReply)
	var (
		count int64
	)

	err = models.ProductList(req.Name).Count(&count).Offset(req.Page).Limit(req.Size).Find(&list).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return
	}
	for _, v := range list {
		v.CreatedAt = helper.RFC3339ToNormalTime(v.CreatedAt)
	}

	resp.Count = count
	resp.List = list
	return
}
