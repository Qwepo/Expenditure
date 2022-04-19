package web

import (
	"app/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func registerAPI(r *gin.RouterGroup, s *service.Service, log *logrus.Logger) {
	payment := r.Group("/payment")
	{
		payment.GET("/", ttt)
		payment.POST("/", createPaymont(s, log))
	}
}
