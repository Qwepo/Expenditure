package web

import (
	"app/internal/service"

	"github.com/gin-gonic/gin"
)

func registerRoutes(r *gin.Engine, s *service.Service) {
	api := r.Group("/api")
	registerAPI(api, s)
}
