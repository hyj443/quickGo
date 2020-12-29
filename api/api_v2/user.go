package apiv2

import (
	"net/http"
	"quickGo/api"
	"quickGo/cache"
	"quickGo/serializer"
	v2 "quickGo/v2"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口函数
func UserRegister(c *gin.Context) {
	var service v2.UserRegisterService

	if err := c.ShouldBind(&service); err == nil {
		response := service.Register()
		c.JSON(http.StatusOK, response.Format())
	} else {
		// c.JSON(http.StatusOK, )
	}
}

func UserLogin(c *gin.Context) {
	var service v2.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login()
		c.JSON(200, res.Format())
	} else {
		// c.JSON(200, api.ErrorResponse(err).Result())
	}
}

func UserMe(c *gin.Context) {
	user := api.GetCurrentUser(c)
	res := serializer.Response{
		Data: serializer.BuildUserResponse(*user),
	}
	c.JSON(http.StatusOK, res.Format())
}

func ChangePassword(c *gin.Context) {
	user := api.GetCurrentUser(c)
	var service v2.ChangePasswordService

	if err := c.ShouldBind(&service); err != nil {
		res := service.ChangePsw(user)
		c.JSON(http.StatusOK, res.Format())
	} else {
		//
	}
}

func Logout(c *gin.Context) {
	token, _ := c.Get("token")
	tokenStr := token.(string)
	// 在redis的jwt:baned这个集合中添加value：tokenStr
	cache.RedisClient.SAdd("jwt:baned", tokenStr)
	c.JSON(http.StatusOK, serializer.Response{
		Message: "注销成功",
	})
}
