package main

import (
	"github.com/MoefulYe/farm-iot/http-server/docs"
	"github.com/MoefulYe/farm-iot/http-server/handler"
	"github.com/MoefulYe/farm-iot/http-server/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.POST("/login", handler.Loginhandler)
	r.POST("/register", handler.Registerhandler)
	s1 := r.Group("/cow")
	s1.Use(middleware.Jwt())
	{
		s1.GET("/", handler.Infohandler)
		s1.GET("/:uuid", handler.Infouuidhandler)
		s1.GET("/keep-alive", handler.Keepalivehandler)
		s1.GET("/keep-alive/:uuid", handler.Keepaliveuuidhandler)
	} //?start=&end=
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")

}
