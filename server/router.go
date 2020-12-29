package server

import (
	"os"
	"quickGo/api"
	apiv1 "quickGo/api/api_v1"
	apiv2 "quickGo/api/api_v2"
	"quickGo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 装载路由
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 加入session中间件 会根据session密钥创建session存储store，然后创建session对象，挂载到context上
	if os.Getenv("RIM") == "use" {
		r.Use(middleware.SessionRedis(os.Getenv("SESSION_SECRE")))
	} else {
		r.Use(middleware.SessionCookie(os.Getenv("SESSION_SECRE")))
	}

	// 加入跨域中间件
	r.Use(middleware.Cors())
	// 中间件，获取session中的user_id，根据id去数据库查询出对应的用户模型，保存到gin context上
	r.Use(middleware.GetCurUser())

	r.GET("/", api.Index)
	r.GET("/ping", api.Ping)

	// V1版本是最基本的网站需要
	if os.Getenv("V1") == "on" {
		v1 := r.Group("/api/v1")

		if os.Getenv("RIM") != "use" {
			panic("v1的Session验证必须依赖于mysql和redis，需要将环境变量RIM设置为use，并配置mysql和redis的连接")
		}

		v1.POST("user/register", apiv1.UserRegister)
		v1.POST("user/login", apiv1.UserLogin)

		auth := v1.Group("")
		auth.Use(middleware.AuthRequired()) // context上有user用户模型（代表登录了）才会继续执行后面的handler，如果没有登录，则不会执行后面的handler

		auth.GET("user/info", apiv1.UserInfo)
		auth.DELETE("user/logout", apiv1.UserLogout) // 清理session对象中所有的值
		auth.PUT("user/changepassword", apiv1.ChangePassword)

		admin := auth.Group("")
		admin.Use(middleware.AuthAdmin())

	}

	if os.Getenv("V2") == "on" {
		if os.Getenv("RIM") != "use" {
			panic("v1的Session验证必须依赖于mysql和redis，需要将环境变量RIM设置为use，并配置mysql和redis的连接")
		}

		jwtGroup:= r.Group("/api/v2")
		jwtGroup.POST("user/register", apiv2.UserRegister)
		jwtGroup.POST("user/login", apiv2.UserLogin)

		jwt:=jwtGroup.Group("")
		jwt.Use(middleware.JwtRequired())
		jwt.GET("user/me", apiv2.UserMe)
		jwt.PUT("user/changepassword", apiv2.ChangePassword)
		// 注销
		jwt.DELETE("user/logout", apiv2.Logout)

	}

	return r
}
