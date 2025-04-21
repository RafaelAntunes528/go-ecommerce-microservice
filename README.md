# Go E-Commerce Microservice System

A complete e-commerce microservice implementation with API Gateway, Product Service, Order Service, and CLI tool - built with Go.

## Table of Contents
- [Project Structure](#project-structure)
- [Services Overview](#services-overview)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Running the System](#running-the-system)
- [CLI Usage](#cli-usage)
- [API Endpoints](#api-endpoints)
- [Adding Logging](#adding-logging)
- [Containerization with Docker](#containerization-with-docker)
- [Future Improvements](#future-improvements)

## Project Structure
go-ecommerce/
├── api-gateway/ # API Gateway service
├── product-service/ # Product catalog service
├── order-service/ # Order processing service
├── cli/ # Command line interface
└── README.md # This documentation


## Services Overview

1. **API Gateway** (`:8080`)
   - Routes requests to appropriate services
   - Single entry point for the system

2. **Product Service** (`:8081`)
   - Manages product catalog
   - Endpoints for listing and viewing products

3. **Order Service** (`:8082`)
   - Handles order creation and management
   - Processes new orders

4. **CLI Tool**
   - Command line interface for system interaction
   - Supports listing products and creating orders

## Getting Started

### Prerequisites

- Go 1.21+
- Git
- Docker (optional, for containerization)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/go-ecommerce.git
   cd go-ecommerce
   ```

2. Install dependencies:
    ```bash
    go mod download
    ```

## Running the System

### Local Development

Run each service in separate terminals:
    ```bash
    # Terminal 1 - Product Service
    cd product-service
    go run main.go

    # Terminal 2 - Order Service
    cd ../order-service
    go run main.go

    # Terminal 3 - API Gateway
    cd ../api-gateway
    go run main.go
    ```

### Building the CLI
    ```bash
    cd cli
    go build -o ecomcli.exe  # .exe for Windows
    ```

## CLI Usage
    ```bash
    # List all products
    ./ecomcli list-products

    # Create new order
    ./ecomcli create-order --product-id=1
    ```

## API Endpoints

Product Service

- **GET /products** - List all products
- **GET /products/{id}** - Get products details

Order Service

- **GET /orders** - List all orders
- **POST /orders** - Create new order

## Containerization with Docker

1. Create Dockerfile for each service

Example for product-service/Dockerfile:
    ```dockerfile
    # Build stage
    FROM golang:1.21 as builder
    WORKDIR /app
    COPY . .
    RUN go mod download
    RUN CGO_ENABLED=0 GOOS=linux go build -o product-service

    # Final stage
    FROM alpine:latest
    WORKDIR /app
    COPY --from=builder /app/product-service .
    EXPOSE 8081
    CMD ["./product-service"]
    ```

2. Create docker-compose.yml

    ```yaml
    version: '3.8'

    services:
    api-gateway:
        build: ./api-gateway
        ports:
        - "8080:8080"
        depends_on:
        - product-service
        - order-service

    product-service:
        build: ./product-service
        ports:
        - "8081:8081"

    order-service:
        build: ./order-service
        ports:
        - "8082:8082"
    ```

3. Build and run

    ```bash
    docker-compose up --build
    ```