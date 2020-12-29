package serializer

import (
	"quickGo/model"
)

type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Nickname  string `json:"nickname"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	Admin     bool   `json:"admin"`
	CreatedAt int64  `json:"created_at"`
}

// UserResponse 单个用户序列化输出格式对象
type UserResponse struct {
	Data User `json:"user"`
}

func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Nickname:  user.NickName,
		Status:    user.Status,
		Avatar:    user.Avatar,
		Admin:     user.SuperUser,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

func BuildUserResponse(user model.User) UserResponse {
	return UserResponse{
		Data: BuildUser(user),
	}
}