# GoZeroAPI Project

English | [简体中文](README_zh.md)

## Introduction

This is a RESTful API service built with go-zero framework, providing backend services for a comprehensive application system. The project includes user management, product management, and order processing functionalities.

## Features

- User Service
  - Login
  - Registration

- Product Service
  - Create product
  - Update product
  - Delete product
  - Get product details
  - List products

- Order Service
  - Create order
  - Cancel order
  - Update order status
  - Get order details
  - List orders

## Requirements

- Go 1.16 or later
- go-zero framework
- MySQL (for data storage)
- Redis (for caching)

## Installation

1. Clone the repository
```bash
git clone [repository-url]
cd gozeroapi
```

2. Install dependencies
```bash
go mod tidy
```

## Configuration

The configuration files are located in the `etc` directory. You need to configure the following files according to your environment:

- `etc/user-api.yaml`: User service configuration (Port: 8888)
- `etc/product-api.yaml`: Product service configuration (Port: 8889)
- `etc/order-api.yaml`: Order service configuration (Port: 8890)

Configuration file structure:
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

## Running the Services

The project uses Cobra CLI framework for command-line interface. You can start different services using the following commands:

1. Build the project
```bash
go build -o goapi main.go
```

2. Run directly (without building)
```bash
# Start user service
go run main.go service user -f etc/user-api.yaml

# Start product service
go run main.go service product -f etc/product-api.yaml

# Start order service
go run main.go service order -f etc/order-api.yaml
```

3. View available commands
```bash
./goapi --help
```

4. Start services
```bash
# Start user service
./goapi service user -f etc/user-api.yaml

# Start product service
./goapi service product -f etc/product-api.yaml

# Start order service
./goapi service order -f etc/order-api.yaml
```

5. View version
```bash
./goapi --version
```

## API Documentation

API documentation is available at `/swagger/doc.html` after starting the service.

## License

This project is licensed under the MIT License - see the LICENSE file for details.