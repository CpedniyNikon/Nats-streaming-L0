package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	routes *gin.Engine
}

func NewHandler() (h *Handler) {
	return &Handler{}
}

func (h *Handler) InitAuthRoutes() *gin.Engine {
	h.routes = gin.New()
	h.routes.Use(cors.Default())
	h.routes.LoadHTMLGlob("templates/*")
	auth := h.routes.Group("/orders")
	{
		auth.GET("/get", h.Get)
		auth.GET("/:orderId/status", h.OrderIdStatus)
	}
	return h.routes
}
