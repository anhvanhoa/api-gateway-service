package main

import (
	"api-gateway/src/bootstrap"
	"api-gateway/src/handler"
	"context"
	"fmt"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	var env bootstrap.Env
	bootstrap.NewEnv(&env)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// CORS
	corsConfig := cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	router := gin.New()
	router.Use(corsConfig)
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(handler.CookieMiddleware())

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
