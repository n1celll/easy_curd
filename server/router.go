package server

import (
	"easy_curd/api"
	"easy_curd/middleware"
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api")
	{

		user := v1.Group("/user")
		{
			// 用户登录
			user.POST("login", api.UserLogin)
			user.POST("socialLogin", api.UserSocialLogin)
			// 需要登录保护的
			user.Use(middleware.AuthRequired())
			// 获取用户信息
			user.GET("userInfo", api.UserInfo)
		}

	}
	return r
}
