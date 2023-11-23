package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"publisher/internal/handler"
	"publisher/pkg/server"
)

func main() {
	fmt.Println("server started")
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	handlers := handler.NewHandler()

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("app_port"), handlers.InitAuthRoutes()); err != nil {
		log.Fatal("error occurred while auth service " + err.Error())
	}
}

func InitConfig() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	return viper.ReadInConfig()
}
