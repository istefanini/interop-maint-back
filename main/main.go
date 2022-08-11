package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/istefanini/goapi/infra"
	"github.com/istefanini/goapi/routes"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	infra.DbPayment, _ = infra.ConnectDB()
	defer infra.DbPayment.Close()
	////CHECK CONNEXION DB
	err := infra.CheckDB()
	infra.FailOnError(err, "Failed to connect to SQLServer")
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
		serverPort := os.Getenv("API_PORT")
		_ = r.Run(":" + serverPort)
	}()
}
