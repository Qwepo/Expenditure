package web

import (
	"app/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ttt(c *gin.Context) {
	c.JSON(200, gin.H{"sms": "babagi"})
}

func createPaymont(s *service.Service, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		var resp service.PaymentFullRequest
		if err := c.BindJSON(&resp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"bad Request": err.Error()})
			return
		}
		fmt.Println(resp.ExpenditureItemName)
		id, err := s.Payment.PaymentCreate(&resp, log)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"bad Request": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": id})

	}

}
