package controllers

import (
	// "log"
	// "strconv"

	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"goapi/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetPayment(c *gin.Context) {
	var payment []models.Payment
	_, err := dbmap.Select(&payment, "USE [Interoperabilidad] SELECT * FROM NotificationMOLPayment")

	if err == nil {
		c.JSON(200, payment)
	} else {
		c.JSON(404, gin.H{"error": "payment not found"})
	}

}

func Healthcheck(c *gin.Context) {

	errDbFacthos := infra.CheckDB("FACTHOS")
	errDbJDE := infra.CheckDB("JDE")

	var sDbFacthos, sDbJDE string

	if errDbFacthos != nil {
		sDbFacthos = errDbFacthos.Error()
	} else {
		sDbFacthos = "OK"
	}

	if errDbJDE != nil {
		sDbJDE = errDbJDE.Error()
	} else {
		sDbJDE = "OK"
	}

	if errDbFacthos != nil || errDbJDE != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"FACTHOS": sDbFacthos,
			"JDE":     sDbJDE,
			"time":    time.Now(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"FACTHOS": sDbFacthos,
			"JDE":     sDbJDE,
			"time":    time.Now(),
		})
	}
	return

}



func PostPayment(c *gin.Context) {

	w := c.Writer
	r := c.Request

	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		errorResponse(w, "Content type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	var newPayment models.Payment
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newPayment)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}

	ctx := context.Background()
	DBConection := initDb()
	tsql := fmt.Sprintf("USE [Interoperabilidad] INSERT INTO [dbo].[NotificationMOLPayment]([Key],[External_Reference],[Status],[Amount]) VALUES (@Key, @External_reference, @Status, @Amount);")
	result, err2 := DBConection.ExecContext(
		ctx,
		tsql,
		sql.Named("Key", newPayment.Key),
		sql.Named("External_reference", newPayment.External_reference),
		sql.Named("Status", newPayment.Status),
		sql.Named("Amount", newPayment.Rate),
	)
	if err2 != nil {
		errorResponse(w, "Error inserting new row: "+err2.Error(), http.StatusBadRequest)
		return
	} else if result != nil {
		errorResponse(w, "Successfully added new row", http.StatusCreated)
	}
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}