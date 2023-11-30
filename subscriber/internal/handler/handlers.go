package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"publisher/pkg/service"
)

type Handler struct {
	routes  *gin.Engine
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitAuthRoutes() *gin.Engine {
	h.routes = gin.New()
	h.routes.Use(cors.Default())
	h.routes.LoadHTMLGlob("templates/*")
	orders := h.routes.Group("/orders")
	{
		orders.GET("/get", h.Get)
		orders.GET("/:orderId/status", h.OrderIdStatus)
		orders.GET("/cache/:orderUId/status", h.OrderUIdStatusCache)
	}

	return h.routes
}
