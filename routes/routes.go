package routes

import (
	"goapi/controllers"

	"github.com/gin-gonic/gin"
)

func CreateRoutes(r *gin.Engine) {

	v1 := r.Group("/payment/api/v1")
	v1.Use()
	{
		v1.GET("/healthcheck", controllers.Healthcheck)
		v1.GET("/payments", controllers.GetPayment)
		v1.POST("/notificaction-mol-payment", controllers.PostPayment)
	}
}
