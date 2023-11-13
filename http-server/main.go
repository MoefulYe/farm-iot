package main

import (
	"github.com/MoefulYe/farm-iot/http-server/docs"
	"github.com/MoefulYe/farm-iot/http-server/handler"
	"github.com/MoefulYe/farm-iot/http-server/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @BasePath /api/v1
func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)
	s1 := r.Group("/cow")
	s1.Use(middleware.Jwt())
	{
		s1.GET("", handler.GetDeviceInfo)
		s1.GET("/:uuid", handler.GetDeviceInfoByUuid)
		s1.GET("/keep-alive", handler.GetKeepalive)
		s1.GET("/keep-alive/:uuid", handler.GetKeepaliveByUuid)
	} //?start=&end=
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err := r.Run(":8080"); err != nil {
		log.Fatalf(err.Error())
	}
}
