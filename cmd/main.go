package main

import (
	config "app/configs"
	"app/pkg/db"
	"app/pkg/hanlder"
	"app/pkg/service"
	"app/server"
	"log"
)

func main() {
	conf := config.LoadConfig()
	db, err := db.Open(&conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	service := service.NewService(db)
	handler := hanlder.NewHandler(service)

	srv := new(server.Server)
	if err := srv.Run(&conf); err != nil {
		log.Fatal(err)
	}

}
