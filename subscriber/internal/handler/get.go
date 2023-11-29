package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"publisher/internal/models"
	"publisher/internal/utils/methods"
)

func (h *Handler) Get(c *gin.Context) {

	db, err := methods.InitDbConnection()
	if err != nil {
		panic(err)
	}

	var Orders []models.Order
	db.Find(&Orders)

	for i := 0; i < len(Orders); i++ {

		var Delivery models.Delivery
		db.First(&Delivery, "order_id = ?", Orders[i].ID)

		var Payment models.Payment
		db.First(&Payment, "order_id = ?", Orders[i].ID)

		var Items []models.Item
		db.Find(&Items, "order_id = ?", Orders[i].ID)

		Orders[i].Delivery = Delivery
		Orders[i].Payment = Payment
		Orders[i].Items = Items
	}

	jsonData, err := json.MarshalIndent(Orders, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"jsonString": string(jsonData),
	})
}
