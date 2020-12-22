package api

import (
	"github.com/gin-gonic/gin"
	"quickGo/v1"
)

func UserRegister(c *gin.Context)  {
	var service v1.UserRegisterService

	if err:= c.ShouldBind(&service);err==nil{
		// service.
	}

}


func UserLogin(c *gin.Context) {
	// 定义一个空的UserLoginService结构体 
	var service v1.UserLoginService 

	// 将接收到的信息绑定到service
	if err:=  c.ShouldBind(&service); err==nil{
		
		// 将拿到的登录信息进行验证
		service.Login()


	}else{

		// c.JSON(200,  )

	}




}
