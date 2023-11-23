package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
	"log"
	"subscriber/internal/models"
	"subscriber/internal/utils"
)

func main() {
	db, err := sql.Open("postgres", utils.ConnectionInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	nc, err := nats.Connect("nats://stan-server:4222")
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()

	nc.Subscribe("Model-order", func(m *nats.Msg) {
		var restoredOrder models.Order
		err = json.Unmarshal(m.Data, &restoredOrder)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(restoredOrder)

		response := []byte("response")
		nc.Publish(m.Reply, response)
	})
	for {

	}
}
