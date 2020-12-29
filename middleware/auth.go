package middleware

import (
	"net/http"
	"quickGo/model"
	"quickGo/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GetCurUser 中间件函数，从context上获取ss，从ss获取当前登录的用户id，根据id获取用户模型，将用户模型保存到context上
func GetCurUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从context上获取session对象
		session := sessions.Default(c)
		// 获取session对象中user_id对应的用户id
		uid := session.Get("user_id")
		if uid != nil {
			// 根据用户id获取到用户模型
			user, err := model.GetUser(uid)
			if err != nil {
				c.Set("user", user)
			}
		}
		c.Next()
	}
}

// AuthRequired 所有需要登录的路由要走的中间件
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 怎么确定有没有登录，就看看context上有没有user的数据
		if user, _ := c.Get("user"); user != nil {
			// 因为本来是interface{}类型，现在要推断它是不是User类型，如果是，那就成功挂载了
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

		// prevents pending handlers from being called
		c.Abort()
	}
}

// AuthAdmin 中间件，获取context上挂载的用户模型，看看SuperUser字段是否为真，看看是不是管理员
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
			Message:"你没有权限进行此操作",
		}.Format())
		c.Abort()
	}
}
