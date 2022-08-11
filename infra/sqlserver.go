package infra

import (
	"context"
	"database/sql"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	DbPayment *sql.DB
)

func ConnectDB() (conection *sql.DB) {
	Driver := os.Getenv("DB_DRIVER")
	Username := os.Getenv("DB_USERNAME")
	Password := os.Getenv("DB_PASSWORD")
	Host := os.Getenv("DB_HOST")
	Instance := os.Getenv("DB_INSTANCE")
	Database := os.Getenv("DB_DATABASE")
	Encrypt := os.Getenv("DB_ENCRYPT")

	conection, err := sql.Open(Driver, Driver+"://"+Username+":"+Password+"@"+Host+"/"+Instance+"?"+"database="+Database+"&"+"Encrypt="+Encrypt+"")
	if err != nil {
		panic(err.Error())
	}
	return conection
}

func CheckDB() error {

	var err error
	ctx := context.Background()
	err = DbPayment.PingContext(ctx)

	if err != nil {
		return err
	}

	return err
}
