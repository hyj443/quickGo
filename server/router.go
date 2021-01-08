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

	// 根据use选择session存储引擎创建store，再返回session中间件
	// 这个中间件会创建一个名为gin-session的session对象，会挂载到Context
	if os.Getenv("RIM") == "use" {
		r.Use(middleware.SessionRedis(os.Getenv("SESSION_SECRET")))
	} else {
		r.Use(middleware.SessionCookie(os.Getenv("SESSION_SECRET")))
	}

	// 添加跨域中间件
	r.Use(middleware.Cors())

	// GetCurUser 中间件，从Context上获取session对象，获取当前登录用户user_id
	// 根据user_id去数据库查询出对应的用户模型，将用户模型保存到Context
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

		// 登录才有权限的路由组
		auth := v1.Group("")

		// 检查是否登录的中间件
		// context上有用户模型（代表登录了）才会继续执行后面的handler，如果没有，则不会执行后面的handler
		auth.Use(middleware.AuthRequired())

		auth.GET("user/info", apiv1.UserInfo)

		auth.DELETE("user/logout", apiv1.UserLogout) 

		auth.PUT("user/changepassword", apiv1.ChangePassword)

		// 需要登录的，并且是管理员身份的路由组
		admin := auth.Group("")
		admin.Use(middleware.AuthAdmin())
	}

	if os.Getenv("V2") == "on" {
		if os.Getenv("RIM") != "use" {
			panic("v2的Session验证必须依赖于mysql和redis，需要将环境变量RIM设置为use，并配置mysql和redis的连接")
		}

		jwtGroup := r.Group("/api/v2")

		jwtGroup.POST("user/register", apiv2.UserRegister)

		jwtGroup.POST("user/login", apiv2.UserLogin)

		jwt := jwtGroup.Group("")

		jwt.Use(middleware.JwtRequired())

		jwt.GET("user/me", apiv2.UserMe)

		jwt.PUT("user/changepassword", apiv2.ChangePassword)

		// 注销
		jwt.DELETE("user/logout", apiv2.Logout)
	}

	return r
}
