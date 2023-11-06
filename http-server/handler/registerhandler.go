package handler

import (
	"github.com/MoefulYe/farm-iot/iot-server/db"
	"github.com/MoefulYe/farm-iot/iot-server/utils"
	"github.com/gin-gonic/gin"
)

func Registerhandler(c *gin.Context) {
	var body struct {
		username string `json:"username"`
		passwd   string `json:"passwd"`
	}
	if err := c.BindJSON(&body); err != nil {
		return
	}
	username := body.username
	passwd := body.passwd
	hashedPasswd, err := utils.Argon2Generate(passwd)
	if _, err = db.PgClient.User.Create().Setusername(username).SetPasswd(hashedPasswd).Save(context.Background()); err != nil {
		handleRegisterResult(
			server, payload.GetUuid(), cow.RegisterResp_FAILED, "",
		)
		return
	}
	return
}
