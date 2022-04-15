package hanlder

import (
	"app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

// func (h *Handler) initRoutes() *gin.Engine {
// 	return nil

// }
