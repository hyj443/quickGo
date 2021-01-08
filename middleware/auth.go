package middleware

import (
	"net/http"
	"quickGo/model"
	"quickGo/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GetCurUser 中间件，从Context上获取session对象，获取它上面的user_id，能获取到代表已登录
// 根据user_id在数据库中找到对应的用户模型，将用户模型保存到Context的user字段
func GetCurUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			user, err := model.GetUser(uid)
			if err != nil {
				c.Set("user", user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录的路由要走的中间件
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			// 因为c.Keys的val是interface{}类型，现在要推断它是不是User类型
			// 如果是，那就Context上有用户模型，即已经登录
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}
		// 用户没有登录
		c.JSON(http.StatusOK, serializer.Response{
			Code:    serializer.UserNotPermissionError,
			Message: "需要先登录",
		}.Format())

		// 防止等待的handler被调用
		c.Abort()
	}
}

// AuthAdmin 中间件，获取Context挂载的用户模型，如果SuperUser字段为真，是管理员，c.Next()
func AuthAdmin() gin.HandlerFunc{
	return func(c *gin.Context){
		if user,_ := c.Get("user");user!= nil {
			if user.(*model.User).SuperUser {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusOK, serializer.Response{
			Code: serializer.UserNotPermissionError,
			Message:"需要管理员权限，才能进行此操作",
		}.Format())
		c.Abort()
	}
}
