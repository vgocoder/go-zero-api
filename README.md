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

- `etc/user-api.yaml`: User service configuration
- `etc/product-api.yaml`: Product service configuration
- `etc/order-api.yaml`: Order service configuration

## Running the Services

1. Start the user service
```bash
cd cmd/user
go run main.go -f ../../etc/user-api.yaml
```

2. Start the product service
```bash
cd cmd/product
go run main.go -f ../../etc/product-api.yaml
```

3. Start the order service
```bash
cd cmd/order
go run main.go -f ../../etc/order-api.yaml
```

## API Documentation

The API definitions can be found in the `api` directory:

- `api/user.api`: User service API definitions
- `api/product.api`: Product service API definitions
- `api/order.api`: Order service API definitions

## Project Structure

```
├── api/                # API definitions
├── cmd/                # Service entry points
├── etc/                # Configuration files
├── internal/           # Internal packages
│   ├── config/        # Configuration structures
│   ├── handler/       # HTTP handlers
│   ├── logic/         # Business logic
│   ├── svc/          # Service context
│   └── types/        # Type definitions
└── README.md          # Project documentation
```

## Contributing

Welcome to contribute to this project. You can:

1. Submit issues
2. Create pull requests
3. Improve documentation

## License

This project is licensed under the MIT License.