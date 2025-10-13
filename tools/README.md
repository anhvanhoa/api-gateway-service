# API Gateway Generator Tool

Tool này tự động hóa việc cập nhật và generate code cho API Gateway.

## Chức năng

Tool sẽ thực hiện các bước sau:

1. **Update thư viện**: Chạy `go get -u github.com/anhvanhoa/sf-proto`
2. **Đọc config**: Đọc `dev.config.yml` hoặc `prod.config.yml` để lấy danh sách services
3. **Quét services**: Lặp qua `Env.Services` và sử dụng field `folder` để tạo đường dẫn import
4. **Kiểm tra tồn tại**: Kiểm tra xem service có tồn tại trong `github.com/anhvanhoa/sf-proto/gen/folder/v1`
5. **Tìm service handlers**: Tìm các function `Register*ServiceHandler` trong các service
6. **Cập nhật main.go**: Tự động import và đăng ký các service handlers vào `handlerMap`

## Cách sử dụng

### Cách 1: Sử dụng Makefile
```bash
make generate
```

### Cách 2: Chạy trực tiếp
```bash
go run tools/generate.go
```

### Cách 3: Sử dụng script
```bash
./scripts/generate.sh
```

## Kết quả

Sau khi chạy tool, sẽ tạo file `src/handler/service_handlers.go` với:

- Import statements cho tất cả các service packages
- Function `GetServiceHandlers()` trả về map chứa tất cả service handlers
- Tự động mapping service folders với handlers

File `src/main.go` sẽ được đơn giản hóa và sử dụng `handler.GetServiceHandlers()`

## Ví dụ output

### File `src/handler/service_handlers.go`:
```go
package handler

import (
    "context"
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"
    proto_auth "github.com/anhvanhoa/sf-proto/gen/auth/v1"
    proto_user "github.com/anhvanhoa/sf-proto/gen/user/v1"
    // ... other service imports
)

type ServiceHandler struct {
    Handler func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
    Swagger string
}

// GetServiceHandlers trả về map chứa tất cả service handlers
func GetServiceHandlers() map[string]ServiceHandler {
    return map[string]ServiceHandler{
        "auth": ServiceHandler{
            Handler: proto_auth.RegisterAuthServiceHandler,
            Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\auth\\v1\\auth.swagger.json",
        },
        "user": ServiceHandler{
            Handler: proto_user.RegisterUserServiceHandler,
            Swagger: "C:\\Users\\a\\go\\pkg\\mod\\github.com\\anhvanhoa\\sf-proto@v0.0.0-20251013045441-e59a7a5395f7\\gen\\user\\v1\\user.swagger.json",
        },
        // ... other handlers
    }
}
```

### File `src/main.go` (đơn giản hóa):
```go
func main() {
    // ... setup code ...
    handlerMap := handler.GetServiceHandlers()
    for _, service := range env.Services {
        serviceHandler, ok := handlerMap[service.Folder]
        if !ok {
            continue
        }
        baseHandler.AddService(handler.NewService(service.Name, service.Route, service.Host, service.Port, serviceHandler.Swagger, serviceHandler.Handler))
    }
    // ... rest of code ...
}
```

## Cấu hình

Tool sử dụng cấu hình từ file `dev.config.yml` hoặc `prod.config.yml`:

```yaml
services:
  - name: AUTH SERVICE
    host: localhost
    port: 50061
    route: /auth
    folder: auth  # Được sử dụng để tạo import path
  - name: USER SERVICE
    host: localhost
    port: 50062
    route: /users
    folder: user  # Được sử dụng để tạo import path
```

## Lưu ý

- Tool sẽ tự động backup file `main.go` hiện tại
- Đảm bảo có kết nối internet để update thư viện
- Tool sẽ tìm tất cả các service handlers có pattern `Register*ServiceHandler`
