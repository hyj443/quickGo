package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string `gorm:"size:1000"` // 规定这个字段的大小（列的大小）
	SuperUser      bool
}
