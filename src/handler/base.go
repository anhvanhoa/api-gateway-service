package handler

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	RegisterIoTDevice(
		ctx context.Context,
		router *gin.Engine,
	)
}

type BaseHandler struct{}

func NewBaseHandler() Handler {
	return &BaseHandler{}
}
