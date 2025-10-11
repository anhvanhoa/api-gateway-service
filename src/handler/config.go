package handler

import (
	"context"
	"log"

	proto_system_configuration "github.com/anhvanhoa/sf-proto/gen/system_configuration/v1"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (h *BaseHandler) RegisterSystemConfiguration(
	ctx context.Context,
	router *gin.Engine,
) {
	// Tạo gRPC connection
	conn, err := grpc.NewClient(
		"localhost:50057",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Không thể kết nối đến gRPC server: %v", err)
	}

	// Tạo gRPC gateway mux
	grpcMux := runtime.NewServeMux(
		runtime.WithErrorHandler(CustomErrorHandler),
	)
	if err := proto_system_configuration.RegisterSystemConfigurationServiceHandler(ctx, grpcMux, conn); err != nil {
		log.Fatalf("Không thể đăng ký handler: %v", err)
	}

	router.Any("/system-configurations/*path", gin.WrapH(grpcMux))
}
