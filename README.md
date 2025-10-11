# API Gateway

A Go-based API Gateway service that provides HTTP REST endpoints by proxying requests to gRPC microservices.

## Overview

This API Gateway acts as a reverse proxy, translating HTTP REST requests to gRPC calls and forwarding them to backend microservices. It's built using:

- **Gin** - HTTP web framework
- **gRPC-Gateway** - Protocol buffer compiler plugin for generating reverse-proxy code
- **gRPC** - High-performance RPC framework

## Features

- HTTP to gRPC protocol translation
- IoT Device service integration
- Custom error handling middleware
- Structured logging
- Graceful error recovery

## Project Structure

```
api-gateway/
├── src/
│   ├── handler/
│   │   ├── base.go              # Base handler structure
│   │   ├── error_middleware.go  # Custom error handling
│   │   └── iot_device.go       # IoT Device service handler
│   └── main.go                  # Application entry point
├── go.mod                       # Go module dependencies
├── go.sum                       # Go module checksums
├── Makefile                     # Build automation
└── README.md                    # This file
```

## Prerequisites

- Go 1.25.0 or later
- gRPC backend services running
- Protocol buffer definitions

## Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd api-gateway
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Build the application**
   ```bash
   go build -o api-gateway src/main.go
   ```

   Or use the Makefile:
   ```bash
   make build
   ```

## Configuration

The API Gateway connects to gRPC services on the following endpoints:

- **IoT Device Service**: `localhost:50060`

To modify these endpoints, update the connection strings in the respective handler files.

## Running the Application

1. **Start the API Gateway**
   ```bash
   ./api-gateway
   ```

   Or:
   ```bash
   go run src/main.go
   ```

2. **Verify the service is running**
   ```bash
   curl http://localhost:8080/health
   ```

The API Gateway will be available at `http://localhost:8080`

## API Endpoints

The gateway automatically exposes all gRPC service methods as HTTP REST endpoints based on the protocol buffer definitions.

### IoT Device Service

All IoT Device service methods are available under the root path with automatic HTTP method mapping.

## Development

### Adding New Services

To add support for new gRPC services:

1. Import the generated protobuf client
2. Create a new handler method in the appropriate handler file
3. Register the service in `main.go`
4. Update the router configuration

### Error Handling

Custom error handling is implemented in `src/handler/error_middleware.go` to provide consistent error responses across all services.

## Dependencies

Key dependencies include:

- `github.com/gin-gonic/gin` - HTTP web framework
- `github.com/grpc-ecosystem/grpc-gateway/v2` - gRPC to HTTP proxy
- `google.golang.org/grpc` - gRPC client library
- `github.com/anhvanhoa/sf-proto` - Protocol buffer definitions

## Building

```bash
# Build for current platform
make build

# Build for multiple platforms
make build-all

# Clean build artifacts
make clean
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

[Add your license information here]

## Support

For issues and questions, please create an issue in the repository or contact the development team.
