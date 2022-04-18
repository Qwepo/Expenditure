package web

import (
	"app/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ttt(c *gin.Context) {
	c.JSON(200, gin.H{"sms": "babagi"})
}

func createPaymont(s *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		var resp service.PaymentFullRequest
		if err := c.BindJSON(&resp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"bad Request": err.Error()})
			return
		}
		id, err := s.Payment.PaymentCreate(&resp)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"bad Request": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": id})

	}

}
