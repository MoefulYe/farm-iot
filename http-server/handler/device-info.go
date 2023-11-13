package handler

import (
	"github.com/MoefulYe/farm-iot/database/postgres/ent/device"
	"github.com/MoefulYe/farm-iot/http-server/db"
	"github.com/MoefulYe/farm-iot/http-server/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// GetDeviceInfoByUuid
// @Tags getInfo
// @Summary get uuid info by uuid
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "jwt"
// @Param uuid path string true "uuid"
// @Success 200 {object} models.Resp[models.DeviceInfo] "success"
// @Failure 400 {object} models.Msg "failure"
// @Router /cow/{uuid} [get]
func GetDeviceInfoByUuid(c *gin.Context) {
	uuid1 := c.Param("uuid")
	id, err := uuid.Parse(uuid1)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, "illegal uuid!"))
		return
	}
	resp := new(models.DeviceInfo)
	d, err := db.PgClient.Device.
		Query().
		Where(device.IDEQ(id)).
		Select(
			device.FieldID, device.FieldBornAt, device.FieldDeadAt,
			device.FieldReason,
		).
		First(c)
	resp.Id = d.ID.String()
	resp.DeadAt = d.DeadAt
	resp.BornAt = d.BornAt
	resp.Reason = d.Reason
	if err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, "no such device"))
		return
	}
	c.JSON(http.StatusOK, models.NewResp(0, "ok", resp))
}

// GetDeviceInfo
// @Tags getInfo
// @Summary get all info
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "jwt"
// @Param Pagination query models.PaginationQuery true "分页"
// @Success 200 {object} models.Resp[Paged[models.DeviceInfo]] "success"
// @Failure 400 {object} models.Msg "failure"
// @Router /cow [get]
func GetDeviceInfo(c *gin.Context) {
	var query models.PaginationQuery
	if err := c.BindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, "bad request"))
		return
	}
	offset := (query.Page - 1) * query.Size
	limit := query.Size

	cnt, err := db.PgClient.Device.Query().Count(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, "db error"))
		return
	}

	var v []models.DeviceInfo
	if err = db.PgClient.Device.
		Query().
		Limit(limit).Offset(offset).Select(
		device.FieldID, device.FieldBornAt, device.FieldDeadAt,
		device.FieldReason,
	).Scan(c, &v); err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, "db error"))
		return
	}
	c.JSON(http.StatusOK, models.NewResp(0, "ok", models.NewPaged(cnt, v)))
}
