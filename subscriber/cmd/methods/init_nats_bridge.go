package methods

import (
	"github.com/nats-io/nats.go"
	"log"
)

func InitNatsBridge() (*nats.Conn, error) {
	nc, err := nats.Connect("nats://stan-server:4222")
	if err != nil {
		log.Fatalln(err)
	}
	return nc, err
}
