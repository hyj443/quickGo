package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	
	config.AllowMethods = []string{"PUT", "PATCH", "GET", "POST", "HEAD", "OPTIONS", "DELETE"}
	config.AllowHeaders = []string{"Content-Type", "Origin", "Cookie", "Content-Length"}

	// 运行在ReleaseMode模式下才会进行跨域保护，开发过程中不会受到跨域的困扰
	if gin.Mode() == gin.ReleaseMode {
		config.AllowOrigins = []string{"http://www.xxx.com"}
	} else {
		config.AllowAllOrigins = true
	}

	config.AllowCredentials = true
	return cors.New(config)
}
