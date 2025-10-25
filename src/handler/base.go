package handler

import (
	"api-gateway/src/bootstrap"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Handler interface {
	AddService(service *Service)
	CloseConnections()
}

type BaseHandler struct {
	env            *bootstrap.Env
	router         *gin.Engine
	ctx            context.Context
	services       map[string]*grpc.ClientConn
	swaggerHandler *SwaggerHandler
}

func NewBaseHandler(env *bootstrap.Env, router *gin.Engine, ctx context.Context) Handler {
	return &BaseHandler{
		env:      env,
		router:   router,
		ctx:      ctx,
		services: make(map[string]*grpc.ClientConn),
	}
}

func (h *BaseHandler) AddService(service *Service) {
	// Tạo gRPC connection
	conn, err := grpc.NewClient(
		service.Host+":"+service.Port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("Không thể kết nối đến gRPC server %s: %v", service.Name, err)
		return
	}

	// Lưu trữ connection để sử dụng sau này
	h.services[service.Name] = conn

	// Tạo gRPC gateway mux
	grpcMux := runtime.NewServeMux(
		runtime.WithErrorHandler(CustomErrorHandler),
	)

	err = service.Handler(h.ctx, grpcMux, conn)
	if err != nil {
		log.Printf("Không thể đăng ký gRPC service handler cho %s: %v", service.Name, err)
		return
	}

	log.Printf("Đã đăng ký service: %s tại route: %s", service.Name, service.Route)

	// Đăng ký gRPC gateway vào Gin router
	fileJSON := service.Route + ".json"
	h.router.GET(fileJSON, func(c *gin.Context) {
		h.swaggerHandler.ServeSwaggerJSON(c, service.Swagger)
	})
	h.router.GET("/swagger"+service.Route, func(c *gin.Context) {
		h.swaggerHandler.ServeSwaggerUI(c, fileJSON)
	})
	h.router.Any(service.Route, gin.WrapH(grpcMux))
	h.router.Any(service.Route+"/*path", gin.WrapH(grpcMux))
}

type Service struct {
	Name    string
	Route   string
	Host    string
	Port    string
	Swagger string
	Handler func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
}

func NewService(name string, route string, host string, port string, swagger string, handler func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error) *Service {
	return &Service{Name: name, Route: route, Host: host, Port: port, Swagger: swagger, Handler: handler}
}

func (h *BaseHandler) CloseConnections() {
	for name, conn := range h.services {
		if err := conn.Close(); err != nil {
			log.Printf("Lỗi khi đóng connection cho service %s: %v", name, err)
		} else {
			log.Printf("Đã đóng connection cho service: %s", name)
		}
	}
}
