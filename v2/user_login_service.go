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

// Login 用户登录函数
func (service UserLoginService) Login() *serializer.Response {
	var user model.User
	// token过期时间点
	ExpireTime := time.Now().Add(time.Hour * time.Duration(720)).Unix()

	// 根据登录的用户名找到数据库中对应的用户模型，找不到就是用户名不存在
	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return &serializer.Response{
			Code:    serializer.UserNotFoundError,
			Message: "账号或密码不正确",
		}
	}

	// 校验密码是否正确
	if !user.CheckPassword(service.Password) {
		return &serializer.Response{
			Code:    serializer.UserNotFoundError,
			Message: "账号或密码不正确",
		}
	}

	// 用户登录，为用户生成专属唯一的token，放到响应中，返回给客户端
	token, err := GenerateToken(user, ExpireTime)
	// 生成token失败，把错误返回给客户端
	if err != nil {
		return &serializer.Response{
			Code:  serializer.ServerPanicError,
			Error: err.Error(),
		}
	}
	// 生成token成功，将放入Data中，还有过期时间，给客户端
	return &serializer.Response{
		Data: gin.H{
			"access_token": token,
			"expire_at":    ExpireTime,
			"token_type":   "Bearer",
		},
	}
}

// GenerateToken 生成token字符串
func GenerateToken(user model.User, expireTime int64) (string, error) {
	claims := auth.Jwt{
		jwt.StandardClaims{
			ExpiresAt: expireTime,        // 传入过期时间
			IssuedAt:  time.Now().Unix(), // 生成token的当前时间
		},
		user,
	}
	// func NewWithClaims(method SigningMethod, claims Claims) *Token {
	// 创建一个token结构体
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Get the complete, signed token
	tokenStr, err := token.SignedString(conf.SignKey)

	return tokenStr, err
}
