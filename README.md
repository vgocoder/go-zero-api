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

The project uses Cobra CLI framework for command-line interface. You can start different services using the following commands:

1. Build the project
```bash
go build -o goapi cmd/main.go
```

2. View available commands
```bash
./goapi --help
```

3. Start services
```bash
# Start user service
./goapi user -f etc/user-api.yaml

# Start product service
./goapi product -f etc/product-api.yaml

# Start order service
./goapi order -f etc/order-api.yaml
```

4. View version
```bash
./goapi --version
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

Contributions are welcome! Please feel free to submit a Pull Request.