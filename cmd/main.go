package main

import (
	config "app/internal/configs"
	"app/internal/db"
	"app/internal/logging"
	"app/internal/server"
	"app/internal/service"
	"app/internal/web"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logging.Init()
	conf := config.LoadConfig()
	db, err := db.Open(&conf)
	chek(log, err)

	service := service.NewService(db, log)
	router := web.NewRouter(service, log)
	srv := new(server.Server)
	err = srv.Run(&conf, router)
	chek(log, err)
}

func chek(log *logrus.Logger, err error) {
	if err != nil {
		log.Fatal(err)
	}
}
