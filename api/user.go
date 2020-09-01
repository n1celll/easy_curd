package api

import (
	"easy_curd/serializer"
	"easy_curd/service"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var s service.UserRegisterService
	if err := c.ShouldBind(&s); err == nil {
		res := s.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var s service.UserLoginService
	if err := c.ShouldBind(&s); err == nil {
		res := s.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserInfo(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(200, res)
}
