package logic

import (
	"context"
	"errors"
	"strings"

	"goapi/internal/svc"
	"goapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// 参数验证
	if len(req.Username) == 0 {
		return nil, errors.New("用户名不能为空")
	}
	if len(req.Password) < 6 {
		return nil, errors.New("密码长度不能小于6位")
	}
	if !strings.Contains(req.Email, "@") {
		return nil, errors.New("邮箱格式不正确")
	}

	// TODO: 检查用户名是否已存在
	// TODO: 密码加密
	// TODO: 保存用户信息到数据库

	return &types.RegisterResponse{
		Message: "注册成功",
	}, nil
}
