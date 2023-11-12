package handler

import (
	"context"
	"github.com/MoefulYe/farm-iot/http-server/db"
	"github.com/MoefulYe/farm-iot/http-server/models"
	"github.com/MoefulYe/farm-iot/http-server/utils"
	"github.com/gin-gonic/gin"
)

// Registerhandler PingExample godoc
// @Schemes
// @Tags register
// @Accept json
// @Produce json
// @Param data body models.RegisterList true "json"
// @Success 200 {object} models.ResponseList{data=models.Token} "xiangying"
// @Router /register [post]
func Registerhandler(c *gin.Context) {
	var body models.RegisterList
	if err := c.BindJSON(&body); err != nil {
		return
	}
	username := body.Username
	passwd := body.Passwd
	hashedPasswd, err := utils.Argon2Generate(passwd)
	if _, err = db.PgClient.User.Create().SetUsername(username).SetPasswd(hashedPasswd).Save(context.Background()); err != nil {
		c.JSON(200, models.ResponseList{Code: 0, Msg: "register fail", Data: models.Token{Token: ""}})
		return
	}
	claims := utils.NewClaims(username)
	token, err := utils.JWTGenerate(claims)
	if err != nil {
		c.JSON(200, models.ResponseList{Code: 0, Msg: "register fail", Data: models.Token{Token: ""}})
		return
	}
	c.JSON(200, models.ResponseList{Code: 1, Msg: "ok", Data: models.Token{Token: token}})
}
