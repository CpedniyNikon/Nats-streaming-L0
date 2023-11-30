package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"subscriber/internal/utils/methods"
	"time"
)

func main() {
	time.Sleep(10 * time.Second)
	if err := methods.InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	_, err := os.ReadFile(viper.GetString("file_path"))
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	nc, err := methods.InitNatsBridge()
	if err != nil {
		log.Fatalln("error initializing nats bridge: %s", err.Error())
	}

	for i := 0; i < 1000; i++ {
		order := methods.GenerateMockOrder()
		b, err := json.Marshal(order)
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
		methods.RequestModelOrder(nc, b)
	}

	// Close connection
	nc.Close()
}
