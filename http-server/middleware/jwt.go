package middleware

import (
	"github.com/MoefulYe/farm-iot/http-server/models"
	"github.com/MoefulYe/farm-iot/http-server/utils"
	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(200, models.ResponseList{Code: 1, Msg: "need token ", Data: ""})
			c.Abort()
			return
		}
		claims, err := utils.JWTParse(token)
		if err != nil {
			// token过期
			//if err == TokenExpired {
			//	c.JSON(200, models.ResponseList{Code: 1, Msg: "expired token ", Data: ""})
			//	c.Abort()
			//	return
			//}
			// 其他错误
			c.JSON(200, models.ResponseList{Code: 1, Msg: "unknown token err", Data: ""})
			c.Abort()
			return
		}
		// 将解析后的有效载荷claims重新写入gin.Context引用对象中
		c.Set("claims", claims)
	}

}
