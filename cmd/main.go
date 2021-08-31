package main

import (
	todo "GoRestAPI"
	"GoRestAPI/pkg/dao"
	"GoRestAPI/pkg/service"
	"GoRestAPI/pkg/view"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error occured reading config file: %s", err.Error())
	}

	daos := dao.NewDao()
	services := service.NewService(daos)
	handler := view.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
