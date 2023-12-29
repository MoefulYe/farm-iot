package handler

import (
	"github.com/MoefulYe/farm-iot/http-server/db"
	"github.com/MoefulYe/farm-iot/http-server/grpc_service"
	"github.com/MoefulYe/farm-iot/http-server/logger"
	"github.com/MoefulYe/farm-iot/http-server/models"
	"github.com/MoefulYe/farm-iot/http-server/protoc-gen/grpc/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"pg/ent/device"
)

// GetCowInfoByUuid
// @Tags cow
// @Summary get cow info by uuid
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "jwt"
// @Param uuid path string true "uuid"
// @Success 200 {object} models.Resp[models.CowInfo] "success"
// @Failure 400 {object} models.Msg "failure"
// @Router /cow/{uuid} [get]
func GetCowInfoByUuid(c *gin.Context) {
	uuid1 := c.Param("uuid")
	id, err := uuid.Parse(uuid1)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, "illegal uuid!"))
		return
	}
	resp := new(models.CowInfo)
	d, err := db.PgClient.Device.
		Query().
		Where(device.IDEQ(id)).
		First(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, "no such cow"))
		logger.Logger.Warnw(err.Error())
		return
	}
	resp.Id = d.ID.String()
	resp.DeadAt = d.DeadAt
	resp.BornAt = d.BornAt
	resp.Reason = d.Reason
	resp.Parent = d.Parent.String()
	c.JSON(http.StatusOK, models.NewResp(0, "ok", resp))
}

// GetCowInfo
// @Tags cow
// @Summary get cow info
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "jwt"
// @Param Pagination query models.PaginationQuery true "分页"
// @Success 200 {object} models.Resp[Paged[models.CowInfo]] "success"
// @Failure 400 {object} models.Msg "failure"
// @Router /cow [get]
func GetCowInfo(c *gin.Context) {
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
		logger.Logger.Warnw(err.Error())
		return
	}

	var v []models.CowInfo
	if err = db.PgClient.Device.
		Query().
		Limit(limit).Offset(offset).Select(
		device.FieldID, device.FieldBornAt, device.FieldDeadAt,
		device.FieldReason, device.FieldParent,
	).Scan(c, &v); err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, "db error"))
		logger.Logger.Warnw(err.Error())
		return
	}
	c.JSON(http.StatusOK, models.NewResp(0, "ok", models.NewPaged(cnt, v)))
}

// SpawnCow
// @Tags cow
// @Summary spawn cow
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "jwt"
// @Success 200 {object} models.Msg "success"
// @Failure 400 {object} models.Msg "failure"
// @Router /cow/spawn [post]
func SpawnCow(c *gin.Context) {
	if _, err := grpc_service.Client.Spawn(c, &service.SpawnReq{}); err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, err.Error()))
		logger.Logger.Warnw(err.Error())
		return
	}
	c.JSON(http.StatusOK, models.MsgOnly(0, "ok"))
}

// KillCow
// @Tags cow
// @Summary Kill cow
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "jwt"
// @Param cows body models.KillCowReq true "uuids"
// @Success 200 {object} models.Msg "success"
// @Failure 400 {object} models.Msg "failure"
// @Router /cow/kill [post]
func KillCow(c *gin.Context) {
	params := new(models.KillCowReq)
	if err := c.BindJSON(params); err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, err.Error()))
	}
	if _, err := grpc_service.Client.Kill(
		c, &service.KillReq{
			List: params.Cows,
		},
	); err != nil {
		logger.Logger.Warnw(err.Error())
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, err.Error()))
	} else {
		c.JSON(http.StatusOK, models.MsgOnly(0, "ok"))
	}
}
