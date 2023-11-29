package methods

import (
	"github.com/nats-io/stan.go"
	"log"
)

func InitNatsBridge() (stan.Conn, error) {
	nc, err := stan.Connect("cluster", "subscriber", stan.NatsURL("nats://stan-server:4222"))
	if err != nil {
		log.Fatalln(err)
	}
	return nc, err
}
