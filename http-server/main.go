package main

import (
	"github.com/MoefulYe/farm-iot/http-server/config"
	"github.com/MoefulYe/farm-iot/http-server/docs"
	"github.com/MoefulYe/farm-iot/http-server/handler"
	"github.com/MoefulYe/farm-iot/http-server/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @BasePath /api
func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"
	r.POST("/api/login", handler.Login)
	r.POST("/api/register", handler.Register)
	s1 := r.Group("/api/cow/")
	s1.Use(middleware.Jwt())
	{
		s1.GET("keep-alive", handler.GetKeepalive)
		s1.GET("keep-alive/:uuid", handler.GetKeepaliveByUuid)
		s1.GET(":uuid", handler.GetDeviceInfoByUuid)
		s1.GET(":uuid/kill", handler.KillDevice)
		s1.GET("", handler.GetDeviceInfo)
	} //?start=&end=
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err := r.Run(config.Conf.Port); err != nil {
		log.Fatalf(err.Error())
	}
}
