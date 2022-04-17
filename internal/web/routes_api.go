package web

import (
	"app/internal/service"

	"github.com/gin-gonic/gin"
)

func registerAPI(r *gin.RouterGroup, s *service.Service) {
	payment := r.Group("/payment")
	{
		payment.GET("/", ttt)
		payment.POST("/", createPaymont(s))
	}
}
