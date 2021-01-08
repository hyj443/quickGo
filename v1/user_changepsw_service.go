package v1

import (
	// "quickGo/model"
	"quickGo/model"
	"quickGo/serializer"
)

func (service *ChangePasswordService) Valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code:    serializer.UserInputError,
			Message: "两次输入的密码不同",
		}
	}
	return nil
}

// ChangePasswordService 修改用户密码
type ChangePasswordService struct {
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// Change 用传来的新密码更新到数据库
func (service *ChangePasswordService) Change(user *model.User) *serializer.Response {
	// 对新密码进行校验，两次输入是否相同
	if err := service.Valid(); err != nil {
		return err
	}
	// 对新密码进行加密
	if err := user.SetPassword(service.Password); err != nil {
		return &serializer.Response{
			Code:    serializer.ServerPanicError,
			Message: "对密码加密出现错误",
		}
	}
	// 将更新了密码后的用户模型保存到数据库
	if err := model.DB.Save(&user).Error; err != nil {
		return &serializer.Response{
			Code:    serializer.DatabaseWriteError,
			Message: "更新数据库出现错误",
		}
	}
	return &serializer.Response{
		Data:    serializer.BuildUser(*user),
		Message: "修改密码成功",
	}
}
