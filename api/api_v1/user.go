package apiv1

import (
	"net/http"
	"quickGo/api"
	"quickGo/serializer"
	v1 "quickGo/v1"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口函数
func UserRegister(c *gin.Context) {
	var service v1.UserRegisterService

	if err := c.ShouldBind(&service); err == nil {
		response := service.Register()
		c.JSON(http.StatusOK, response.Format())
	} else {
		// c.JSON(http.StatusOK, )
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	// 定义一个空的UserLoginService结构体
	var service v1.UserLoginService

	// 将接收到的信息绑定到service
	if err := c.ShouldBind(&service); err == nil {
		// 将拿到的登录信息进行验证，然后返回一个Response
		res := service.Login()
		// 获取res的Data
		if userReponse, ok := res.Data.(serializer.UserResponse); ok {
			// 拿到了登录成功后的响应的用户结构体
			// 现在登录成功了，要把用户信息保存到session
			// 获取session对象
			s := sessions.Default(c)
			// 先清理一下
			s.Clear()
			// Set sets the session value associated to the given key.
			// 将用户模型的ID保存到session的"user"字段
			s.Set("user", userReponse.Data.ID)
			s.Save()
		}

		c.JSON(http.StatusOK, res.Format())

	} else {
		// c.JSON(200,  )
	}
}

// UserInfo 用户详情接口函数
func UserInfo(c *gin.Context) {
	// 获取c上的用户模型
	user := api.GetCurrentUser(c)

	// 构建出response的结构体
	res := serializer.Response{
		Data: serializer.BuildUserResponse(*user),
	}

	// 渲染输出成JSON
	c.JSON(http.StatusOK, res.Format())
}

// UserLogout 用户登出的接口
// 就是清理session对象中保存的值
func UserLogout(c *gin.Context) {
	// 获取context上的session对象
	s := sessions.Default(c)
	// 清理session对象的所有键值对
	s.Clear()
	s.Save()
	c.JSON(http.StatusOK, serializer.Response{
		Message: "登出成功",
	}.Format())
}

// ChangePassword 修改密码的接口
func ChangePassword(c *gin.Context) {
	user := api.GetCurrentUser(c)
	var service v1.ChangePasswordService

	if err := c.ShouldBind(&service); err == nil {
		res := service.Change(user)
		c.JSON(http.StatusOK, res.Format())
	} else {

	}

}
