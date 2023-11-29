package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"publisher/internal/handler"
	"publisher/internal/models"
	methods2 "publisher/internal/utils/methods"
	"publisher/pkg/server"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	fmt.Println("server started")
	if err := methods2.InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	nc, err := methods2.InitNatsBridge()
	if err != nil {
		log.Fatalln("error initializing nats bridge: %s", err.Error())
	}

	db, err := methods2.InitDbConnection()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Delivery{})
	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Payment{})

	subscriber := methods2.SubscribeModelOrder(db, nc)

	handlers := handler.NewHandler()

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("app_port"), handlers.InitAuthRoutes()); err != nil {
		log.Fatal("error occurred while auth service " + err.Error())
	}

	subscriber.Unsubscribe()
}
