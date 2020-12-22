package v1

import (
	"quickGo/model"
	"quickGo/serializer"
)

type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=18"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=18"`
}

// Valid 验证表单
func (service UserRegisterService) Valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code:    serializer.UserPasswordError,
			Message: "两次输入的密码不相同",
		}
	}

	count := 0
	model.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code:    serializer.UserRepeatError,
			Message: "昵称已存在",
		}
	}
	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code:    serializer.UserRepeatError,
			Message: "用户名已经存在",
		}
	}

	return nil
}


// Valid 做的事是将post过来的账密进行验证，首先看两次输入的密码是否一致，再看用户名和昵称是否已经存在于数据库，即是否已经有重复的数据

func (service *UserRegisterService)Register ()*serializer.Response  {
	user:= model.User{
		NickName: service.Nickname,
		UserName: service.UserName,
		Status: model.Active,
	}
	// 先进行表单验证
	if err:= service.Valid();err!=nil{
		return err
	}

	// 进行密码加密
	if  err:=  user.SetPassword(service.Password);err!=nil {
		return &serializer.Response{
			Code:serializer.ServerPanicError,
			Message:"密码加密失败",
		}
	}

	if err:=model.DB.Create(&user).Error ; err!=nil {
		return &serializer.Response{
			code:serializer.DatabaseWriteError,
			Message:"注册信息存入数据库失败",
		}
	}
	return &serializer.Response{
		// Data: serializer.
	}
	
}