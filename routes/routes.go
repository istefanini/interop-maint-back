package routes

import (
	"interop-maint-back/controllers"

	"github.com/gin-gonic/gin"
)

func CreateRoutes(r *gin.Engine) {

	v1 := r.Group("")
	v1.Use()
	{
		v1.GET("/healthcheck", controllers.Healthcheck)
		v1.POST("/getLogs", controllers.getLogs)
	}
}
