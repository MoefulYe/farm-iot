package handler

import (
	"fmt"
	"github.com/MoefulYe/farm-iot/http-server/db"
	"github.com/MoefulYe/farm-iot/http-server/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"strings"
)

// Keepaliveuuidhandler PingExample godoc
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
// @Success 200 {object} models.Resp[models.Info] "xiangying"
// @Router /cow/keep-alive/{uuid} [get]
func Keepaliveuuidhandler(c *gin.Context) {
	//page := c.Request.URL.Query().Get("page")
	//size := c.Request.URL.Query().Get("size")
	field := c.Request.URL.Query().Get("field")
	uuid1 := c.Param("uuid")
	var data []models.KeepalivePackage
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
		fmt.Printf("%v", err)
		c.JSON(200, models.ResponseList{Code: 1, Msg: "db.query err", Data: ""})
		return
	}
	c.JSON(200, models.ResponseList{Code: 0, Msg: "keepalivePackage ok", Data: data})
}
