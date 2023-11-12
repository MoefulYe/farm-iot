package handler

import (
	"fmt"
	"github.com/MoefulYe/farm-iot/database/postgres/ent/device"
	"github.com/MoefulYe/farm-iot/http-server/db"
	"github.com/MoefulYe/farm-iot/http-server/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

// Infohandler PingExample godoc
// @Tags getInfo
// @Schemes
// @Summary get all  info
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "json"
// @Success 200 {object} models.ResponseList{data=[]models.Info} "xiangying"
// @Router /cow [get]
func Infohandler(c *gin.Context) {
	var v []models.Info
	err := db.PgClient.Device.
		Query().
		Select(device.FieldID, device.FieldBornAt, device.FieldDeadAt, device.FieldReason).
		Scan(context.Background(), &v)

	if err != nil {
		fmt.Printf("%v", err)
		c.JSON(200, models.ResponseList{Code: 1, Msg: "getInfo failed", Data: ""})
		return
	}
	c.JSON(200, models.ResponseList{Code: 0, Msg: "Info ok", Data: v})
}
