package handler

import (
	"github.com/MoefulYe/farm-iot/http-server/db"
	"github.com/MoefulYe/farm-iot/http-server/logger"
	"github.com/MoefulYe/farm-iot/http-server/models"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "pg/ent"
	"pg/ent/balance"
	"time"
)

// GetBalance
// @Tags balance
// @Summary get balance
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "token"
// @Param query-params query models.TimeRange true "范围"
// @Success 200 {object} models.Resp[[]float64] "success"
// @Failure 400 {object} models.Msg "failure"
// @Router /balance [get]
func GetBalance(c *gin.Context) {
	params := new(models.TimeRange)
	if err := c.BindQuery(params); err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, err.Error()))
		return
	}
	query := db.PgClient.Balance.Query()
	if params.From != "" {
		from, err := time.Parse(time.RFC3339, params.From)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.MsgOnly(1, err.Error()))
			return
		}
		query.Where(balance.WhenGTE(from))
	}
	if params.To != "" {
		to, err := time.Parse(time.RFC3339, params.To)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.MsgOnly(1, err.Error()))
			return
		}
		query.Where(balance.WhenLTE(to))
	}
	ret, err := query.All(c)
	if err != nil {
		logger.Logger.Warnw(err.Error())
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, models.NewResp(0, "ok", ret))
}
