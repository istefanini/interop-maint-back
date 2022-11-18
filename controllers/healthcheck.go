package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/istefanini/goapi/infra"
)

func Healthcheck(c *gin.Context) {
	errDbPayment := infra.CheckDB()
	var sDbPayment string
	if errDbPayment != nil {
		sDbPayment = errDbPayment.Error()
	} else {
		sDbPayment = "OK"
	}
	if errDbPayment != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"INTEROPERABILIDAD": sDbPayment,
			"time":              time.Now(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"INTEROPERABILIDAD": sDbPayment,
			"time":              time.Now(),
		})
	}
	return
}
