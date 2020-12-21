package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User 用户的模型
type User struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string `gorm:"size:1000"` // 规定这个字段的大小（列的大小）
	SuperUser      bool
}

// SetPassword 对密码进行加密，生成哈希密码，保存到user模型的PasswordDigest，加密的强度是12
func (user *User) SetPassword(password string) error {
	//func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return err
	}

	user.PasswordDigest = string(bytes)

	return nil

}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	// func CompareHashAndPassword(hashedPassword, password []byte) error
	// 获取用户user的哈希密码，通过下面的函数，与传入的密码进行比对
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
