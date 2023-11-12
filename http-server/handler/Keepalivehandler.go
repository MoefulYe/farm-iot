package handler

import (
	"fmt"
	"github.com/MoefulYe/farm-iot/http-server/db"
	"github.com/MoefulYe/farm-iot/http-server/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"strings"
)

var field2query = map[string]string{
	"weight": `from(bucket:"farm-iot")
	|> range(start: -1m)
	|> filter(fn: (r) => r._measurement == "cow")
	|> filter(fn: (r) => r._field == "weight")`,
	"health": `from(bucket:"farm-iot")
	|> range(start: -5m)
	|> filter(fn: (r) => r._measurement == "cow")
	|> filter(fn: (r) => r._field == "health")`,
	"latitude": `from(bucket:"farm-iot")
	|> range(start: -5m)
	|> filter(fn: (r) => r._measurement == "cow")
	|> filter(fn: (r) => r._field == "latitude")`,
	"longitude": `from(bucket:"farm-iot")
	|> range(start: -5m)
	|> filter(fn: (r) => r._measurement == "cow")
	|> filter(fn: (r) => r._field == "longitude")`,
}

// Keepalivehandler PingExample godoc
// @Tags  get Keep-alive package
// @Summary keepalive package
// @Schemes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "token"
// @Param       page    query     int  false  "need page"
// @Param       size    query     int  false  "need size"
// @Param       field    query     string  true  "field name"
// @Success 200 {object} models.Resp[models.Info] "xiangying"
// @Router /cow/keep-alive [get]
func Keepalivehandler(c *gin.Context) {
	//page := c.Request.URL.Query().Get("page")
	//size := c.Request.URL.Query().Get("size")
	field := c.Request.URL.Query().Get("field")
	var data []models.KeepalivePackage
	query, ok := field2query[field]
	if !ok {
		return
	}
	result, err := db.QueryAPI.Query(context.Background(), query)
	if err == nil {
		// Use Next() to iterate over query result lines
		for result.Next() {
			k := strings.Split(result.Record().String(), ",")
			time := strings.Split(k[4], ":")
			value := strings.Split(k[5], ":")
			uuid := strings.Split(k[8], ":")
			c := models.KeepalivePackage{uuid[1], value[1], time[1]}
			//fmt.Printf("%v\n", c)
			data = append(data, c)
		}
		if result.Err() != nil {
			c.JSON(200, models.ResponseList{Code: 1, Msg: "db.query err", Data: ""})
			fmt.Printf("Query error: %s\n", result.Err().Error())
			return
		}
	} else {
		c.JSON(200, models.ResponseList{Code: 1, Msg: "db.query err", Data: ""})
		fmt.Printf("%v", err)
		return
	}
	c.JSON(200, models.ResponseList{Code: 0, Msg: "keepalivePackage ok", Data: data})
}
