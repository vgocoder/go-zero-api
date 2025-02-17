package logic

import (
	"context"
	"errors"
	"time"

	"goapi/internal/svc"
	"goapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateProductLogic) CreateProduct(req *types.CreateProductRequest) (resp *types.CreateProductResponse, err error) {
	// 参数验证
	if len(req.Name) == 0 {
		return nil, errors.New("商品名称不能为空")
	}
	if req.Price <= 0 {
		return nil, errors.New("商品价格必须大于0")
	}
	if req.Stock < 0 {
		return nil, errors.New("商品库存不能为负数")
	}

	// TODO: 创建商品记录到数据库
	// 模拟创建商品
	product := types.Product{
		Id:          1, // 模拟ID
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	return &types.CreateProductResponse{
		Product: product,
	}, nil
}
