# GoZeroAPI 项目

[English](README.md) | 简体中文

## 项目介绍

这是一个使用 go-zero 框架构建的 RESTful API 服务，为应用系统提供后端服务支持。项目包含用户管理、产品管理和订单处理等功能模块。

## 功能特点

- 用户服务
  - 登录
  - 注册

- 产品服务
  - 创建产品
  - 更新产品
  - 删除产品
  - 获取产品详情
  - 产品列表

- 订单服务
  - 创建订单
  - 取消订单
  - 更新订单状态
  - 获取订单详情
  - 订单列表

## 环境要求

- Go 1.16 或更高版本
- go-zero 框架
- MySQL（用于数据存储）

## 安装

1. 克隆仓库
```bash
git clone [repository-url]
cd gozeroapi
```

2. 安装依赖
```bash
go mod tidy
```

## 配置

配置文件位于 `etc` 目录下，需要根据您的环境配置以下文件：

- `etc/user-api.yaml`: 用户服务配置
- `etc/product-api.yaml`: 产品服务配置
- `etc/order-api.yaml`: 订单服务配置

## 运行服务

项目使用 Cobra CLI 框架作为命令行接口。您可以使用以下命令启动不同的服务：

1. 构建项目
```bash
go build -o goapi cmd/main.go
```

2. 查看可用命令
```bash
./goapi --help
```

3. 启动服务
```bash
# 启动用户服务
./goapi user -f etc/user-api.yaml

# 启动产品服务
./goapi product -f etc/product-api.yaml

# 启动订单服务
./goapi order -f etc/order-api.yaml
```

4. 查看版本
```bash
./goapi --version
```

## API 文档

API 定义文件位于 `api` 目录下：

- `api/user.api`: 用户服务 API 定义
- `api/product.api`: 产品服务 API 定义
- `api/order.api`: 订单服务 API 定义

## 项目结构

```
├── api/                # API 定义文件
├── cmd/                # 服务入口点
├── etc/                # 配置文件
├── internal/           # 内部包
│   ├── config/        # 配置结构
│   ├── handler/       # HTTP 处理器
│   ├── logic/         # 业务逻辑
│   ├── svc/          # 服务上下文
│   └── types/        # 类型定义
└── README.md          # 项目文档
```

## 贡献

欢迎贡献！请随时提交 Pull Request。