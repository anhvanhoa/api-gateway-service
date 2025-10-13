package main

import (
	"api-gateway/src/bootstrap"
	"api-gateway/src/handler"
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	var env bootstrap.Env
	bootstrap.NewEnv(&env)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Tạo Gin router
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// Tạo handler và đăng ký service routes trước
	baseHandler := handler.NewBaseHandler(&env, router, ctx)
	handlerMap := handler.GetServiceHandlers()
	for _, service := range env.Services {
		serviceHandler, ok := handlerMap[service.Folder]
		if !ok {
			continue
		}
		baseHandler.AddService(handler.NewService(service.Name, service.Route, service.Host, service.Port, serviceHandler.Swagger, serviceHandler.Handler))
	}

	fmt.Println("API Gateway running on http://localhost:" + strconv.Itoa(env.Port))
	router.Run(":" + strconv.Itoa(env.Port))
}
