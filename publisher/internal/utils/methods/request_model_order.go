package methods

import (
	"github.com/nats-io/stan.go"
	"log"
)

func RequestModelOrder(nc stan.Conn, data []byte) {
	if err := nc.Publish("Model-order", data); err != nil {
		log.Println(err)
	}
}
