package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"publisher/internal/models"
	"time"
)

func (h *Handler) get(c *gin.Context) {
	fmt.Println("get request called")

	content, err := os.ReadFile(viper.GetString("file_path"))
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var order models.Order
	err = json.Unmarshal(content, &order)
	if err != nil {
		fmt.Println(err)
	}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(order)

	nc, err := nats.Connect("nats://stan-server:4222")
	if err != nil {
		log.Println(err)
	} else {
		defer nc.Close()
	}

	msg, err := nc.Request("Model-order", reqBodyBytes.Bytes(), 100*time.Millisecond)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(msg.Data))
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
