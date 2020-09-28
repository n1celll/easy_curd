package api

import (
	"easy_curd/serializer"
	"easy_curd/service"

	"github.com/gin-gonic/gin"
)

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var s service.UserLoginService
	if err := c.ShouldBind(&s); err == nil {
		res := s.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ParamsErr(err))
	}
}

func UserSocialLogin(c *gin.Context) {
	var s service.UserSocialLoginService
	if err := c.ShouldBind(&s); err == nil {
		res := s.SocialLogin(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ParamsErr(err))
	}
}

// UserMe 用户详情
func UserInfo(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(200, res)
}
