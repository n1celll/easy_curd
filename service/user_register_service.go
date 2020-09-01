package service

import (
	"easy_curd/model"
	"easy_curd/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Status: 400,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Status: 400,
			Msg:  "用户名已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() serializer.Response {
	user := model.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Password: service.Password,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Err(400, "注册失败", err)
	}
	return serializer.BuildUserWithTokenResponse(user)
}
