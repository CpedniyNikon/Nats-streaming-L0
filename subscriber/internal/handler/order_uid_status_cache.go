package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"publisher/internal/models"
	"publisher/internal/utils/methods"
)

func (h *Handler) OrderUIdStatusCache(c *gin.Context) {
	orderUId := c.Param("orderUId")

	db, err := methods.InitDbConnection()
	if err != nil {
		panic(err)
	}

	var order models.Order
	db.First(&order, "order_uid = ?", orderUId)

	var Payment models.Payment
	db.First(&Payment, "order_id = ?", order.ID)

	var Delivery models.Delivery
	db.First(&Delivery, "order_id = ?", order.ID)

	var Items []models.Item
	db.Find(&Items, "order_id = ?", order.ID)

	order.Payment = Payment
	order.Delivery = Delivery
	order.Items = Items

	jsonData, err := json.MarshalIndent(order, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"jsonString": string(jsonData),
	})
}
