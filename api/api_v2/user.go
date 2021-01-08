package apiv2

import (
	"net/http"
	"quickGo/api"
	"quickGo/cache"
	"quickGo/serializer"
	v2 "quickGo/v2"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var service v2.UserRegisterService

	if err := c.ShouldBind(&service); err == nil {
		response := service.Register()
		c.JSON(http.StatusOK, response.Format())
	} else {
		// c.JSON(http.StatusOK, )
	}
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	var service v2.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login()
		c.JSON(200, res.Format())
	} else {
		// c.JSON(200, api.ErrorResponse(err).Result())
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	// 从Context上获取当前用户模型
	user := api.GetCurrentUser(c)
	res := serializer.Response{
		Data: serializer.BuildUserResponse(*user),
	}
	c.JSON(http.StatusOK, res.Format())
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	// 从Context上获取当前用户模型
	user := api.GetCurrentUser(c)

	var service v2.ChangePasswordService
	// 将新密码信息绑定到service
	if err := c.ShouldBind(&service); err != nil {
		res := service.ChangePsw(user)
		c.JSON(http.StatusOK, res.Format())
	} else {
		//
	}
}

// Logout 用户注销
func Logout(c *gin.Context) {
	// 获取保存在Context上的token字符串
	token, _ := c.Get("token")
	tokenStr := token.(string)
	// 在redis的jwt:banned这个集合中添加tokenStr
	cache.RedisClient.SAdd("jwt:banned", tokenStr)
	
	c.JSON(http.StatusOK, serializer.Response{
		Message: "注销成功",
	})
}
