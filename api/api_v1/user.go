package apiv1

import (
	"net/http"
	"quickGo/api"
	"quickGo/serializer"
	v1 "quickGo/v1"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service v1.UserRegisterService
	// POST过来的注册信息绑定到service，调用它的Register进行注册
	if err := c.ShouldBind(&service); err == nil {
		// 先创建用户模型，再进行表单验证，进行密码加密，将用户模型存入数据库，最后返回响应给客户端的Response
		response := service.Register()
		c.JSON(http.StatusOK, response.Format())
	} else {
		// c.JSON(http.StatusOK, )
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service v1.UserLoginService

	// 将接收到的信息绑定到service
	if err := c.ShouldBind(&service); err == nil {
		// 将传过来的用户名去数据库找到对应的用户模型，验证密码是否正确，返回UserResponse
		res := service.Login()
		// 现在登录校验通过，要把用户id保存到Context上的session对象
		if userReponse, ok := res.Data.(serializer.UserResponse); ok {
			// 获取Context上的session对象
			s := sessions.Default(c)
			// 先清理session对象中所有键值对
			s.Clear()
			// 将用户的id保存为session的"user_id"字段
			s.Set("user_id", userReponse.Data.ID)
			s.Save()
		}
		c.JSON(http.StatusOK, res.Format())
	} else {
		// c.JSON(200,  )
	}
}

// UserInfo 用户详情接口函数
func UserInfo(c *gin.Context) {
	// 获取挂在Context上的当前用户模型
	user := api.GetCurrentUser(c)

	// 构建出response的结构体
	res := serializer.Response{
		Data: serializer.BuildUserResponse(*user),
	}

	// 渲染输出成JSON
	c.JSON(http.StatusOK, res.Format())
}

// UserLogout 用户登出的接口，清理session对象中保存的值
func UserLogout(c *gin.Context) {
	// 获取Context上的session对象
	s := sessions.Default(c)
	// 清理session对象的所有键值对，其实就是user_id字段
	s.Clear()
	s.Save()
	c.JSON(http.StatusOK, serializer.Response{
		Message: "登出成功",
	}.Format())
}

// ChangePassword 修改密码的接口
func ChangePassword(c *gin.Context) {
	// 获取挂在Context上的当前用户模型
	user := api.GetCurrentUser(c)

	var service v1.ChangePasswordService

	if err := c.ShouldBind(&service); err == nil {
		// service绑定了新的密码，调用Change，传入当前的user进行修改
		res := service.Change(user)
		c.JSON(http.StatusOK, res.Format())
	} else {

	}
}
