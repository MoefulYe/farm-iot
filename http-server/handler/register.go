package handler

import (
	"context"
	"github.com/MoefulYe/farm-iot/http-server/db"
	"github.com/MoefulYe/farm-iot/http-server/models"
	"github.com/MoefulYe/farm-iot/http-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register 注册用户
// @Tags register
// @Accept json
// @Produce json
// @Param body-params body models.RegisterReq true "用户名和密码"
// @Success 200 {object} models.Resp[models.Token] "success"
// @Failure 400 {object} models.Msg "failure"
// @Router /register [post]
func Register(c *gin.Context) {
	var body models.RegisterReq
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, "illegal body params!"))
		return
	}
	username := body.Username
	passwd := body.Passwd
	hashedPasswd, err := utils.Argon2Generate(passwd)
	if _, err = db.PgClient.User.Create().SetUsername(username).SetPasswd(hashedPasswd).Save(context.Background()); err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, "register failed!"))
		return
	}
	claims := utils.NewClaims(username)
	token, err := utils.JWTGenerate(claims)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.MsgOnly(1, "fail to gen token!"))
		return
	}
	c.JSON(http.StatusOK, models.NewResp(0, "ok", models.Token{Token: token}))
}
