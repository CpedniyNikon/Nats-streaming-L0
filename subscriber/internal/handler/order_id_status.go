package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"publisher/internal/models"
	"publisher/internal/utils/methods"
)

func (h *Handler) OrderIdStatus(c *gin.Context) {
	orderId := c.Param("orderId")

	db, err := methods.InitDbConnection()
	if err != nil {
		panic(err)
	}

	var order models.Order
	db.First(&order, orderId)

	var Payment models.Payment
	db.First(&Payment, "order_id = ?", orderId)

	var Delivery models.Delivery
	db.First(&Delivery, "order_id = ?", orderId)

	var Items []models.Item
	db.Find(&Items, "order_id = ?", orderId)

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
