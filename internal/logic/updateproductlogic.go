package logic

import (
	"context"
	"errors"
	"time"

	"goapi/internal/svc"
	"goapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProductLogic) UpdateProduct(req *types.UpdateProductRequest) (resp *types.UpdateProductResponse, err error) {
	// 参数验证
	if req.Id <= 0 {
		return nil, errors.New("无效的商品ID")
	}
	if req.Price != 0 && req.Price < 0 {
		return nil, errors.New("商品价格不能为负数")
	}
	if req.Stock != 0 && req.Stock < 0 {
		return nil, errors.New("商品库存不能为负数")
	}

	// TODO: 检查商品是否存在
	// TODO: 更新商品信息到数据库

	// 模拟更新商品
	product := types.Product{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	return &types.UpdateProductResponse{
		Product: product,
	}, nil
}
