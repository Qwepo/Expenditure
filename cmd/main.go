package main

import (
	config "app/internal/configs"
	"app/internal/db"
	"app/internal/server"
	"app/internal/service"
	"app/internal/web"
	"log"
)

func main() {
	conf := config.LoadConfig()
	db, err := db.Open(&conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	service := service.NewService(db)
	router := web.NewRouter(service)
	srv := new(server.Server)
	if err := srv.Run(&conf, router); err != nil {
		log.Fatal(err)
	}

}
