package methods

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"gorm.io/gorm"
	"publisher/internal/models"
)

func SubscribeModelOrder(db *gorm.DB, nc stan.Conn) stan.Subscription {
	subscriber, err := nc.Subscribe("Model-order", func(m *stan.Msg) {
		var restoredOrder models.Order
		err := json.Unmarshal(m.Data, &restoredOrder)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(restoredOrder)
		db.Create(&restoredOrder)
	})
	if err != nil {
		fmt.Println(err)
	}
	return subscriber
}
