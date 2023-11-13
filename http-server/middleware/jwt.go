package middleware

import (
	"github.com/MoefulYe/farm-iot/http-server/models"
	"github.com/MoefulYe/farm-iot/http-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusBadRequest, models.MsgOnly(1, "expect token"))
			c.Abort()
			return
		}
		claims, err := utils.JWTParse(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.MsgOnly(1, "invalid token"))
			c.Abort()
			return
		}
		// 将解析后的有效载荷claims重新写入gin.Context引用对象中
		c.Set("claims", claims)
	}

}
