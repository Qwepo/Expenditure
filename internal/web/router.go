package web

import (
	"app/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewRouter(s *service.Service, log *logrus.Logger) *gin.Engine {
	router := gin.New()

	registerRoutes(router, s, log)
	return router

}
