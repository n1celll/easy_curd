package service

import (
	"easy_curd/model"
	"easy_curd/serializer"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("username = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.Err(401, "账号或密码错误", nil)
	}

	if user.Password != service.Password {
		return serializer.Err(401, "账号密码错误", nil)
	}

	return serializer.BuildUserWithTokenResponse(user)
}
