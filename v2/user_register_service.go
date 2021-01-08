package v2

import (
	"quickGo/model"
	"quickGo/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=18"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=18"`
}

// Valid 验证表单
func (service *UserRegisterService) Valid() *serializer.Response {
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
			Message: "昵称被占用",
		}
	}

	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code:    serializer.UserRepeatError,
			Message: "用户名被占用",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() *serializer.Response {
	user := model.User{
		NickName: service.Nickname,
		UserName: service.UserName,
		Status:   model.Active,
	}

	// 表单验证
	if err := service.Valid(); err != nil {
		return err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return &serializer.Response{
			Code:    serializer.ServerPanicError,
			Message: "密码加密失败",
		}
	}

	// 将user用户模型保存到数据库
	if err := model.DB.Create(&user).Error; err != nil {
		return &serializer.Response{
			Code:    serializer.DatabaseWriteError,
			Message: "注册失败",
		}
	}

	return &serializer.Response{
		Data: serializer.BuildUserResponse(user),
	}
}
