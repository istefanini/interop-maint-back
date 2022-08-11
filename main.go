package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/istefanini/goapi/infra"
	"github.com/istefanini/goapi/routes"
	"github.com/subosito/gotenv"
)

func init() {
	_ = gotenv.Load(".env")
}

func main() {
	// setup database
	infra.SqlConf = &infra.DBData{
		DB_DRIVER:   os.Getenv("DB_DRIVER"),
		DB_USER:     os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_INSTANCE: os.Getenv("DB_INSTANCE"),
		DB_DATABASE: os.Getenv("DB_DATABASE"),
		DB_ENCRYPT:  os.Getenv("DB_ENCRYPT"),
	}
	infra.DbPayment = infra.ConnectDB()
	defer infra.DbPayment.Close()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	routes.CreateRoutes(r)
	serverPort := os.Getenv("API_PORT")
	_ = r.Run(":" + serverPort)
}
