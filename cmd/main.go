package main

import (
	todo "GoRestAPI"
	"GoRestAPI/pkg/dao"
	"GoRestAPI/pkg/service"
	"GoRestAPI/pkg/view"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error occured reading config file: %s", err.Error())
	}

	db, err := dao.NewPostgresDB(dao.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("Error occured connection to database: %s", err.Error())
	}

	daos := dao.NewDao(db)
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
