package v2

import (
	"quickGo/model"
	"quickGo/serializer"
)

// ChangePasswordService 修改用户密码
type ChangePasswordService struct {
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// Valid 校验密码输入
func (service ChangePasswordService) Valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code:    serializer.UserInputError,
			Message: "两次输入的密码不相同",
		}
	}
	return nil
}

// ChangePsw 修改密码
func (service ChangePasswordService) ChangePsw(user *model.User) *serializer.Response {
	// 新密码的两次输入是否相同
	if err := service.Valid(); err != nil {
		return err
	}
	// 对新密码进行加密
	if err := user.SetPassword(service.Password); err != nil {
		return &serializer.Response{
			Code:    serializer.ServerPanicError,
			Message: "密码加密出现错误",
		}
	}
	// 将修改后的用户模型保存到数据库
	if err := model.DB.Save(&user).Error; err != nil {
		return &serializer.Response{
			Code:    serializer.DatabaseWriteError,
			Message: "更新数据库出现错误。",
		}
	}

	return &serializer.Response{
		Data:    serializer.BuildUser(*user),
		Message: "修改密码成功",
	}
}
