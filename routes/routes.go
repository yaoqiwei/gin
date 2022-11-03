package routes

import (
	"gin/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouter ...
func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {

	router := gin.New()
	router.ForwardedByClientIP = true
	router.Use(middlewares...)

	// 错误中间件
	router.Use(middleware.RecoveryMiddleware())

	// 跨域中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	// 公共中间件
	router.Use(cors.New(config))
	router.Use(middleware.LogMiddleware())

	return router
}
