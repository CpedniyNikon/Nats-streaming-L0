package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"publisher/cmd/methods"
	"publisher/internal/handler"
	"publisher/internal/models"
	"publisher/pkg/server"
	"time"
)

func main() {
	time.Sleep(2 * time.Second)
	fmt.Println("server started")
	if err := methods.InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	nc, err := methods.InitNatsBridge()
	if err != nil {
		log.Fatalln("error initializing nats bridge: %s", err.Error())
	}

	db, err := methods.InitDbConnection()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Delivery{})
	db.AutoMigrate(&models.Payment{})
	db.AutoMigrate(&models.Order{})

	methods.SubscribeModelOrder(nc)

	handlers := handler.NewHandler()

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("app_port"), handlers.InitAuthRoutes()); err != nil {
		log.Fatal("error occurred while auth service " + err.Error())
	}
}
