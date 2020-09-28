package model

import (
	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName      string `gorm:"column:username"`
	Password      string
	Nickname      string
	Avatar        string `gorm:"size:1000"`
	WechatUnionid string
}

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}
