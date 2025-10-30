package main

import (
	"api-gateway/src/bootstrap"
	"api-gateway/src/handler"
	"context"
	"fmt"
	"net/http"
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
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	router := gin.New()
	router.Use(corsConfig)
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	swaggerHandler := handler.NewSwaggerHandler()
	baseHandler := handler.NewBaseHandler(&env, router, ctx, swaggerHandler)
	handlerMap := handler.GetServiceHandlers()
	for _, service := range env.Services {
		serviceHandler, ok := handlerMap[service.Folder]
		if !ok {
			continue
		}
		if service.Folder == env.AuthService.Folder {
			handler.HandleAuth(ctx, &env, router, swaggerHandler, serviceHandler)
			continue
		}
		baseHandler.AddService(handler.NewService(service.Name, service.Route, service.Host, service.Port, serviceHandler.Swagger, serviceHandler.Handler))
	}

	fmt.Println("API Gateway running on http://localhost:" + strconv.Itoa(env.Port))
	router.Run(":" + strconv.Itoa(env.Port))
}
