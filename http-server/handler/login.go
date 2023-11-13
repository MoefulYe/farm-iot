package handler

import (
	"context"
	"github.com/MoefulYe/farm-iot/database/postgres/ent/user"
	"github.com/MoefulYe/farm-iot/http-server/db"
	"github.com/MoefulYe/farm-iot/http-server/models"
	"github.com/MoefulYe/farm-iot/http-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 登录
// @Tags Login
// @Accept json
// @Produce json
// @Param data body models.LoginReq true "账号和密码"
// @Success 200 {object} models.Resp[models.Token] "success"
// @Failure 400 {object} models.Msg "failure"
// @Router /login [post]
func Login(c *gin.Context) {
	var body models.LoginReq
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, "illegal body params!"))
		return
	}
	username := body.Username
	passwd := body.Passwd
	d, err := db.PgClient.User.Query().Where(user.UsernameEQ(username)).First(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, "no such username!"))
		return
	}
	if match, err := utils.Argon2Verify(
		passwd, d.Passwd,
	); match == false || err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, "wrong passwd!"))
		return
	}
	token, err := utils.JWTGenerate(utils.NewClaims(username))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, "fail to gen token!"))
		return
	}

	c.JSON(http.StatusOK, models.NewResp(0, "ok", models.Token{Token: token}))
	return
}
