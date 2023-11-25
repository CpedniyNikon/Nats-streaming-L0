package methods

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"publisher/internal/models"
)

func SubscribeModelOrder(nc *nats.Conn) {
	nc.Subscribe("Model-order", func(m *nats.Msg) {
		var restoredOrder models.Order
		err := json.Unmarshal(m.Data, &restoredOrder)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(restoredOrder)

		response := []byte("response")
		nc.Publish(m.Reply, response)
	})
}
