package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/istefanini/goapi/infra"
)

type Log struct {
	EventID    string  `json:"EventID"`
	SysFechaC  string  `json:"sysFechaC"`
	Estado     float64 `json:"Estado"`
	LogProceso string  `json:"LogProceso"`
	MsgFinal   string  `json:"MsgFinal"`
}

var logs []Log

func getLogs(c *gin.Context) {
	logs := GetLogs(c)
	if logs == nil || len(logs) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, logs)
	}
}

func GetLogs(c *gin.Context) []Log {
	results, err := infra.ConnectDB().Query("SELECT TOP (10) [Transaccion_ID],[EventID],[sysFechaC],[Estado],[LogProceso],[MsgFinal] FROM [Interoperabilidad].[dbo].[MQMsgDisparados]")
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	for results.Next() {
		var log Log
		err = results.Scan(&log.EventID, &log.SysFechaC, &log.Estado, &log.LogProceso, &log.MsgFinal)
		if err != nil {
			panic(err.Error())
		}
		logs = append(logs, log)
		//fmt.Println("product.code :", prod.Code+" : "+prod.Name)
	}
	fmt.Println(logs)
	return logs
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
