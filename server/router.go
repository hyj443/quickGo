package server

import (
	"os"
	"quickGo/middleware"
	"quickGo/api"
	"github.com/gin-gonic/gin"
)

// NewRouter 装载路由
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 加入session中间件
	if os.Getenv("RIM") == "use" {
		// 
		r.Use(middleware.SessionRedis(os.Getenv("SESSION_SECRE")))
	} else {
		r.Use(middleware.SessionCookie(os.Getenv("SESSION_SECRE")))
	}
	// 加入跨域中间件
	r.Use(middleware.Cors())
	// 获取当前的用户
	r.Use(middleware.GetCurUser())

	if os.Getenv("V1") == "on" {
		v1:= r.Group("/api/v1")

		v1.POST("user/login", api.UserLogin)



	}



	








	return r
}
