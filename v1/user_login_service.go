package v1

import (
	"quickGo/model"
	"quickGo/serializer"
)

type UserLoginService struct {
	UserName string `form:"user_name",json:"user_name",binding:"required,min=5,max=20"`
	PassWord string `form:"password",json:"password",binding:"required,min=8,max=20"`
}

func (service UserLoginService) Login() *serializer.Response {

	// 现在要用收到的用户信息去数据库查找对应的用户

	var user model.User

	// 从数据库中查找user_name字段和传入的用户名匹配的第一条数据，存储到user中
	err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error

	// 找不到该用户
	if err != nil {
		return &serializer.Response{
			Code:    40000,
			Message: "用户名或密码不正确",
		}
	}

	// 找到了这个用户，但是传过来的密码不对
	if !user.CheckPassword(service.PassWord){
		return &serializer.Response{
			Code: 40001,
			Message:"用户名或密码不准确",
		}
	}

	// 现在用户名和密码都对了，现在要为这个登录上来的用户生成一个token
		
	






	return &serializer.Response{
		// Data:
	}
}
