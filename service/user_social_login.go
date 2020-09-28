package service

import (
	"easy_curd/model"
	"easy_curd/serializer"
	"easy_curd/wechat"
	"github.com/gin-gonic/gin"
)

type UserSocialLoginService struct {
	Channel string `form:"channel" json:"channel" binding:"required"`
	Code    string `form:"code" json:"code"`
}

// Login 用户登录函数
func (service *UserSocialLoginService) SocialLogin(c *gin.Context) serializer.Response {
	var user model.User
	channel := service.Channel
	if channel == "wechat" {
		userInfo, err := wechat.GetUserInfo(service.Code)
		if userInfo == nil {
			return serializer.Err(401, "微信登录失败", err)
		}
		unionid := userInfo.Unionid
		if unionid == "" {
			return serializer.Err(401, "", nil)
		}
		err = model.DB.Where("wechat_unionid = ? ", unionid).First(&user).Error
		if err != nil {
			user.Nickname = userInfo.Nickname
			user.Avatar = userInfo.Avatar
			user.WechatUnionid = userInfo.Unionid
			model.DB.Create(&user)
		}
	} else {
		return serializer.ParamsErr(nil)
	}
	return serializer.BuildUserResponse(user)
}
