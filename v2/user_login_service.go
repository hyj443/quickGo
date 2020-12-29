package v2

import (
	"quickGo/auth"
	"quickGo/conf"
	"quickGo/model"
	"quickGo/serializer"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=18"`
}

func (service UserLoginService) Login() *serializer.Response {
	var user model.User
	ExpireTime := time.Now().Add(time.Hour * time.Duration(720)).Unix()

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return &serializer.Response{
			Code:    serializer.UserNotFoundError,
			Message: "账号或密码错误",
		}
	}

	if user.CheckPassword(service.Password) == false {
		return &serializer.Response{
			Code:    serializer.UserNotFoundError,
			Message: "账号或密码错误",
		}
	}

	// 为当前的用户模型，生成一段token字符串，放到响应中，返回给客户端
	token, err := GenerateToken(user, ExpireTime)
	if err != nil {
		return &serializer.Response{
			Code:  serializer.ServerPanicError,
			Error: err.Error(),
		}
	}

	return &serializer.Response{
		Data: gin.H{
			"access_token": token,
			"expire_at":    ExpireTime,
			"token_type":   "Bearer",
		},
	}
}

func GenerateToken(user model.User, expireTime int64) (string, error) {
	claims := auth.Jwt{
		jwt.StandardClaims{
			ExpiresAt: expireTime,        // 传入过期时间
			IssuedAt:  time.Now().Unix(), // 生成token的当前时间
		},
		user,
	}
	// func NewWithClaims(method SigningMethod, claims Claims) *Token {
	// 创建一个新的token结构体
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Get the complete, signed token
	// 获取完整的签名token
	jwtStr, err := token.SignedString(conf.SignKey)

	return jwtStr, err
}
