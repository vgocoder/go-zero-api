package logic

import (
	"context"
	"errors"

	"goapi/internal/svc"
	"goapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 参数验证
	if len(req.Username) == 0 {
		return nil, errors.New("用户名不能为空")
	}
	if len(req.Password) == 0 {
		return nil, errors.New("密码不能为空")
	}

	// TODO: 从数据库查询用户信息
	// TODO: 验证密码是否正确
	// TODO: 生成 JWT token

	// 临时返回一个模拟的 token
	return &types.LoginResponse{
		Token: "mock_jwt_token",
	}, nil
}
