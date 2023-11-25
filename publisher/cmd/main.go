package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"log"
	"os"
	"subscriber/internal/models"
	"time"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

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

	nc, err := InitNatsBridge()
	if err != nil {
		log.Fatalln("error initializing nats bridge: %s", err.Error())
	}

	for {
		RequestModelOrder(nc, reqBodyBytes.Bytes())
		time.Sleep(5 * time.Second)
	}
}

func InitConfig() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	return viper.ReadInConfig()
}

func InitNatsBridge() (*nats.Conn, error) {
	nc, err := nats.Connect("nats://stan-server:4222")
	if err != nil {
		log.Fatalln(err)
	}
	return nc, err
}

func RequestModelOrder(nc *nats.Conn, data []byte) {
	msg, err := nc.Request("Model-order", data, 100*time.Millisecond)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(msg.Data))
	}
}
