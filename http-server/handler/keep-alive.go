package handler

import (
	"fmt"
	"github.com/MoefulYe/farm-iot/http-server/db"
	"github.com/MoefulYe/farm-iot/http-server/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"strings"
)

// GetKeepaliveByUuid PingExample godoc
// @Tags  get Keep-alive package
// @Summary keepalive by uuid
// @Schemes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "token"
// @Param uuid path string    true  "uuid"
// @Param       field    query     string  true  "field name"
// @Param       start    query     int  false  "start time"
// @Param       end    query     int  false  "end time"
// @Success 200 {object} models.Resp[models.DeviceInfo] "xiangying"
// @Router /cow/keep-alive/{uuid} [get]
func GetKeepaliveByUuid(c *gin.Context) {
	//page := c.Request.URL.Query().Get("page")
	//size := c.Request.URL.Query().Get("size")
	field := c.Request.URL.Query().Get("field")
	uuid1 := c.Param("uuid")
	var data []models.KeepAlive
	Query := `from(bucket:"farm-iot")
				|> range(start: -1m)
				|> filter(fn: (r) => r._measurement == "cow")
				|> filter(fn: (r) => r._field == "` + field + "\")\n" +
		`|> filter(fn: (r) => r["uuid"] =="` + uuid1 + "\")"
	fmt.Printf("%v", Query)
	result, err := db.QueryAPI.Query(context.Background(), Query)
	if err == nil {
		// Use Next() to iterate over query result lines
		for result.Next() {
			k := strings.Split(result.Record().String(), ",")
			time := strings.Split(k[4], ":")
			value := strings.Split(k[5], ":")
			uuid := strings.Split(k[8], ":")
			c := models.KeepAlive{uuid[1], value[1], time[1]}
			//fmt.Printf("%v\n", c)
			data = append(data, c)
		}
		if result.Err() != nil {
			c.JSON(
				200,
				models.ResponseList{Code: 1, Msg: "db.query err", Data: ""},
			)
			fmt.Printf("Query error: %s\n", result.Err().Error())
			return
		}
	} else {
		fmt.Printf("%v", err)
		c.JSON(200, models.ResponseList{Code: 1, Msg: "db.query err", Data: ""})
		return
	}
	c.JSON(
		200,
		models.ResponseList{Code: 0, Msg: "keepalivePackage ok", Data: data},
	)
}

var field1query = map[string]string{
	"weight": `from(bucket:"farm-iot")
	|> range(start: -2m)
	|> filter(fn: (r) => r._measurement == "cow")
	|> filter(fn: (r) => r._field == "weight")`,
	"health": `from(bucket:"farm-iot")
	|> range(start: -6m)
	|> filter(fn: (r) => r._measurement == "cow")
	|> filter(fn: (r) => r._field == "health")`,
	"latitude": `from(bucket:"farm-iot")
	|> range(start: -6m)
	|> filter(fn: (r) => r._measurement == "cow")
	|> filter(fn: (r) => r._field == "latitude")`,
	"longitude": `from(bucket:"farm-iot")
	|> range(start: -6m)
	|> filter(fn: (r) => r._measurement == "cow")
	|> filter(fn: (r) => r._field == "longitude")`,
}

// GetKeepalive PingExample godoc
// @Tags  get Keep-alive package
// @Summary keepalive package
// @Schemes
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "token"
// @Success 200 {object} models.Resp[models.KeepAlive] "success"
// @Router /cow/keep-alive [get]
func GetKeepalive(c *gin.Context) {
	//page := c.Request.URL.Query().Get("page")
	//size := c.Request.URL.Query().Get("size")

	func buildQuery(query models.RangeQuery) string {

	}
	field := c.Request.URL.Query().Get("field")
	var data []models.KeepAlive
	query, ok := field1query[field]
	if !ok {
		return
	}
	result, err := db.QueryAPI.Query(context.Background(), query)
	if err == nil {
		// Use Next() to iterate over query result lines
		for result.Next() {
			k := strings.Split(result.Record().String(), ",")
			time := strings.Split(k[3], ":")
			value := strings.Split(k[4], ":")
			uuid := strings.Split(k[7], ":")
			c := models.KeepAlive{uuid[0], value[1], time[1]}
			//fmt.Printf("%v\n", c)
			data = append(data, c)
		}
		if result.Err() != nil {
			c.JSON(
				199,
				models.ResponseList{Code: 1, Msg: "db.query err", Data: ""},
			)
			fmt.Printf("Query error: %s\n", result.Err().Error())
			return
		}
	} else {
		c.JSON(199, models.ResponseList{Code: 1, Msg: "db.query err", Data: ""})
		fmt.Printf("%v", err)
		return
	}
	c.JSON(
		199,
		models.ResponseList{Code: 0, Msg: "keepalivePackage ok", Data: data},
	)
}
