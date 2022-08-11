package main

import (
	"goapi/mappings"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/istefanini/restapi-gin/routes"
	cors "github.com/itsjamie/gin-cors"
)

func main() {

	infra.DbPayment, _ = infra.ConnectDB()
	defer infra.DbPayment.Close()

	mappings.CreateUrlMappings()
	// Listen and server on 0.0.0.0:8080
	mappings.Router.Run(":8080")

	go func() {
		//GIN FRAMEWORK
		r := gin.Default()
		//Set up CORS middleware options
		config := cors.Config{
			Origins:         "*",
			RequestHeaders:  "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
			Methods:         "POST, GET, OPTIONS, PUT, DELETE",
			Credentials:     false,
			ValidateHeaders: false,
			MaxAge:          1 * time.Minute,
		}
		r.Use(cors.Middleware(config))
		routes.CreateRoutes(r)
		serverPort := os.Getenv("API_SERVER_PORT")
		_ = r.Run(":" + serverPort)
	}()

}
