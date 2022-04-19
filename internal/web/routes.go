package web

import (
	"app/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func registerRoutes(r *gin.Engine, s *service.Service, log *logrus.Logger) {
	api := r.Group("/api")
	registerAPI(api, s, log)
}
