package main

import (
	"github.com/MoefulYe/farm-iot/http-server/docs"
	"github.com/MoefulYe/farm-iot/http-server/handler"
	"github.com/MoefulYe/farm-iot/http-server/logger"
	"github.com/MoefulYe/farm-iot/http-server/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api
func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"
	r.POST("/api/login", handler.Login)
	r.POST("/api/register", handler.Register)
	cowApi := r.Group("/api/cow/")
	cowApi.Use(middleware.Jwt())
	{
		cowApi.GET("heartbeat", handler.GetHeartbeat)
		cowApi.GET("heartbeat/:uuid", handler.GetHeartbeatByUuid)
		cowApi.POST("spawn", handler.SpawnCow)
		cowApi.POST("kill", handler.KillCow)
		cowApi.GET(":uuid", handler.GetCowInfoByUuid)
		cowApi.GET("", handler.GetCowInfo)
	}
	r.GET("/api/balance", middleware.Jwt(), handler.GetBalance)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err := r.Run(":8080"); err != nil {
		logger.Logger.Fatalw(err.Error())
	}
}
