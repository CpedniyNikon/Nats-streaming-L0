package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"publisher/internal/handler"
	"publisher/internal/models"
	"publisher/internal/utils/methods"
	"publisher/internal/utils/structs/cache"
	"publisher/pkg/server"
	"publisher/pkg/service"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
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

	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Delivery{})
	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Payment{})

	Cache := cache.NewCache()
	Service := service.NewService(Cache)

	subscriber := methods.SubscribeModelOrder(db, nc, Cache)

	handlers := handler.NewHandler(Service)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("app_port"), handlers.InitAuthRoutes()); err != nil {
		log.Fatal("error occurred while auth service " + err.Error())
	}

	subscriber.Unsubscribe()
}
