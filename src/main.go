package main

import (
	"api-gateway/src/handler"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Tạo Gin router
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	// Tạo handler và đăng ký routes
	handler := handler.NewBaseHandler()
	handler.RegisterIoTDevice(ctx, router)

	fmt.Println("API Gateway running on http://localhost:8080")
	router.Run(":8080")
}
