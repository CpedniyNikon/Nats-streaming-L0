package methods

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"gorm.io/gorm"
	"publisher/internal/models"
	"publisher/internal/utils/structs/cache"
)

func SubscribeModelOrder(db *gorm.DB, nc stan.Conn, c *cache.Cache) stan.Subscription {
	subscriber, err := nc.Subscribe("Model-order", func(m *stan.Msg) {
		var restoredOrder models.Order
		err := json.Unmarshal(m.Data, &restoredOrder)
		if err != nil {
			fmt.Println(err)
		}
		c.Add(restoredOrder)
		//fmt.Println(restoredOrder)
		db.Create(&restoredOrder)
	})
	if err != nil {
		fmt.Println(err)
	}

	return subscriber
}
