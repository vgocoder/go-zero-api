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
- Redis（用于缓存）

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

- `etc/user-api.yaml`: 用户服务配置（端口：8888）
- `etc/product-api.yaml`: 产品服务配置（端口：8889）
- `etc/order-api.yaml`: 订单服务配置（端口：8890）

配置文件结构：
```yaml
Name: service-name
Host: 0.0.0.0
Port: port-number

customLog:
  mode: file
  level: info
  path: logs/service-name.log

database:
  host: localhost
  port: 3306
  username: root
  password: password
  dbname: database_name

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0
```

## 运行服务

项目使用 Cobra CLI 框架作为命令行接口。您可以使用以下命令启动不同的服务：

1. 构建项目
```bash
go build -o goapi main.go
```

2. 直接运行（无需构建）
```bash
# 启动用户服务
go run main.go service user -f etc/user-api.yaml

# 启动产品服务
go run main.go service product -f etc/product-api.yaml

# 启动订单服务
go run main.go service order -f etc/order-api.yaml
```

3. 查看可用命令
```bash
./goapi --help
```

4. 启动服务
```bash
# 启动用户服务
./goapi service user -f etc/user-api.yaml

# 启动产品服务
./goapi service product -f etc/product-api.yaml

# 启动订单服务
./goapi service order -f etc/order-api.yaml
```

5. 查看版本
```bash
./goapi --version
```

## API 文档

服务启动后，API 文档可在 `/swagger/doc.html` 查看。

## 许可证

本项目采用 MIT 许可证 - 查看 LICENSE 文件了解详情。