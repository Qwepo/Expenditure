package web

import (
	"app/internal/service"

	"github.com/gin-gonic/gin"
)

func NewRouter(s *service.Service) *gin.Engine {
	router := gin.New()

	registerRoutes(router, s)
	return router

}
