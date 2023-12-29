package handler

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/MoefulYe/farm-iot/http-server/db"
	"github.com/MoefulYe/farm-iot/http-server/logger"
	"github.com/MoefulYe/farm-iot/http-server/models"
	"github.com/MoefulYe/farm-iot/http-server/utils"
	"github.com/gin-gonic/gin"
)

// GetHeartbeatByUuid
// @Tags heartbeat
// @Summary heartbeat by uuid
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "jwt"
// @Param uuid path string true  "uuid"
// @Param query-params query models.RangeQuery true "范围, 查询字段"
// @Success 200 {object} models.Resp[[]models.CowInfo] "success"
// @Failure 400 {object} models.Msg "failure"
// @Router /cow/heartbeat/{uuid} [get]
func GetHeartbeatByUuid(c *gin.Context) {
	var query models.RangeQuery
	if err := c.BindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, err.Error()))
		return
	}
	uuid := c.Param("uuid")
	fields := strings.Split(query.Fields, ",")
	ranges := ranges(query.Start, query.Stop)
	filter := fieldFilter(fields)
	flux := buildFlux4SelectedUuid(ranges, filter, uuid)
	res, err := db.QueryAPI.Query(c, flux)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, err.Error()))
		logger.Logger.Fatalw(err.Error())
		return
	}
	var ret []models.HeartBeat
	for res.Next() {
		var data models.HeartBeat
		value := reflect.ValueOf(&data).Elem()
		data.Id = res.Record().ValueByKey("uuid").(string)
		data.Time = res.Record().ValueByKey("_time").(time.Time)
		for _, field := range fields {
			col := res.Record().ValueByKey(field)
			if val, ok := col.(float64); ok {
				value.FieldByName(utils.FirstToUpper(field)).SetFloat(val)
			}
		}
		ret = append(ret, data)
	}
	if res.Err() != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, res.Err().Error()))
		logger.Logger.Warnw(res.Err().Error())
		return
	}
	c.JSON(http.StatusOK, models.NewResp(0, "ok", ret))
}

// GetHeartbeat
// @Tags heartbeat
// @Summary get heartbeat
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "token"
// @Param query-params query models.RangeQuery true "范围, 查询字段"
// @Success 200 {object} models.Resp[[]models.HeartBeat] "success"
// @Failure 400 {object} models.Msg "failure"
// @Router /cow/heartbeat [get]
func GetHeartbeat(c *gin.Context) {
	var query models.RangeQuery
	if err := c.BindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, err.Error()))
		return
	}
	fields := strings.Split(query.Fields, ",")
	ranges := ranges(query.Start, query.Stop)
	filter := fieldFilter(fields)
	flux := buildFlux(ranges, filter)
	res, err := db.QueryAPI.Query(c, flux)
	if err != nil {
		logger.Logger.Warnw(err.Error())
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, err.Error()))
		return
	}
	var ret []models.HeartBeat
	for res.Next() {
		var data models.HeartBeat
		value := reflect.ValueOf(&data).Elem()
		data.Id = res.Record().ValueByKey("uuid").(string)
		data.Time = res.Record().ValueByKey("_time").(time.Time)
		for _, field := range fields {
			col := res.Record().ValueByKey(field)
			if val, ok := col.(float64); ok {
				value.FieldByName(utils.FirstToUpper(field)).SetFloat(val)
			}
		}
		ret = append(ret, data)
	}
	if res.Err() != nil {
		logger.Logger.Warnw(res.Err().Error())
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, res.Err().Error()))
		return
	}
	c.JSON(http.StatusOK, models.NewResp(0, "ok", ret))
}

func fieldFilter(fields []string) string {
	term := func(field string) string {
		return fmt.Sprintf("r._field == \"%s\"", field)
	}
	arr := make([]string, 0, len(fields))
	for _, field := range fields {
		arr = append(arr, term(field))
	}
	return strings.Join(arr, " or ")
}

func ranges(start string, stop string) string {
	if start == "" && stop == "" {
		panic("unreachable")
	} else if start == "" {
		return fmt.Sprintf("stop: %s", stop)
	} else if stop == "" {
		return fmt.Sprintf("start: %s", start)
	} else {
		return fmt.Sprintf("start: %s, stop: %s", start, stop)
	}
}

func buildFlux(ranges string, fieldFilter string) string {
	return fmt.Sprintf(
		`from(bucket:"farm-iot")
|> range(%s)
|> filter(fn: (r) => r._measurement == "cow") 
|> filter(fn: (r) => %s) 
|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")`,
		ranges,
		fieldFilter,
	)
}

func buildFlux4SelectedUuid(ranges string, fieldFilter string, uuid string) string {
	return fmt.Sprintf(
		`from(bucket:"farm-iot")
|> range(%s)
|> filter(fn: (r) => r._measurement == "cow" and  r.uuid == "%s") 
|> filter(fn: (r) => %s) 
|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")`,
		ranges,
		uuid,
		fieldFilter,
	)
}
