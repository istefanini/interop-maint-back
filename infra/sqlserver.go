package infra

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	DbPayment *sql.DB
)

func ConnectDB() (*sql.DB, error) {
	Driver := os.Getenv("DB_DRIVER")
	Username := os.Getenv("DB_USERNAME")
	Password := os.Getenv("DB_PASSWORD")
	Host := os.Getenv("DB_HOST")
	Instance := os.Getenv("DB_INSTANCE")
	Database := os.Getenv("DB_DATABASE")
	Encrypt := os.Getenv("DB_ENCRYPT")
	return sql.Open(Driver, Driver+"://"+Username+":"+Password+"@"+Host+"/"+Instance+"?"+"database="+Database+"&"+"Encrypt="+Encrypt+"")
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

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}
