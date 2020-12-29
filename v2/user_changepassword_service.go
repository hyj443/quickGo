package v2

import (
	"quickGo/model"
	"quickGo/serializer"
)

// ChangePassword 修改用户密码
type ChangePasswordService struct {
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

func (service ChangePasswordService) Valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code:    serializer.UserInputError,
			Message: "两次输入的密码不相同",
		}
	}
	return nil
}

func (service ChangePasswordService) ChangePsw(user *model.User) *serializer.Response {
	if err := service.Valid(); err != nil {
		return err
	}
	if err := user.SetPassword(service.Password); err != nil {
		return &serializer.Response{
			Code:    serializer.ServerPanicError,
			Message: "密码加密出现错误",
		}
	}
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
