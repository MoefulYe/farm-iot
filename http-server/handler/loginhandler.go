package handler

import (
	"context"
	"github.com/MoefulYe/farm-iot/database/postgres/ent/user"
	"github.com/MoefulYe/farm-iot/http-server/db"
	"github.com/MoefulYe/farm-iot/http-server/models"
	"github.com/MoefulYe/farm-iot/http-server/utils"
	"github.com/gin-gonic/gin"
)

// Loginhandler PingExample godoc
// @Tags Login
// @Schemes
// @Accept json
// @Produce json
// @Param data body models.RegisterList true "json"
// @Success 200 {object} models.ResponseList{data=models.Token} "xiangying"
// @Router /login [post]
func Loginhandler(c *gin.Context) {
	var body models.RegisterList
	if err := c.BindJSON(&body); err != nil {
		return
	}
	username := body.Username
	passwd := body.Passwd
	d, err := db.PgClient.User.Query().Where(user.UsernameEQ(username)).First(context.Background())
	if err != nil {
		c.JSON(200, models.ResponseList{Code: 0, Msg: "login fail", Data: models.Token{Token: ""}})
		return
	}
	if match, err := utils.Argon2Verify(passwd, d.Passwd); match == false || err != nil {
		c.JSON(200, models.ResponseList{Code: 0, Msg: "login fail", Data: models.Token{Token: ""}})
		return
	}
	token, err := utils.JWTGenerate(utils.NewClaims(username))
	if err != nil {
		c.JSON(200, models.ResponseList{Code: 0, Msg: "login fail", Data: models.Token{Token: ""}})
		return
	}

	c.JSON(200, models.ResponseList{Code: 1, Msg: "ok", Data: models.Token{Token: token}})
	return
}
