package handler

import (
	"context"
	"log"

	proto_iot_device "github.com/anhvanhoa/sf-proto/gen/iot_device/v1"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (h *BaseHandler) RegisterIoTDevice(
	ctx context.Context,
	router *gin.Engine,
) {
	// Tạo gRPC connection
	conn, err := grpc.NewClient(
		"localhost:50060",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Không thể kết nối đến gRPC server: %v", err)
	}

	// Tạo gRPC gateway mux
	grpcMux := runtime.NewServeMux(
		runtime.WithErrorHandler(CustomErrorHandler),
	)
	if err := proto_iot_device.RegisterIoTDeviceServiceHandler(ctx, grpcMux, conn); err != nil {
		log.Fatalf("Không thể đăng ký handler: %v", err)
	}

	// Đăng ký gRPC gateway vào Gin router
	router.Any("/iot-devices", gin.WrapH(grpcMux))
	router.Any("/iot-devices/*path", gin.WrapH(grpcMux))
}
