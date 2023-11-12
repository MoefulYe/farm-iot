package handler

import (
	"fmt"
	"github.com/MoefulYe/farm-iot/database/postgres/ent/device"
	"github.com/MoefulYe/farm-iot/http-server/db"
	"github.com/MoefulYe/farm-iot/http-server/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

// Infouuidhandler PingExample godoc
// @Tags getInfo
// @Schemes
// @Summary get uuid info
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "json"
// @Param uuid path string    true  "ID"
// @Success 200 {object} models.Resp[models.Info] "xiangying"
// @Router /cow/{uuid} [get]
func Infouuidhandler(c *gin.Context) {
	uuid1 := c.Param("uuid")
	id, err := uuid.Parse(uuid1)
	if err != nil {
		fmt.Printf("%v", err)
		c.JSON(200, models.ResponseList{Code: 1, Msg: "getInfo failed", Data: ""})
		return
	}
	//var v models.Info
	d, err := db.PgClient.Device.
		Query().
		Where(device.IDEQ(id)).
		Select(device.FieldID, device.FieldBornAt, device.FieldDeadAt, device.FieldReason).
		First(context.Background())
	fmt.Printf("%v", d)
	if err != nil {
		fmt.Printf("%v", err)
		c.JSON(200, models.ResponseList{Code: 1, Msg: "getInfo failed", Data: ""})
		return
	}
	c.JSON(200, models.ResponseList{Code: 0, Msg: "Info ok", Data: d})
}
